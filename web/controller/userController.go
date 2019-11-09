package controller

import (
	"TaskManager/service"
	"github.com/kataras/iris"
)

type IUserController interface {

}

type UserController struct {
	Ctx iris.Context
	Service service.UserService
}

func (this *UserController) Get() {

}

func (this *UserController) GetQueryusers() {

}

func (this *UserController)
