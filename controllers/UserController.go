package controllers

import (
	"GoApi/models"
	"github.com/kataras/iris"
)


func UserLogin(ctx iris.Context) {
	// username := ctx.FormValue("username")
	// password := ctx.FormValue("password")

	// response, status, msg := models.CheckLogin(username,password)
	
}

func Create(ctx iris.Context) {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	res := models.CreateUser(username,password)

	ctx.JSON(iris.Map{
		"messageL": res,
	})

	return
}
