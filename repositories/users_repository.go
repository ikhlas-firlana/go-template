package repositories

import (
	"github.com/ikhlas-firlana/go-template/interfaces"
	"github.com/ikhlas-firlana/go-template/models"
	"github.com/jinzhu/gorm"
)

// UsersRepository repository for user
type UsersRepository struct {
	Db *gorm.DB
}

// Provide constructor
func Provide(Db *gorm.DB) interfaces.IUsers {
	return &UsersRepository{
		Db: Db,
	}
}

// GetUserByUsername is getname
func (u *UsersRepository) GetUserByUsername(username string) (user models.Users, err error) {

	err = u.Db.Table("users").Where("username = ?", username).First(&user).Error
	return
}
