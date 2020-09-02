package domain

import (
	"time"
)

type Users []User

type User struct {
	// gorm.Model
	ID        int       `gorm:"id" json:"id"`
	UserName  string    `gorm:"column:user_name" json:"user_name"`
	Email     string    `gorm:"column:email" json:"email"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
	// Reservation Reservation `gorm:"foreignKey: UserId"`
}
