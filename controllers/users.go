package controllers

import (
	"gopkg.in/go-playground/validator.v9"
	"log"

	"../models"
	"github.com/kataras/iris"
)

// GetUsers receives a list of user entries
func GetUsers(ctx iris.Context) {
	users := models.GetUsers()

	_, _ = ctx.JSON(users)
}

// GetUser By Id
func GetUser(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	user := models.GetUserByID(id)
	if user == nil {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(user)
}

// CreateUser creates a new user entry
func CreateUser(ctx iris.Context) {
	user := new(models.User)

	if err := ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(errorData(err))
	} else {
		err := validate.Struct(user)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			for _, err := range err.(validator.ValidationErrors) {
				log.Println()
				log.Println(err.Namespace())
				log.Println(err.Field())
				log.Println(err.Type())
				log.Println(err.Param())
				log.Println()
			}
		} else {
			log.Println("name:", user.Name)
			log.Println("mail:", user.Email)
			createdUser := models.CreateUser(user)

			ctx.StatusCode(iris.StatusOK)
			_, _ = ctx.JSON(createdUser)
		}
	}

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(user)
}
