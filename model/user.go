package model

type User struct {
	Id       int
	Username string
	Password string
}

func (User) TableName() string {
	return "user"
}

func (u *User) Register() string {

	return ""
}
