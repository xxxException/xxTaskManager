package dao

import (
	"TaskManager/dataSource"
	"xorm.io/xorm"
)

type IRoleDao interface {
}

type RoleDao struct {
	EngineGroup *xorm.EngineGroup
}

func NewRoleDao() *RoleDao {
	return &RoleDao{EngineGroup: dataSource.GetMysqlGroup()}
}
