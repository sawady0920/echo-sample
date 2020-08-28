package domain

import "time"

type Users []User

type User struct {
	ID        int
	UserName  string
	Email     string
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
