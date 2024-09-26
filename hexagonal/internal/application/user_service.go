package application

import "hexagonal/internal/domain"

var _ UserServicePort = (*UserService)(nil)

type UserService struct {
	repo domain.UserRepository
}

// GetUserByID implements UserServicePort.
func (u *UserService) GetUserByID(id int) (*domain.User, error) {
	return u.repo.GetByID(id)
}

// SaveUser implements UserServicePort.
func (u *UserService) SaveUser(user *domain.User) error {
	return u.repo.Save(user)
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo}
}
