package cases

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pmas98/go-architecture/modular-architecture/internal/domain"
	"github.com/pmas98/go-architecture/modular-architecture/internal/services"
)

type UserUseCase struct {
	userService *services.UserService
	secretKey   string
}

type CustomClaims struct {
	UserID uint `json:"userID"`
	jwt.RegisteredClaims
}

func NewUserUseCase(userService *services.UserService, secretKey string) *UserUseCase {
	return &UserUseCase{
		userService: userService,
		secretKey:   secretKey,
	}
}

func (uc *UserUseCase) AuthenticateAndGenerateTokens(username, password string) (string, string, error) {
	isAuthenticated, user, err := uc.userService.Authenticate(username, password)
	if err != nil || !isAuthenticated {
		return "", "", errors.New("invalid credentials")
	}

	accessToken, err := uc.generateToken(username, user, 15*time.Minute)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := uc.generateToken(username, user, 24*time.Hour)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (uc *UserUseCase) generateToken(username string, user *domain.User, duration time.Duration) (string, error) {
	expirationTime := time.Now().Add(duration)

	fmt.Println(user.ID)

	claims := &CustomClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   username,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Create a new token with the custom claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte(uc.secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (uc *UserUseCase) Register(username, email, password string) (uint, error) {
	if username == "" || email == "" || password == "" {
		return 0, errors.New("all fields are required")
	}

	userID, err := uc.userService.Register(username, email, password)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
