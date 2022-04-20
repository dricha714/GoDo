package resolvers

import "github.com/CodingProjects/Go/GoDo/models"

type TodoItemResolver struct {
	T *models.TodoItem
}

// func (r *TodoItemResolver) Completed() *bool { return r.T.Completed }
func (r *TodoItemResolver) Label() *string   { return r.T.Label }
