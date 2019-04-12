package database

import (
	"GoApi/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pelletier/go-toml"
)

var (
	DB = New()
)

/**
*设置数据库连接
*@param diver string
 */
func New() *gorm.DB {
		driver := config.Conf.Get("database.dirver").(string)
		configTree := config.Conf.Get(driver).(*toml.Tree)
		userName := configTree.Get("databaseUserName").(string)
		password := configTree.Get("databasePassword").(string)
		databaseName := configTree.Get("databaseName").(string)

		connect := userName + ":" + password + "@tcp(127.0.0.1:3306)/" + databaseName + "?charset=utf8&parseTime=True&loc=Local"

		DB, err := gorm.Open(driver, connect)
		if err != nil {
			panic(fmt.Sprintf("No error should happen when connecting to  database, but got err=%+v", err))
		}

		return DB
}

