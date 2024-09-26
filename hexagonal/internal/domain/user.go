package domain

type User struct {
	ID    int
	Name  string
	Email string
}

type UserRepository interface {
	GetByID(id int) (*User, error)
	Save(user *User) error
}
