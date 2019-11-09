package dataModel

type UserRole struct {
	Id     int
	UserId int `xorm:"int 'userId'"`
	RoleId int `xorm:"int 'roleId'"`
}

func (this *UserRole) TableName() string {
	return "user_role"
}
