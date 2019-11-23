package dao

import (
	"TaskManager/dataModel"
	"TaskManager/dataSource"
	"errors"
	"xorm.io/xorm"
)

type IProjectDao interface {
}

type ProjectDao struct {
	EngineGroup *xorm.EngineGroup
}

func NewProjectDao() *ProjectDao {
	return &ProjectDao{EngineGroup: dataSource.GetMysqlGroup()}
}

func (dao *ProjectDao) InsertProject(project *dataModel.Project) error {
	var session = dao.EngineGroup.NewSession()
	defer session.Close()

	var err error

	_, err = session.Insert(project)
	if err != nil {
		_ = session.Rollback()
		return errors.New("insert project fail -> " + err.Error())
	}

	err = session.Commit()
	if err != nil {
		_ = session.Rollback()
		return errors.New("commit project insertion fail -> " + err.Error())
	}

	return nil
}

func (dao *TaskDao) UpdateProject(project *dataModel.Project, projectTask *dataModel.ProjectTask) error {
	var session = dao.EngineGroup.NewSession()
	defer session.Close()

	var err error

	projectId := project.Id
	project.Id = 0
	_, err = session.Update(project, &dataModel.Task{Id: projectId})
	if err != nil {
		_ = session.Rollback()
		return errors.New("update project fail -> " + err.Error())
	}

	_, err = session.Update(projectTask, &dataModel.TaskUser{TaskId: projectId})
	if err != nil {
		_ = session.Rollback()
		return errors.New("update project'task fail -> " + err.Error())
	}

	err = session.Commit()
	if err != nil {
		_ = session.Rollback()
		return errors.New("commit project fail -> " + err.Error())
	}

	return nil
}
