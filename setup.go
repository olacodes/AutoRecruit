package main


import (
	"autorecruit/db"
	"autorecruit/db/model"
)


func init() {
	db := db.InitDB()
	db.AutoMigrate(&model.User{})
	// user := model.User{ID:1, CompanyName: "olacodes", Email: "olatundesodiq@gmail.com", Password: "olacodes"}
	// db.Create(&user)

}
