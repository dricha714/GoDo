package resolvers

import "github.com/CodingProjects/Go/GoDo/models"

type TodoFileResolver struct {
	T *models.TodoFile
}

func (r *TodoFileResolver) Name() string  { return r.T.Name }
func (r *TodoFileResolver) Path() *string { return r.T.Path }
func (r *TodoFileResolver) Items() TodoItemResolver {
	//return r.T.Items
	 b := "false"
	// item := models.TodoItem{
	// 	// Completed: &b,
	// } 
	res := models.TodoItem{
		Label: &b,
	}
	// res = append(res, item)
	r.T.Items = res

	re := TodoItemResolver{
		T: &models.TodoItem{
			 Label: &b,
		},
	}

	return re
}
