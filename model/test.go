package model

type test struct {
	Name     string `gorm:"primary_key"`
	Password string
}

func (test) TableName() string {
	return "test"
}
