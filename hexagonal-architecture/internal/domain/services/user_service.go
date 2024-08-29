package services

import (
	"errors"

	"github.com/pmas98/go-architecture/hexagonal-architecture/internal/domain/entity"
	"github.com/pmas98/go-architecture/hexagonal-architecture/internal/domain/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Register(username, email, password string) (uint, error) {
	existingUser, _ := s.userRepo.FindByUsernameOrEmail(username, email)

	if existingUser != nil {
		return 0, errors.New("user already registered with this username or email")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	user := &entity.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	id, err := s.userRepo.CreateUser(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *UserService) Authenticate(username, password string) (bool, *entity.User, error) {
	user, err := s.userRepo.FindByUsernameOrEmail(username, username)
	if err != nil {
		return false, nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, nil, err
	}
	return true, user, nil
}
