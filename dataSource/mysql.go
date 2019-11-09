package dataSource

import (
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

var mysqlGroup *xorm.EngineGroup

func NewMysqlGroup() error {
	var driverName = "mysql"
	var dataSourceName = "root:199762@tcp(127.0.0.1:3306)/task_manager"
	var err error

	var dataSourceSlice = []string{dataSourceName, dataSourceName, dataSourceName, dataSourceName}
	mysqlGroup, err = xorm.NewEngineGroup(driverName, dataSourceSlice)
	if err != nil {
		return err
	}
	return nil
}

func GetMysqlGroup() *xorm.EngineGroup {
	return mysqlGroup
}
