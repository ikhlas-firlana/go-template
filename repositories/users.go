package repositories

import (
	"github.com/ikhlas-firlana/go-template/interfaces"
)

// UserRepository repository for user
type UserRepository struct{}

// Provide constructor
func Provide() interfaces.IUsers {
	return &UserRepository{}
}

