package query

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/CodingProjects/Go/GoDo/common"

	"github.com/CodingProjects/Go/GoDo/inputs"
	"github.com/CodingProjects/Go/GoDo/models"
	"github.com/CodingProjects/Go/GoDo/resolvers"
)

type Query struct{}

func (q *Query) TodoFile(ctx context.Context, args struct{ Name string }) *resolvers.TodoFileResolver {
	value := resolvers.TodoFileResolver{}

	files, err := os.ReadDir(common.BaseTodoDirectory())
	if err != nil {
		log.Println(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), args.Name) {
			filepath, err := filepath.Abs(filepath.Join(common.BaseTodoDirectory(), file.Name()))
			if err != nil {
				log.Println(err)
			}

			// text, err := os.ReadFile(filepath)
			// if err != nil {
			// 	log.Println(err)
			// }
			// data := string(text)
			text := "test"

			value = resolvers.TodoFileResolver{
				T: &models.TodoFile{
					Name: args.Name,
					Path: &filepath,
					Items: models.TodoItem{
						Label: &text,
					},
				},
			}
		}
	}

	return &value
}

func (m *Query) CreateTodoFile(args struct{ TodoFile *inputs.TodoFileInput }) *resolvers.TodoFileResolver {
	thepath := filepath.Join(common.BaseTodoDirectory(), args.TodoFile.Name)
	//var todoFile *os.File
	_, err := os.Stat(thepath)
	if err != nil {
		if _, err = os.Create(thepath); err != nil {
			return nil
		}
	}
	if args.TodoFile.Items != nil {
		println(args.TodoFile)
		//item models.TodoItem
		//if err := json.Unmarshal([]byte(*args.TodoFile.Items), items); err != nil {
		//	return nil
		//}
		//if _, err = todoFile.WriteString(); err != nil {
		//	return nil
		//}
	}

	text := "test"

	value := resolvers.TodoFileResolver{
		T: &models.TodoFile{
			Name: args.TodoFile.Name,
			Path: &thepath,
			Items: models.TodoItem{
				Label: &text,
			},
		},
	}

	return &value
}
