package model

import (
	"time"
)

type User struct {
	ID          int    `gorm:"primaryKey"`
	CompanyName string `gorm:"unique" gorm:"not null"`
	Email       string `gorm:"unique" gorm:"not null"`
	Password    string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
