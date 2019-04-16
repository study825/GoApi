package controllers

import (
	"GoApi/models"
	"github.com/kataras/iris"
)

func CreateUser(ctx iris.Context) {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	res, err := models.CreateUser(username,password)

	if res != false {
		ctx.JSON(iris.Map{
			"message": "success",
		})
	}else {
		ctx.JSON(iris.Map{
			"message": err,
		})
	}
	
	return
}

func UserList(ctx iris.Context)  {
	companyId,_ := ctx.Params().GetInt64("company_id")

	res := models.UserList(companyId)

	ctx.JSON(iris.Map{
		"data": res,
	})
	return
}

func UpdateUser(ctx iris.Context) {

}