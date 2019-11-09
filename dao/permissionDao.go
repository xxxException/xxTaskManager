package dao

import (
	"TaskManager/dataSource"
	"xorm.io/xorm"
)

type IPermissionDao interface {
}

type PermissionDao struct {
	EngineGroup *xorm.EngineGroup
}

func NewPermissionDao() *PermissionDao {
	return &PermissionDao{EngineGroup: dataSource.GetMysqlGroup()}
}
