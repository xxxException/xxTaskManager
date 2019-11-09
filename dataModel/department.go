package dataModel

type Department struct {
	Id   int
	Name string `xorm:"varchar 'name'"`
	Desc string `xorm:"varchar 'desc'"`
}

func (this *Department) TableName() string {
	return "department"
}
