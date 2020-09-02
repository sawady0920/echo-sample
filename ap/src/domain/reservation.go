package domain

import (
	"time"
)

type Reservation struct {
	// gorm.Model
	ID        uint       `gorm:"column:id" json:"id"`
	UserId    uint       `gorm:"column:user_id" json:"userId"`
	StartsAt  *time.Time `gorm:"column:starts_at" json:"startsAt"`
	EndsAt    *time.Time `gorm:"column:ends_at" json:"endsAt"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deletedAt"`
	// User      User       `gorm:"foreignKey:UserId" json:"user"`
}

func (Reservation) TableName() string { return "reservations" }
