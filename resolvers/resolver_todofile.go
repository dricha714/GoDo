package resolvers

import "github.com/CodingProjects/Go/GoDo/models"

type TodoFileResolver struct {
	T *models.TodoFile
}

func (r *TodoFileResolver) Name() string  { return r.T.Name }
func (r *TodoFileResolver) Path() *string { return r.T.Path }
func (r *TodoFileResolver) Items() *[]*models.TodoItem {
	return r.T.Items
}
