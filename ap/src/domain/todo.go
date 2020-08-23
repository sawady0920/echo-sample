package domain

type Todos []Todo

type Todo struct {
	ID        int
	Title string
	Body  string
}