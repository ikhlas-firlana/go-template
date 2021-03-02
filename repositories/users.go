package repositories

import (
	"github.com/ikhlas-firlana/go-template/interfaces"
	"github.com/jinzhu/gorm"
)

// UserRepository repository for user
type UserRepository struct {
	Db *gorm.DB
}

// Provide constructor
func Provide(Db *gorm.DB) interfaces.IUsers {
	return &UserRepository{
		Db: Db,
	}
}

// GetName is getname
func (u *UserRepository) GetName(username string) (name string, err error) {
	err = u.Db.Table("users").Where("username = ?", username).First(&name).Error
	return
}
