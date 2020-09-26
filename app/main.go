package main

import (
	"app/database"
	"app/handler"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Origin",
			"Content-Type",
			"Content-Length",
		},
		AllowOrigins: []string{
			"*",
		},
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello, World")
	})
	r.GET("/api/user", handler.ApiGetAllUser)
	r.POST("/api/login", handler.ApiLoginUser)
	r.POST("/api/regist", handler.ApiRegistUser)
	r.GET("/api/user_stamp/:user_id", handler.ApiGetStampWithUserId)
	r.GET("/api/user_stamp", handler.ApiGetAllStamp)
	r.POST("/api/increment", handler.ApiIncrementStamp)
	r.GET("/api/item/:stamp_type", handler.ApiGetItemWithStampType)
	r.GET("/api/item", handler.ApiGetAllItem)
	r.GET("/api/user_study/:user_id", handler.ApiGetUserStudyWithUserId)
	r.GET("/api/user_study", handler.ApiGetAllUserStudy)
	// r.GET("/posts", GetPost)
	// r.GET("/quetions/:id", handler.ApiGetQuetion)
	// r.POST("create/quetion", handler.ApiRegistQuetion)
	r.Run()
}

type Post struct {
	id      int
	title   string
	content string
}

func GetPost(c *gin.Context) {
	var posts []Post
	db := database.GetDB()
	db.Find(&posts)
	fmt.Println("確認！！", posts)
	c.JSON(200, posts)
}
