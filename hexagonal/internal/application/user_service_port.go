package application

import "hexagonal/internal/domain"

type UserServicePort interface {
	GetUserByID(id int) (*domain.User, error)
	SaveUser(user *domain.User) error
}
