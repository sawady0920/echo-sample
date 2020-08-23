package usecase

import (
	"echo-sample/domain"
)

type TodoRepository interface {
	FindById(id int) (domain.Todo, error)
	FindAll() (domain.Todos, error)
	Store(domain.Todo) (domain.Todo, error)
	Update(domain.Todo) (domain.Todo, error)
	DeleteById(domain.Todo) error
}

type TodoInteractor struct {
	TodoRepository TodoRepository
}

func (interactor *TodoInteractor) TodoById(id int) (user domain.Todo, err error) {
	user, err = interactor.TodoRepository.FindById(id)
	return
}

func (interactor *TodoInteractor) Todos() (users domain.Todos, err error) {
	users, err = interactor.TodoRepository.FindAll()
	return
}

func (interactor *TodoInteractor) Add(u domain.Todo) (user domain.Todo, err error) {
	user, err = interactor.TodoRepository.Store(u)
	return
}

func (interactor *TodoInteractor) Update(u domain.Todo) (user domain.Todo, err error) {
	user, err = interactor.TodoRepository.Update(u)
	return
}

func (interactor *TodoInteractor) DeleteById(u domain.Todo) (err error) {
	err = interactor.TodoRepository.DeleteById(u)
	return
}
