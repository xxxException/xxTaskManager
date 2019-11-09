package dataModel

type UserDepartment struct {
	Id           int
	UserId       int `xorm:"int 'userId'"`
	DepartmentId int `xorm:"int 'departmentId'"`
}

func (this *UserDepartment) TableName() string {
	return "user_department"
}
