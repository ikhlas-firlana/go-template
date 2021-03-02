package interfaces

import "github.com/ikhlas-firlana/go-template/models"

// IUsers interface model users
type IUsers interface {
	GetUserByUsername(username string) (models.Users, error)
}
