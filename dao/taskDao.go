package dao

import (
	"TaskManager/dataSource"
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
