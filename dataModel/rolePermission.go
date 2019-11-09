package dataModel

type RolePermission struct {
	Id           int
	RoleId       int `xorm:"int 'RoleId'"`
	PermissionId int `xorm:"int 'permissionId'"`
}

func (this *RolePermission) TableName() string {
	return "role_permission"
}
