package auth

import "gorm.io/gorm"

type User struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	Name     string  `json:"name"`
	Email    string  `gorm:"unique" json:"email"`
	Password string  `json:"-"`
	Role     string  `gorm:"default:siswa" json:"role"`
	PhotoURL *string `gorm:"photo_url,omitempty"`

	gorm.Model
}
