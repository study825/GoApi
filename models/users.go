package models

import (
	"GoApi/database"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model

	ID       int
	Username string
	Password string
}

func (Users) TableName() string {
	return "users"
}

func GetUsers(name string) Users {
	fmt.Printf(name)
	u := Users{}
	database.DB.First(&u)

	return u
}

func CreateUser(username string, password string) (bool, error) {
	user := &Users{Username: username, Password: password}

	if err := database.DB.Create(user).Error; err != nil {
		fmt.Printf("CreateUserErr:%s", err)
		return false, err;
	} else {
		return true, nil;
	}
}

func UserList(companyId int64) Users {
	userList := Users{}

	database.DB.Where("company_id",companyId).Find(&userList)

	return userList
}

