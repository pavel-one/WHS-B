package controllers

import (
	"app/base"
	"app/helpers"
	"app/resources"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	*Controller
}

func (c UserController) User(ctx *gin.Context) {
	token, _ := helpers.GetToken(ctx)
	user, _ := helpers.GetUserWithToken(token, c.DB)
	resource := resources.GetUserResource(&user)

	c.Success(resource, ctx)
}

func (c UserController) SetSetting(ctx *gin.Context) {
	token, _ := helpers.GetToken(ctx)
	user, _ := helpers.GetUserWithToken(token, c.DB)

}

func NewUser(db *base.DB) *UserController {
	controller := Controller{DB: db}

	return &UserController{&controller}
}
