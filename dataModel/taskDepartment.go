package dataModel

type TaskDepartment struct {
	Id           int
	DepartmentId int `xorm:"int 'departmentId'"`
	TaskId       int `xorm:"int 'taskId'"`
}

func (this *TaskDepartment) TableName() string {
	return "TaskDepartment"
}
