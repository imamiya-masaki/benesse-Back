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
	r.GET("/user", handler.ApiGetAllUser)
	r.POST("/login", handler.ApiLoginUser)
	r.POST("/regist", handler.ApiRegistUser)
	r.GET("/user_stamp/:user_id", handler.ApiGetStampWithUserId)
	r.GET("/user_stamp", handler.ApiGetAllStamp)
	r.POST("/increment", handler.ApiIncrementStamp)
	r.GET("/item/:stamp_type", handler.ApiGetItemWithStampType)
	r.GET("/item", handler.ApiGetAllItem)
	r.GET("/user_study/:user_id", handler.ApiGetUserStudyWithUserId)
	r.GET("/user_study", handler.ApiGetAllUserStudy)
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
