package dao

import (
	"TaskManager/dataModel"
	"TaskManager/dataSource"
	"errors"
	"xorm.io/xorm"
)

type IDepartmentDao interface {
}

type DepartmentDao struct {
	EngineGroup *xorm.EngineGroup
}

func NewDepartmentDao() *DepartmentDao {
	return &DepartmentDao{EngineGroup: dataSource.GetMysqlGroup()}
}

func (dao *DepartmentDao) InsertDepartment(department dataModel.Department) error {
	var session = dao.EngineGroup.NewSession()
	defer session.Close()

	var err error

	_, err = session.Insert(department)
	if err != nil {
		_ = session.Rollback()
		return errors.New("insert department fail -> " + err.Error())
	}

	return session.Commit()
}
