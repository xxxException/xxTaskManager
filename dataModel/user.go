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

type UserQueryCondition struct {
	UserName string `json:"user_name"`
	JobId    int    `json:"job_id"`
	Mail     string `json:"mail"`
	Tel      string `json:"tel"`
	Profile  string `json:"profile"`

	RoleId       int   `json:"role_id"`
	DepartmentId int   `json:"department_id"`
	Page         int   `json:"page"`
	Ids          []int `json:"ids"`
}

func (this *User) TableName() string {
	return "user"
}
