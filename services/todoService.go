package services

import (
	"mongoapp.com/dto"
	"mongoapp.com/model"
	"mongoapp.com/repository"
)

//go:generate mockgen -destination=../mocks/service/todo.go -package=services mongoapp.com/services TodoService
type DefaultTodoService struct {
	Repo repository.TodoRepository
}

type TodoService interface {
	Insert(todo model.Todo) (*dto.TodoDTO, error)
	GetAll() (model.Todos, error)
}

func (t DefaultTodoService) Insert(todo model.Todo) (*dto.TodoDTO, error) {
	var res dto.TodoDTO
	if len(todo.Title) <= 2 {
		res.Status = false
		return &res, nil
	}

	result, err := t.Repo.Insert(todo)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res.Status = true

	return &res, nil
}

func NewTodoService(Repo repository.TodoRepository) DefaultTodoService {
	return DefaultTodoService{Repo: Repo}
}

func (t DefaultTodoService) GetAll() (model.Todos, error) {
	result, err := t.Repo.GetAll()

	if err != nil {
		return nil, err
	}

	return result, nil
}
