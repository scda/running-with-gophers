package controllers

import (
	"github.com/kataras/iris"
)

// GetHealth returns the application status
func GetHealth(ctx iris.Context) {
	ctx.StatusCode(iris.StatusNoContent)
}
