package controller

import (
	"TaskManager/common"
	"TaskManager/dataModel"
	"TaskManager/service"
	"github.com/kataras/iris"
)

type IUserController interface {
}

type UserController struct {
	Ctx     iris.Context
	Service service.UserService
}

func (this *UserController) Get() {

}

func (this *UserController) GetQueryusers() {
	condition := &dataModel.UserQueryCondition{}
	var err error

	err = this.Ctx.ReadJSON(condition)
	if err != nil {
		//错误信息
		this.Ctx.StatusCode(iris.StatusBadRequest)
		this.Ctx.Next()
		return
	}

	users, err := this.Service.GetUser(condition.JobId, condition.UserName,
		condition.RoleId, condition.DepartmentId, condition.Page)
	if err != nil {
		print("ERROR: ", err.Error())
		this.Ctx.StatusCode(iris.StatusBadRequest)
		this.Ctx.Next()
		return
	}

	this.Ctx.StatusCode(iris.StatusOK)
	_, _ = this.Ctx.JSON(ApiResource(true, "success", users))
}

func (this *UserController) DeleteDeleteuser() {
	condition := &dataModel.UserQueryCondition{}
	var err error

	err = this.Ctx.ReadJSON(condition)
	if err != nil {
		//错误信息
		this.Ctx.StatusCode(iris.StatusBadRequest)
		this.Ctx.Next()
		return
	}

	err = this.Service.DeleteUser(condition.Ids)
	if err != nil {
		print("ERROR: ", err.Error())
	}
	this.Ctx.StatusCode(iris.StatusOK)
	_, _ = this.Ctx.JSON(ApiResource(true, "success", nil))
}

func (this *UserController) Putuser() {
	condition := &dataModel.UserQueryCondition{}
	var err error

	err = this.Ctx.ReadJSON(condition)
	if err != nil {
		//错误信息
		this.Ctx.StatusCode(iris.StatusBadRequest)
		this.Ctx.Next()
		return
	}
}
