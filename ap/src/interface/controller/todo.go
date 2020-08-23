package controller

import (
	"strconv"

	"echo-sample/domain"
	"echo-sample/interface/database"
	"echo-sample/usecase"

	"github.com/labstack/echo"
)

type TodoController struct {
	Interactor usecase.TodoInteractor
}

func NewTodoController(sqlHandler database.SqlHandler) *TodoController {
	return &TodoController{
		Interactor: usecase.TodoInteractor{
			TodoRepository: &database.TodoRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TodoController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := controller.Interactor.TodoById(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, todo)
	return
}

func (controller *TodoController) Index(c echo.Context) (err error) {
	todos, err := controller.Interactor.Todos()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, todos)
	return
}

func (controller *TodoController) Create(c echo.Context) (err error) {
	t := domain.Todo{} 
	c.Bind(&t)
	todo, err := controller.Interactor.Add(t)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, todo)
	return
}

func (controller *TodoController) Save(c echo.Context) (err error) {
	t := domain.Todo{}
	c.Bind(&t)
	todo, err := controller.Interactor.Update(t)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, todo)
	return
}

func (controller *TodoController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo := domain.Todo{
		ID: id,
	}
	err = controller.Interactor.DeleteById(todo)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, todo)
	return
}
