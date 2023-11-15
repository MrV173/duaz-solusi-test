package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey:autoIncrement"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	FullName string `json:"fullname" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Gender   string `json:"gender" gorm:"type: varchar(255)"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"fullname"`
}

func (UserResponse) TableName() string {
	return "users"
}