package auth

import (
	"errors"

	"github.com/ihsankarim/backend-brighted/pkg/utils"
)

type AuthService interface {
	Register(req RegisterRequest) (*User, error)
	Login(req LoginRequest) (*User, error)
	GetProfile(id uint) (*User, error)
	UpdateProfile(id uint, name string, photoURL *string) error
}

type authService struct {
	repo AuthRepository
}

func NewAuthService(r AuthRepository) AuthService {
	return &authService{repo: r}
}

// GetProfile implements AuthService.
func (a *authService) GetProfile(id uint) (*User, error) {
	return a.repo.FindByID(id)
}

// Login implements AuthService.
func (a *authService) Login(req LoginRequest) (*User, error) {
	user, err := a.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("email or password invalid")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("email or password invalid")
	}
	return user, nil
}

// Register implements AuthService.
func (a *authService) Register(req RegisterRequest) (*User, error) {
	hashedPassword, _ := utils.HashPassword(req.Password)
	user := &User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "siswa",
	}

	if err := a.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateProfile implements AuthService.
func (a *authService) UpdateProfile(id uint, name string, photoURL *string) error {
	user, err := a.repo.FindByID(id)
	if err != nil {
		return err
	}
	user.Name = name
	user.PhotoURL = photoURL
	return a.repo.Update(user)
}
