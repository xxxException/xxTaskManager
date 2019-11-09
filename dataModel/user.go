package dataModel

type User struct {
	Id       int    //mysql自动返回的key会填进Id里
	JobId    int    `xorm:"int 'jobId'"`
	Username string `xorm:"varchar 'username'"`
	Password string `xorm:"varchar 'password'"`
	Salt     string `xorm:"varchar 'salt'"`
	Mail     string `xorm:"varchar 'mail'"`
	Tel      string `xorm:"varchar 'tel'"`
	Profile  string `xorm:"varchar 'profile'"`

	DepartmentName string
	RoleName       string
}

func (this *User) TableName() string {
	return "user"
}
