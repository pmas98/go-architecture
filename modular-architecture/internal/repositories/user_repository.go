package repositories

import (
	"github.com/pmas98/go-architecture/modular-architecture/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *domain.User) (uint, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (r *UserRepository) FindByUsernameOrEmail(username, email string) (*domain.User, error) {
	var user domain.User
	result := r.db.Where("username = ? OR email = ?", username, email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
