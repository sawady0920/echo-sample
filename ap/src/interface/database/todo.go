package database

import "echo-sample/domain"

type TodoRepository struct {
	SqlHandler
}

func (repo *TodoRepository) FindById(id int) (todo domain.Todo, err error) {
	if err = repo.Find(&todo, id).Error; err != nil {
		return
	}
	return
}

func (repo *TodoRepository) FindAll() (todos domain.Todos, err error) {
	if err = repo.Find(&todos).Error; err != nil {
		return
	}
	return
}

func (repo *TodoRepository) Store(u domain.Todo) (todo domain.Todo, err error) {
	if err = repo.Create(&u).Error; err != nil {
		return
	}
	todo = u
	return
}

func (repo *TodoRepository) Update(u domain.Todo) (todo domain.Todo, err error) {
	if err = repo.Save(&u).Error; err != nil {
		return
	}
	todo = u
	return
}

func (repo *TodoRepository) DeleteById(todo domain.Todo) (err error) {
	if err = repo.Delete(&todo).Error; err != nil {
		return
	}
	return
}
