package auth

import "gorm.io/gorm"

type AuthRepository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
	FindByID(id uint) (*User, error)
	Update(user *User) error
}

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{DB: db}
}

// Create implements AuthRepository.
func (a *authRepository) Create(user *User) error {
	return a.DB.Create(user).Error
}

// FindByEmail implements AuthRepository.
func (a *authRepository) FindByEmail(email string) (*User, error) {
	var user User
	err := a.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

// FindByID implements AuthRepository.
func (a *authRepository) FindByID(id uint) (*User, error) {
	var user User
	err := a.DB.First(&user, id).Error
	return &user, err
}

// Update implements AuthRepository.
func (a *authRepository) Update(user *User) error {
	return a.DB.Save(user).Error
}
