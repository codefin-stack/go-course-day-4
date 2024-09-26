package db

import (
	"fmt"
	"hexagonal/internal/domain"

	"gorm.io/gorm"
)

var _ domain.UserRepository = (*DBUserRepository)(nil)

type DBUserRepository struct {
	db *gorm.DB
}

// GetByID implements domain.UserRepository.
func (d *DBUserRepository) GetByID(id int) (*domain.User, error) {
	var user domain.User
	result := d.db.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, result.Error
	}
	return &user, nil
}

func NewDBUserRepository(db *gorm.DB) *DBUserRepository {
	return &DBUserRepository{db: db}
}

// Save implements domain.UserRepository.
func (d *DBUserRepository) Save(user *domain.User) error {
	result := d.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
