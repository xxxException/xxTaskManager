package service

import (
	"TaskManager/dao"
	"TaskManager/dataModel"
)

type IUserService interface {
	GetUser(*dataModel.User, int, int, int, int) ([]dataModel.User, error)
	AddUser(*dataModel.User, int, int) error
	DeleteUser(*dataModel.User) error
	UpdateUser(dataModel.User, int, int) error
}

type UserService struct {
	userDao *dao.UserDao
}

func NewNoteService() *UserService {
	return &UserService{userDao: dao.NewUserDao()}
}

func (service *UserService) GetUser(user *dataModel.User, roleId int, departmentId int, limitStart int,
	limitEnd int) ([]dataModel.User, error) {
	return service.userDao.QueryUser(user, roleId, departmentId, limitStart, limitEnd)
}

func (service *UserService) AddUser(user *dataModel.User, roleId int, departmentId int) error {
	return service.userDao.InsertUser(user, roleId, departmentId)
}

func (service *UserService) DeleteUser(user *dataModel.User) error {
	return service.userDao.DeleteUser(user)
}

func (service *UserService) UpdateUser(user dataModel.User, roleId int, departmentId int) error {
	return service.userDao.UpdateUser(user, roleId, departmentId)
}
