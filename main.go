package main

import (
	"./controllers"
	"./database"
	"./models"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/middleware/logger"
)

func initData() {
	database.Connect()
	models.Initialize()
}

func initApp() (api *iris.Application) {
	api = iris.New()
	api.Use(logger.New())

	api.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.WriteString("404 not found")
	})

	api.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.WriteString("Something went wrong. Try again later.")
	})

	iris.RegisterOnInterrupt(func() {
		database.Disconnect()
	})

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		ExposedHeaders:   []string{"Accept", "Accept-Encoding", "Authorization", "Content-Length", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})

	v1 := api.Party("/_api/v1", crs).AllowMethods(iris.MethodOptions)
	{
		v1.Get("/health", controllers.GetHealth)
		v1.PartyFunc("/users", func(users router.Party) {
			//TODO : ENABLE users.Get("/", controllers.GetAllUsers)
			users.Get("/{id:uint}", controllers.GetUser)
			//TODO : ENABLE users.Post("/", controllers.CreateUser)
			//TODO : ENABLE users.Put("/{id:uint}", controllers.UpdateUser)
			//TODO : ENABLE users.Delete("/{id:uint}", controllers.DeleteUser)
		})
	}

	return
}

func main() {

	app := initApp()
	app.Run(iris.Addr(":8080"))
}
