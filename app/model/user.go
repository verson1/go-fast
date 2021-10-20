package model

type User struct {
	Id       string `gorm:"column:id" json:"id"`
	User     string `gorm:"column:user" json:"user"`
	Password string `gorm:"column:password" json:"password"`
	Role     int    `gorm:"column:role" json:"role"`
}

func (u *User) TableName() string {
	return "user"
}
