package dataModel

type Task struct {
	Id               int
	Name             string `xorm:"varchar 'name'"`
	CreatedTime      string `xorm:"datetime 'createdTime'"`
	StatusChangeTime string `xorm:"varchar 'statusChangeTime'"`
	Status           string `xorm:"varchar 'roleName'"`
	Note             string `xorm:"varchar 'note'"`
}

func (this *Task) TableName() string {
	return "task"
}
