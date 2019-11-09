package dataModel

type TaskUser struct {
	Id     int
	UserId int `xorm:"int 'userId'"`
	TaskId int `xorm:"int 'taskId'"`
}

func (this *TaskUser) TableName() string {
	return "task_user"
}
