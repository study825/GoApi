package models

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"time"
	"iris/database"
)


type Users struct {
	gorm.Model

	ID	int
	ClientId int 
	ClientSecret string
	Name string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (Users) TableName() string {
	return "users"
}

func GetUsers(name string) Users {
	fmt.Printf(name)
	u := Users{}
	database.DB.First(&u)
	// if err := database.DB.Where("name = ?", 234).First(&u).Error; err != nil {
	// 	fmt.Printf("UserAdminCheckLoginErr:%s", err)
	// }

	return u
}