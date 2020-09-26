package handler

import (
	"app/database"
	"app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// func ApiGetAllQuetions(c *gin.Context) {
// 	posts := models.GetAllQuetion(database.GetDB())
// 	println("確認！！AA", posts)
// 	c.JSON(200, posts)
// }

// func ApiRegistQuetion(c *gin.Context) {
// 	req := &models.Quetion{}
// 	err := c.BindJSON(req)
// 	if err != nil {
// 		println(err)
// 	}
// 	println("req", req)
// 	req.CreateQuetion(database.GetDB())
// 	c.JSON(200, req)
// }

// func ApiGetQuetion(c *gin.Context) {
// 	posts := models.GetQuetion(database.GetDB())
// 	c.JSON(200, posts)
// }

func ApiRegistUser(c *gin.Context) {
	req := &models.UserLogin{}
	err := c.BindJSON(req)
	if err != nil {
		println(err)
	}
	println("req", req)
	req.CreateUser(database.GetDB())
	c.JSON(200, req)
}

func ApiLoginUser(c *gin.Context) {
	req := &models.UserLogin{}
	err := c.BindJSON(req)
	if err != nil {
		println(err)
	}
	println("req", req)
	get := req.LoginUser(database.GetDB())
	c.JSON(200, get)
}

func ApiGetAllUser(c *gin.Context) {
	posts := models.GetAllUsers(database.GetDB())
	println("確認！！AA", posts)
	c.JSON(200, posts)
}

func ApiGetStampWithUserId(c *gin.Context) {
	param := c.Param("user_id")
	intParam, _ := strconv.Atoi(param)
	posts := models.GetStampWithUserId(database.GetDB(), intParam)
	println("確認！！AA", posts)
	c.JSON(200, posts)
}

func ApiGetAllStamp(c *gin.Context) {
	posts := models.GetAllStamp(database.GetDB())
	println("確認！！AA", posts)
	c.JSON(200, posts)
}

func ApiIncrementStamp(c *gin.Context) {
	req := &models.UserStamp{}
	err := c.BindJSON(req)
	if err != nil {
		println(err)
	}
	println("req", req)
	req.IncrementStamp(database.GetDB())
	c.JSON(200, req)
}

func ApiGetItemWithStampType(c *gin.Context) {
	param := c.Param("stamp_type")
	intParam, _ := strconv.Atoi(param)
	posts := models.GetItemWithStampType(database.GetDB(), intParam)
	println("確認！！AA", posts)
	c.JSON(200, posts)
}

func ApiGetAllItem(c *gin.Context) {
	posts := models.GetAllItem(database.GetDB())
	println("確認！！AA", posts)
	c.JSON(200, posts)
}

func ApiGetUserStudyWithUserId(c *gin.Context) {
	param := c.Param("user_id")
	intParam, _ := strconv.Atoi(param)
	posts := models.GetUserStudyWithUserId(database.GetDB(), intParam)
	println("確認！！AA", posts)
	c.JSON(200, posts)
}

func ApiGetAllUserStudy(c *gin.Context) {
	posts := models.GetAllUserStudy(database.GetDB())
	println("確認！！AA", posts)
	c.JSON(200, posts)
}
