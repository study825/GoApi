package main

import (
	// "iris/models"
	"GoApi/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/iris-contrib/middleware/cors"

)

func newApp() (api *iris.Application){
	api = iris.New()

	// api.Get("/ping/{name}",func(ctx iris.Context) {
	// 	name := ctx.Params().Get("name")

	// 	ctx.JSON(iris.Map{
	// 		"messageL": name,
	// 	})
	// })
	// api.Get("/users", controllers.index)

	// api.PartyFunc("/users", func(users router.Party) {
	// 	roles.Get("/", controllers.index)
	// })

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})

	v1 := api.Party("/v1", crs).AllowMethods(iris.MethodOptions)
	{
		v1.PartyFunc("/users", func(users router.Party) {
			// 根据账号密码获取access_token
			//users.Get("/access_token", controllers.In)
			users.Post("/create", controllers.Create)

		})
		
	}



	return 
}

func main() {
	app := newApp()

	addr := ":8081";
	app.Run(iris.Addr(addr))
}