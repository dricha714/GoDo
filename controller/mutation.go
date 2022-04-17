package controller

import (
	"encoding/json"
	"github.com/CodingProjects/Go/GoDo/common"
	"github.com/CodingProjects/Go/GoDo/inputs"
	"github.com/CodingProjects/Go/GoDo/models"
	"github.com/CodingProjects/Go/GoDo/resolvers"
	"os"
	"path/filepath"
)

func (_ *Controller) CreateTodoFile(args struct{ TodoFile *inputs.TodoFileInput }) *resolvers.TodoFileResolver {
	thepath := filepath.Join(common.BaseTodoDirectory(), args.TodoFile.Name)
	_, err := os.Stat(thepath)
	if err != nil {
		if _, err = os.Create(thepath); err != nil {
			return nil
		}
	}

	var todoItems []*models.TodoItem
	if args.TodoFile.Items != nil {
		for _, todoItemInput := range *args.TodoFile.Items {
			var completed bool
			if todoItemInput.Completed != nil {
				completed = *todoItemInput.Completed
			}
			todoItem := models.TodoItem{
				completed,
				todoItemInput.Label,
			}
			todoItems = append(todoItems, &todoItem)
		}
	}

	todoFileModel := models.TodoFile{
		Name:  args.TodoFile.Name,
		Path:  &thepath,
		Items: &todoItems,
	}

	buff, err := json.Marshal(todoFileModel)

	if err := os.WriteFile(thepath, buff, 0666); err != nil {
		return &resolvers.TodoFileResolver{}
	}

	value := resolvers.TodoFileResolver{
		T: &todoFileModel,
	}

	return &value
}
