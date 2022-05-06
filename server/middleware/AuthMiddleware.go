package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/wdana-cn/ginpro/common"

	"github.com/gin-gonic/gin"
)

type UserModel struct {
	ID        uint      `json:"userId"`
	CreatedAt time.Time `json:"time"`
	Name      string    `json:"name"`
	Avator    string    `json:"avator"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足！"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足！"})
			ctx.Abort()
			return
		}
		userID := claims.UserId
		DB := common.GetDB()
		var user UserModel
		//DB.First(&user, userID)
		DB.Table("users").Where("id=?",userID).Scan(&user)
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足！"})
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
//"github.com/wdana-cn/ginpro/model"
