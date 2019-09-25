package controllers

import (
	"../models"
	"github.com/kataras/iris"
)

// GetUser By Id
func GetUser(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	user := models.GetUserByID(id)
	if user == nil {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(user)
}
