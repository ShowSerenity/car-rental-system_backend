package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email,omitempty" gorm:"unique"`
	Password string `json:"password"`
	FullName string `json:"full_name,omitempty"`
}
