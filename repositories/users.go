package repositories

import (
	"github.com/ikhlas-firlana/go-template/interfaces"
	"github.com/ikhlas-firlana/go-template/models"
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

// GetUserByUsername is getname
func (u *UserRepository) GetUserByUsername(username string) (user models.Users, err error) {
	err = u.Db.Table("users").Where("username = ?", username).First(&user).Error
	return
}
