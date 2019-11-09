package dataModel

type Permission struct {
	Id             int
	PermissionName string `xorm:"varchar 'permissionName'"`
	Desc           string `xorm:"varchar 'desc'"`
}

func (this *Permission) TableName() string {
	return "permission"
}
