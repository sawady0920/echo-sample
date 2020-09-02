package domain

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Todos []Todo

type Todo struct {
	gorm.Model
	ID        int       `gorm:"column:id" json:"id"`
	Title     string    `gorm:"column:title" json:"title"`
	Body      string    `gorm:"column:body" json:"body"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
	UserId    uint      `gorm:"column:user_id" json:"userId"`
}
