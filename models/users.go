package models

// Users struct
type Users struct {
	Username string `gorm:"column:username"`
	Name     string `gorm:"column:name"`
}
