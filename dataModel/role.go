package dataModel

type Role struct {
	Id       int
	RoleName string `xorm:"varchar 'roleName'"`
	Desc     string `xorm:"varchar 'desc'"`
}

func (this *Role) TableName() string {
	return "role"
}
