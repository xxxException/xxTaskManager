package dataModel

type ProjectTask struct {
	Id        int
	ProjectId int `xorm:"int 'projectId'"`
	TaskId    int `xorm:"int 'taskId'"`
}

func (this *ProjectTask) TableName() string {
	return "project_task"
}
