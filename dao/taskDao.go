package dao

import (
	"TaskManager/dataModel"
	"TaskManager/dataSource"
	"errors"
	"xorm.io/xorm"
)

type ITaskDao interface {
}

type TaskDao struct {
	EngineGroup *xorm.EngineGroup
}

func NewTaskDao() *TaskDao {
	return &TaskDao{EngineGroup: dataSource.GetMysqlGroup()}
}

func (dao *TaskDao) InsertTask(task *dataModel.Task) error {
	//开启事务
	var session = dao.EngineGroup.NewSession()
	defer session.Close()

	var err error
	_, err = session.Insert(task)
	if err != nil {
		_ = session.Rollback()
		return errors.New("insert task fail -> " + err.Error())
	}
	err = session.Commit()
	if err != nil {
		_ = session.Rollback()
		return errors.New("insert task fail -> " + err.Error())
	}
	return nil
}

func (dao *TaskDao) InsertTaskDepartments(taskDepartments []*dataModel.TaskDepartment) error {
	//开启事务
	var session = dao.EngineGroup.NewSession()
	defer session.Close()

	var err error
	for taskDepartment := range taskDepartments {
		_, err = session.Insert(taskDepartment)
		if err != nil {
			_ = session.Rollback()
			return errors.New("insert task'department fail -> " + err.Error())
		}
	}

	err = session.Commit()
	if err != nil {
		_ = session.Rollback()
		return errors.New("insert task'department fail -> " + err.Error())
	}

	return nil
}

func (dao *TaskDao) InsertTaskUsers(taskUsers []*dataModel.TaskUser) error {
	//开启事务
	var session = dao.EngineGroup.NewSession()
	defer session.Close()

	var err error
	for taskUser := range taskUsers {
		_, err = session.Insert(taskUser)
		if err != nil {
			_ = session.Rollback()
			return errors.New("insert task'user fail -> " + err.Error())
		}
	}

	err = session.Commit()
	if err != nil {
		_ = session.Rollback()
		return errors.New("insert task'user fail -> " + err.Error())
	}

	return nil
}

func (dao *TaskDao) DeleteTask(taskId int) error {
	var session = dao.EngineGroup.NewSession()
	defer session.Close()

	var err error

	_, err = session.ID(taskId).Delete(&dataModel.Task{})
	if err != nil {
		_ = session.Rollback()
		return errors.New("delete task fail -> " + err.Error())
	}

	_, err = session.Where("taskId", taskId).Delete(&dataModel.TaskUser{})
	if err != nil {
		_ = session.Rollback()
		return errors.New("delete task'user fail -> " + err.Error())
	}

	_, err = session.Where("taskId", taskId).Delete(&dataModel.TaskDepartment{})
	if err != nil {
		_ = session.Rollback()
		return errors.New("delete task'department fail -> " + err.Error())
	}

	err = session.Commit()
	if err != nil {
		_ = session.Rollback()
		return errors.New("commit task fail -> " + err.Error())
	}

	return nil
}

func (dao *TaskDao) UpdateTask(task *dataModel.Task, taskUser *dataModel.TaskUser,
	taskDepartment *dataModel.TaskDepartment) error {
	var session = dao.EngineGroup.NewSession()
	defer session.Close()

	var err error

	taskId := task.Id
	task.Id = 0
	_, err = session.Update(task, &dataModel.Task{Id: taskId})
	if err != nil {
		_ = session.Rollback()
		return errors.New("update task fail -> " + err.Error())
	}

	_, err = session.Update(taskUser, &dataModel.TaskUser{TaskId: taskId})
	if err != nil {
		_ = session.Rollback()
		return errors.New("update task'user fail -> " + err.Error())
	}

	_, err = session.Update(taskDepartment, &dataModel.TaskDepartment{TaskId: taskId})
	if err != nil {
		_ = session.Rollback()
		return errors.New("update task'department fail -> " + err.Error())
	}

	err = session.Commit()
	if err != nil {
		_ = session.Rollback()
		return errors.New("commit task fail -> " + err.Error())
	}

	return nil
}
