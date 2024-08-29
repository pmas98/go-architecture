package repositories

import (
	"github.com/pmas98/go-architecture/hexagonal-architecture/internal/domain/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *entity.User) (uint, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (r *UserRepository) FindByUsernameOrEmail(username, email string) (*entity.User, error) {
	var user entity.User
	result := r.db.Where("username = ? OR email = ?", username, email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
