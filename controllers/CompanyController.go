package controllers

import (
	"GoApi/models"
	"github.com/kataras/iris"
)

func In(ctx iris.Context)  {
	name := ctx.Params().Get("name")
	company := models.GetUsers(name)

	// ctx.JSON()
	
	ctx.JSON(iris.Map{
		"messageL": company,
	})
	return
}