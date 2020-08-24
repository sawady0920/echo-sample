package domain

type Todos []Todo

type Todo struct {
	ID        int `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}