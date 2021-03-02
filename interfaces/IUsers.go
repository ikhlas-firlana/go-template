package interfaces

// IUsers interface model users
type IUsers interface {
	GetName(username string) (string, error)
}
