package dao

import (
	"TaskManager/dataModel"
	"TaskManager/dataSource"
	"container/list"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"xorm.io/xorm"
)

type IUserDao interface {
	InsertUser(*dataModel.User, int) error
	DeleteUser(*dataModel.User) error
	UpdateUser(*dataModel.User) error
	QueryUser(*dataModel.User, int, int) (list.List, error)
}

type UserDao struct {
	EngineGroup *xorm.EngineGroup
}

func NewUserDao() *UserDao {
	return &UserDao{EngineGroup: dataSource.GetMysqlGroup()}
}

func (dao *UserDao) InsertUser(user *dataModel.User, roleId int, departmentId int) error {
	//开启事务
	var session = dao.EngineGroup.NewSession()
	defer session.Close()

	var err error
	//user insert
	_, err = session.Insert(user)
	if err != nil {
		_ = session.Rollback()
		return errors.New("insert user fail ->" + err.Error())
	}

	//添加角色
	var userRole = dataModel.UserRole{
		UserId: user.Id,
		RoleId: roleId,
	}
	_, err = session.Insert(userRole)
	if err != nil {
		_ = session.Rollback()
		return errors.New("when inserting user , insert user'role fail ->" + err.Error())
	}

	//添加部门
	var userDepartment = dataModel.UserDepartment{
		UserId:       user.Id,
		DepartmentId: departmentId,
	}
	_, err = session.Insert(userDepartment)
	if err != nil {
		_ = session.Rollback()
		return errors.New("when inserting user , insert user'department fail ->" + err.Error())
	}

	//事务提交
	return session.Commit()
}

func (dao *UserDao) DeleteUser(user *dataModel.User) error {
	var session = dao.EngineGroup.NewSession()
	defer session.Close()

	var err error
	//受影响行数
	var affected int64

	affected, err = session.ID(user.Id).Delete(user)
	if err != nil {
		_ = session.Rollback()
		return errors.New("Error number of user deleted->" + err.Error())
	}
	if affected != 1 {
		_ = session.Rollback()
		return errors.New("Error number of user deleted-> expected: 1 , actual: " + string(affected))
	}

	affected, err = session.Delete(&dataModel.UserRole{UserId: user.Id})
	if err != nil {
		_ = session.Rollback()
		return errors.New("delete user'department fail -> " + err.Error())
	}
	if affected != 1 {
		_ = session.Rollback()
		return errors.New("delete user'role fail, error number of user_role deleted-> expected: 1 , actual: " + string(affected))
	}

	affected, err = session.Delete(&dataModel.UserDepartment{UserId: user.Id})
	if err != nil {
		_ = session.Rollback()
		return errors.New("delete user'department fail -> " + err.Error())
	}
	if affected != 1 {
		_ = session.Rollback()
		return errors.New("delete user'department fail, error number of user_role deleted-> expected: 1 , actual: " + string(affected))
	}

	return session.Commit()
}

func (dao *UserDao) UpdateUser(user *dataModel.User, roleId int, departmentId int) error {
	var session = dao.EngineGroup.NewSession()
	defer session.Close()

	var err error
	//受影响行数
	var affected int64

	affected, err = session.ID(user.Id).Update(user)
	if err != nil {
		_ = session.Rollback()
		return errors.New("update user fail-> " + err.Error())
	}
	if affected != 1 {
		_ = session.Rollback()
		return errors.New("Error number of user updated-> expected: 1 , actual: " + string(affected))
	}

	//需要更新角色
	if roleId != 0 {
		//update  第一个参数是更新项，只更新不为空的，第二个参数是更新的where条件
		affected, err = session.Update(&dataModel.UserRole{RoleId: roleId}, &dataModel.UserRole{UserId: user.Id})
		if err != nil {
			_ = session.Rollback()
			return errors.New("update user'role fail-> " + err.Error())
		}
		if affected != 1 {
			_ = session.Rollback()
			return errors.New("Error number of user'role updated-> expected: 1 , actual: " + string(affected))
		}
	}

	//需要更新部门
	if departmentId != 0 {
		affected, err = session.Where("").Update(&dataModel.UserDepartment{DepartmentId: departmentId}, &dataModel.UserDepartment{UserId: user.Id})
		if err != nil {
			_ = session.Rollback()
			return errors.New("update user'department fail-> " + err.Error())
		}
		if affected != 1 {
			_ = session.Rollback()
			return errors.New("Error number of user'department updated-> expected: 1 , actual: " + string(affected))
		}
	}

	return session.Commit()
}

func (dao *UserDao) QueryUser(user *dataModel.User, departmentId int, roleId int, limitStart int, limitEnd int) ([]dataModel.User, error) {
	var users = make([]dataModel.User, 0, 10)

	var session = dao.EngineGroup.NewSession()
	defer session.Close()

	//构建where条件
	session = session.Where("id >", 0)
	if user.JobId != 0 {
		session = session.And("jobId=", strconv.Itoa(user.JobId))
	}
	if user.Mail != "" {
		session = session.And("mail=", user.Mail)
	}
	if user.Tel != "" {
		session = session.And("tel=", user.Tel)
	}
	if user.Username != "" {
		session = session.And("jobId=", user.Username)
	}
	if limitStart != -1 && limitEnd != -1 {
		session = session.Limit(limitStart, limitEnd)
	}

	err := session.Find(users)
	if err != nil {
		return nil, err
	}
	return users, err
}
