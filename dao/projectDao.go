package dao

import (
	"TaskManager/dataSource"
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
