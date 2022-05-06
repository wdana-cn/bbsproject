package controller

import (
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/wdana-cn/ginpro/common"
	"github.com/wdana-cn/ginpro/model"
	"github.com/wdana-cn/ginpro/response"
	"github.com/wdana-cn/ginpro/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func IsPhoneNumExist(DB *gorm.DB, telephone string) bool {
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	//获取参数
	user_name := ctx.PostForm("name")
	user_telephone := ctx.PostForm("telephone")
	user_passwd := ctx.PostForm("passwd")
	user_email := ctx.PostForm("email")
	user_avatar, err := ctx.FormFile("avator")
	if err != nil {
		log.Println(err)
	}
	filepath := path.Join("./uploads/", user_avatar.Filename)
	fileExt := strings.ToLower(path.Ext(user_avatar.Filename))
	if fileExt != ".jpg" && fileExt != ".png" {
		log.Println("file doesn't jpgtype")
		response.Fail(ctx, "only access jpg&png", gin.H{})
	}
	err = ctx.SaveUploadedFile(user_avatar, filepath)
	if err != nil {
		log.Println(err)
	}

	//数据验证
	if len(user_telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为十一位！")
		return
	} else {
		if IsPhoneNumExist(DB, user_telephone) == true {
			response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号已存在！")
			return
		}
	}
	//判断参数合法性
	if len(user_passwd) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码长度不能小于六位！")
		return
	}

	if len(user_name) == 0 {
		user_name = utils.RandonString(10)
	}
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(user_passwd), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误！")
		return
	}

	//创建用户
	newUser := model.User{
		Name:      user_name,
		Telephone: user_telephone,
		Password:  string(hasedPassword),
		Email:     user_email,
		Avator:    user_avatar.Filename,
	}
	if DB.HasTable(&newUser) == false{
		DB.CreateTable(&newUser)
	}
	DB.Create(&newUser)
	//返回结果

	response.Success(ctx, nil, "注册成功")
}
func Login(ctx *gin.Context) {
	DB := common.GetDB()
	//获取参数
	user_telephone := ctx.PostForm("telephone")
	user_passwd := ctx.PostForm("passwd")
	//数据验证
	if len(user_telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为十一位！")
		return
	}
	//判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", user_telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnauthorized, 422, nil, "用户不存在！")
		return

	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user_passwd)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误！")
		return
	}
	//发放token前端
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常！"})
		log.Printf("token generate error: %v \n", err)
		return
	}
	response.Success(ctx, gin.H{"code": 200, "token": token, "name": user.Name, "userId": user.ID, "telephone": user.Telephone}, "登录成功")
}
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK,gin.H{"code":200,"user":user})
	//response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.User))}, "登陆成功")
}
func Signout(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "user": user})
}
//"github.com/wdana-cn/ginpro/dto"
