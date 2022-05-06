package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wdana-cn/ginpro/common"
	"github.com/wdana-cn/ginpro/model"
	"github.com/wdana-cn/ginpro/response"
)

type UserModel struct {
	ID        uint      `json:"userId"`
	CreatedAt time.Time `json:"time"`
	Name      string    `json:"name"`
	Avator    string    `json:"avator"`
}

type PostsData struct {
	ID        uint      `json:"postId"`
	CreatedAt time.Time `json:"time"`
	Title     string    `json:"title"`
	Postvalue string    `json:"type"`
	Content   string    `json:"content"`
	UserId    uint      `json:"userId"`
}
type CommentData struct {
	ID        uint      `json:"commentId"`
	CreatedAt time.Time `json:"time"`
	Postid    uint      `json:"postId"`
	Comment   string    `json:"comment"`
	Userid    uint      `json:"userId"`
}

func GetPosts(ctx *gin.Context) {
	DB := common.GetDB()
	value := ctx.Param("value")
	var postsdata []PostsData
	if value == "0" {
		DB.Table("postdata").Scan(&postsdata)
		ctx.JSON(http.StatusOK, postsdata)
	} else {
		DB.Table("postdata").Where("postvalue = ?", value).Find(&postsdata)
		ctx.JSON(http.StatusOK, postsdata)
	}
}

func GetProfile(ctx *gin.Context) {
	DB := common.GetDB()
	var usermodel UserModel
	userid := ctx.Param("userId")
	if userid == "" {
		response.Fail(ctx, "ID null", gin.H{"code": 500})
	} else {
		modelId, err := strconv.Atoi(userid)
		if err != nil {
			log.Println(err)
		}
		DB.Table("users").Where("id=?", uint(modelId)).Scan(&usermodel)

		log.Println(usermodel)
		ctx.JSON(http.StatusOK, usermodel)
		//response.Success(ctx, gin.H{"profile": usermodel}, "")
	}
}
func GetPostDetail(ctx *gin.Context) {
	db := common.GetDB()
	pid := ctx.Query("postId")
	var PostDataDetail PostsData
	if pid == "" {
		response.Fail(ctx, "please input postid", gin.H{"post": "null"})
	} else {
		db.Table("postdata").Where("id=?", pid).Scan(&PostDataDetail)
		ctx.JSON(http.StatusOK, PostDataDetail)
		//response.Success(ctx, gin.H{}, "")
	}

}
func AddPosts(ctx *gin.Context) {
	DB := common.GetDB()
	PostValue := ctx.Param("value")
	PostTitle := ctx.PostForm("title")
	PostContent := ctx.PostForm("content")
	postuserid := ctx.PostForm("userId")
	PostUserID, _ := strconv.Atoi(postuserid)
	var postmodel model.Postdata
	postmodel.Title = PostTitle
	postmodel.Postvalue = PostValue
	postmodel.Content = PostContent
	postmodel.UserId = uint(PostUserID)
	if DB.HasTable(&postmodel) == false {
		DB.CreateTable(&postmodel)
	}
	DB.Create(&postmodel)
	//ctx.JSON(http.StatusOK, "addpost done!")
	response.Success(ctx, gin.H{}, "addpost done!")
}
func AddPostComment(ctx *gin.Context) {
	pid := ctx.Param("postid")
	postcomment := ctx.PostForm("comment")
	uid := ctx.PostForm("userId")
	db := common.GetDB()
	var comment model.Comment
	comment.Comment = postcomment
	postid, err1 := strconv.Atoi(pid)
	userid, err2 := strconv.Atoi(uid)
	if err1 != nil {
		log.Println(err1)
	}
	if err2 != nil {
		log.Println(err2)
	}
	comment.Postid = uint(postid)
	comment.Userid = uint(userid)
	log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	log.Println(comment)
	db.Create(&comment)
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "comment": comment})
}
func GetComment(ctx *gin.Context) {
	pid := ctx.Param("postid")
	id, err := strconv.Atoi(pid)
	if err != nil {
		log.Println(err)
	}
	postid := uint(id)
	db := common.GetDB()
	var commentdata []CommentData
	db.Table("comments").Where("postid=?", postid).Scan(&commentdata)
	ctx.JSON(http.StatusOK, commentdata)
	//response.Success(ctx, gin.H{"comementdata": commentdata}, "add comment done!")

}
func GetPostComment(ctx *gin.Context) {
	db := common.GetDB()
	pid := ctx.Param("id")
	pcid, err := strconv.Atoi(pid)
	if err != nil {
		log.Println(err)
	}
	var commentdata []CommentData
	db.Table("comments").Where("postid=?", uint(pcid)).Scan(&commentdata)
	ctx.JSON(http.StatusOK, commentdata)
}
