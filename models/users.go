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
	// if err := database.DB.Where("name = ?", 234).First(&u).Error; err != nil {
	// 	fmt.Printf("UserAdminCheckLoginErr:%s", err)
	// }

	return u
}

func CreateUser(username string, password string) (bool, error) {
	user := Users{Username: username, Password: password}

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
/**
 * 判断用户是否登录
 * @method CheckLogin
 * @param  {[type]}  id       int    [description]
 * @param  {[type]}  password string [description]
 */
//  func CheckLogin(username, password string) (response Token, status bool, msg string) {
// 	user := UserAdminCheckLogin(username)
// 	if user.ID == 0 {
// 		msg = "用户不存在"
// 		return
// 	} else {
// 		if ok := bcrypt.Match(password, user.Password); ok {
// 			token := jwt.New(jwt.SigningMethodHS256)
// 			claims := make(jwt.MapClaims)
// 			claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
// 			claims["iat"] = time.Now().Unix()
// 			token.Claims = claims
// 			tokenString, err := token.SignedString([]byte("secret"))

// 			if err != nil {
// 				msg = err.Error()
// 				return
// 			}

// 			oauth_token := new(OauthToken)
// 			oauth_token.Token = tokenString
// 			oauth_token.UserId = user.ID
// 			oauth_token.Secret = "secret"
// 			oauth_token.Revoked = false
// 			oauth_token.ExpressIn = time.Now().Add(time.Hour * time.Duration(1)).Unix()
// 			oauth_token.CreatedAt = time.Now()

// 			response = oauth_token.OauthTokenCreate()
// 			status = true
// 			msg = "登陆成功"

// 			return

// 		} else {
// 			msg = "用户名或密码错误"
// 			return
// 		}
// 	}
// }

/**
 * 校验用户登录
 * @method UserAdminCheckLogin
 * @param  {[type]}       username string [description]
 */
func UserAdminCheckLogin(username string) Users {
	u := Users{}
	if err := database.DB.Where("username = ?", username).First(&u).Error; err != nil {
		fmt.Printf("UserAdminCheckLoginErr:%s", err)
	}
	return u
}
