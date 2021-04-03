package db

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func InitDB() *gorm.DB {
	cfg, err := ini.Load("my.ini")

	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	dbPassword := cfg.Section("db").Key("db_password").String()
	dbHost := cfg.Section("db").Key("db_host").String()
	dbUser := cfg.Section("db").Key("db_user").String()
	dbName := cfg.Section("db").Key("db_name").String()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", dbHost, dbUser, dbPassword, dbName)
	Db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return Db
}

