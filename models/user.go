package models

type User struct {
	Id        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Username  string `json:"username" gorm:"unique;not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password"`
	CreateAt  string `json:"create_at"`
	UpdatedAt string `json:"updated_at"`
}
