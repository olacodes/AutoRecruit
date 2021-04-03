package main

import (
	"autorecruit/db"
	"autorecruit/db/model"
)



func main() {
	Db := db.InitDB()

	user := model.User{ID:1, CompanyName: "olacodes", Email: "olatundesodiq@gmail.com", Password: "olacodes"}
	
	Db.Create(&user)


}
