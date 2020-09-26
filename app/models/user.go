package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	Id         int    `json:"id"`
	Login_id   string `json:"login_id"`
	Login_pass string `json:login_pass`
	Benese_id  string `json:"benese_id"`
}
type UserLogin struct {
	Login_id   string `json:"login_id"`
	Login_pass string `json:login_pass`
	Benese_id  string `json:"benese_id"`
}

type UserStamp struct {
	Id          int `json:"id"`
	User_id     int `json:"user_id"`
	Stamp_count int `json:"stamp_count"`
	Stamp_type  int `json:"stamp_type"`
}

type Item struct {
	Id          int `json:"id"`
	Stamp_type  int `json:"stamp_type"`
	Stamp_count int `json:"Stamp_count"`
	Item_name   int `json:"item_name"`
}

type user_stu struct {
	User_id        int `json:"user_id"`
	Japanase_score int `json:"Japanase_score"`
	Society_score  int `json:"society_score"`
	Math_score     int `json:"math_score"`
	Science_score  int `json:"science_score"`
	English_score  int `json:"english_score"`
}

type IncrementSend struct {
	Stamp_type int `json:"stamp_type"`
	User_id    int `json:"user_id"`
}

func (userLogin *UserLogin) CreateUser(db *gorm.DB) {
	result := db.Create(&userLogin)
	if err := result.Error; err != nil {
		fmt.Println("err", err)
		// println(err)
	}
}

func (userLogin *UserLogin) LoginUser(db *gorm.DB) User {
	var get User
	result := db.Where("login_id = ? AND login_pass = ?", userLogin.Login_id, userLogin.Login_pass).Find(&get)
	fmt.Println("result", result)
	if err := result.Error; err != nil {
		fmt.Println("err", err)
		// println(err)
	}
	return get
}

func GetAllUsers(db *gorm.DB) []User {
	var users []User
	db.Find(&users)
	return users
}

func GetStampWithUserId(db *gorm.DB, userId int) []UserStamp {
	var stamps []UserStamp
	db.Where("user_id = ? ", userId).Find(&stamps)
	return stamps
}
func GetAllStamp(db *gorm.DB) []UserStamp {
	var stamps []UserStamp
	db.Find(&stamps)
	return stamps
}

func GetItemWithStampType(db *gorm.DB, Stamp_type int) []Item {
	var items []Item
	db.Where("stamp_type = ? ", Stamp_type).Find(&items)
	return items
}

func (incre *UserStamp) IncrementStamp(db *gorm.DB) {
	db.Model(&incre).Update("stamp_count", incre.Stamp_count+1)
}

func GetAllItem(db *gorm.DB) []Item {
	var items []Item
	db.Find(&items)
	return items
}

func GetUserStudyWithUserId(db *gorm.DB, user_id int) []user_stu {
	var userStudys []user_stu
	db.Where("user_id = ? ", user_id).Find(&userStudys)
	return userStudys
}

func GetAllUserStudy(db *gorm.DB) []user_stu {
	var userStudys []user_stu
	db.Find(&userStudys)
	return userStudys
}

// func (quetion *GetQuetionWithStringTags) CreateQuetionWithTags(db *gorm.DB) {
// 	allTags := GetAllTags(db)
// 	mapTags := map[string]int{}
// 	for _, value := range allTags {
// 		mapTags[value.Tag_name] = value.Id
// 	}
// 	setQuetion := Quetion{
// 		Title:          quetion.Title,
// 		Content:        quetion.Content,
// 		Created_at:     quetion.Created_at,
// 		Modifyed_at:    quetion.Modifyed_at,
// 		Group_id:       quetion.Group_id,
// 		Privated:       quetion.Privated,
// 		Create_user_id: quetion.Create_user_id}
// 	db.Create(&setQuetion)
// 	// var maxQuetion Quetion
// 	// maxCount := maxQuetion.Id + 1
// 	// db.Select("MAX(id)").Find(&maxQuetion)
// 	for _, value := range quetion.Tags {
// 		if _, ok := mapTags[value]; !ok {
// 			// ユニークなtagが追加
// 			setTag := Tag{Tag_name: value}
// 			db.Create(setTag)
// 			mapTags[value] = setTag.Id
// 		}
// 		result := db.Create(&QuetionTags{TargetId: setQuetion.Id, Tag_Id: mapTags[value]})
// 		if err := result.Error; err != nil {
// 			fmt.Println("err", err)
// 			// println(err)
// 		}
// 	}

// }

// func GetAllQuetion(db *gorm.DB) []Quetion {
// 	var quetions []Quetion
// 	db.Find(&quetions)
// 	fmt.Println("確認！！", quetions)
// 	return quetions
// }

// func GetQuetion(db *gorm.DB) Quetion {
// 	var quetion Quetion
// 	db.Find(&quetion)
// 	fmt.Println("確認!!T", quetion)
// 	return quetion
// }

// func (quetion *Quetion) CreateQuetion(db *gorm.DB) {
// 	result := db.Create(&quetion)
// 	if err := result.Error; err != nil {
// 		fmt.Println("err", err)
// 		// println(err)
// 	}
// }
