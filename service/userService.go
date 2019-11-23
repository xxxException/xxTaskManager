package service

import (
	"TaskManager/common"
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

func (service *UserService) GetUser(jobId int, userName string, roleId int, departmentId int, page int) ([]*dataModel.User, error) {
	user := &dataModel.User{JobId: jobId, Username: userName}
	pageSize := common.Global.PageSize
	if pageSize == 0 {
		pageSize = 10
	}
	limitStart := (page - 1) * pageSize
	limitEnd := page * pageSize
	return service.userDao.QueryUser(user, roleId, departmentId, limitStart, limitEnd)
}

func (service *UserService) AddUser(user *dataModel.User, roleId int, departmentId int) error {
	return service.userDao.InsertUser(user, roleId, departmentId)
}

func (service *UserService) DeleteUser(ids []int) error {
	users := make([]*dataModel.User, len(ids), len(ids))
	for idx, id := range ids {
		users[idx] = &dataModel.User{Id: id}
	}
	return service.userDao.DeleteUser(users)
}

func (service *UserService) UpdateUser(user *dataModel.User, roleId int, departmentId int) error {
	return service.userDao.UpdateUser(user, roleId, departmentId)
}
