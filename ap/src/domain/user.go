package domain

type Users []User

type User struct {
	ID        int
	UserName string
	Email  string
}