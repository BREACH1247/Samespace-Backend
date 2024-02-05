package api

import (
	"github.com/BREACH1247/todo-api/internal/models"
	"github.com/BREACH1247/todo-api/internal/repository"
	"github.com/gocql/gocql"
)

type TodoAPI struct {
	repo repository.TodoRepository
}

func NewTodoAPI(repo repository.TodoRepository) *TodoAPI {
	return &TodoAPI{
		repo: repo,
	}
}

func (api *TodoAPI) AddTodo(todo models.Todo) error {
	return api.repo.Create(todo)
}
func (api *TodoAPI) GetTodo(id gocql.UUID) (models.Todo, error) {
	return api.repo.GetByID(id)
}

func (api *TodoAPI) UpdateTodo(todo models.Todo) error {
	return api.repo.Update(todo)
}

func (api *TodoAPI) DeleteTodo(id gocql.UUID) error {
	return api.repo.Delete(id)
}

func (api *TodoAPI) GetTodoList(page int, pageSize int, offset int, status string) ([]models.Todo, error) {
	return api.repo.List(page, pageSize, offset , status)
}


