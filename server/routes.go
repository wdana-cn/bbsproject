package main

import (
	"net/http"

	"github.com/wdana-cn/ginpro/controller"
	"github.com/wdana-cn/ginpro/middleware"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,X-Access-Token")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func CollectRouter(r *gin.Engine) *gin.Engine {
	//User Register
	r.POST("/api/user/register", controller.Register)
	//User Login
	r.POST("/api/user/login", controller.Login)
	//Authorization User status
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	//User Logout
	r.GET("/api/user/signout", middleware.AuthMiddleware(), controller.Signout)
	//Get User Profile
	//r.POST("/api/user/profile/:userId", controller.GetProfile)
	r.GET("/api/user/profile/:userId", controller.GetProfile)
	r.StaticFS("/uploads", http.Dir("./uploads"))
	//Get Post Details
	r.GET("/api/user/postDetail", controller.GetPostDetail)
	//Get/Add  Post+$
	r.GET("/api/user/posts/:value", controller.GetPosts)
	r.POST("/api/user/posts/:value", controller.AddPosts)
	//Get/Add Comments+$
	r.GET("/api/user/comments/:postid", controller.GetComment)
	r.POST("/api/user/comments/:postid", controller.AddPostComment)
	//gePostComment
	r.GET("/api/user/getcomment/:id", controller.GetPostComment)
	return r
}
