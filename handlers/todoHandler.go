package handlers

import (
	"net/http"

	"mongoapp.com/model"
	"mongoapp.com/services"
)

type Todo struct {
	Service services.TodoService
}

func NewTodoHandler(Service services.TodoService) *Todo {
	return &Todo{Service: Service}
}

func (h Todo) Insert(rw http.ResponseWriter, r *http.Request) {
	var todo = model.Todo{}
	err := todo.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Error while parsing todo", http.StatusBadRequest)
		return
	}

	result, err := h.Service.Insert(todo)

	if err != nil || !result.Status {
		http.Error(rw, "Error while inserting todo", http.StatusBadRequest)
		return
	}

	result.ToJSON(rw)

}

func (h Todo) GetAll(rw http.ResponseWriter, r *http.Request) {
	result, err := h.Service.GetAll()

	if err != nil {
		http.Error(rw, "Error while getting todos", http.StatusInternalServerError)
		return
	}

	result.ToJSON(rw)
}
