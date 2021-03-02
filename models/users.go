package models

// Users struct
type Users struct {
	ID       int    `gorm:"column:id;primary_key"`
	Username string `gorm:"column:username"`
	Name     string `gorm:"column:name"`
}
