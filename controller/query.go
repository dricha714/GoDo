package controller

import (
	"context"
	"encoding/json"
	"github.com/CodingProjects/Go/GoDo/common"
	"github.com/CodingProjects/Go/GoDo/models"
	"github.com/CodingProjects/Go/GoDo/resolvers"
	"io/ioutil"
	"log"
	"path/filepath"
)

func (_ *Controller) TodoFile(ctx context.Context, args struct{ Name string }) *resolvers.TodoFileResolver {
	var value resolvers.TodoFileResolver

	thepath, err := filepath.Abs(filepath.Join(common.BaseTodoDirectory(), args.Name))
	if err != nil {
		log.Println(err)
		return &value
	}

	var todoFileModel models.TodoFile
	var buff []byte
	if buff, err = ioutil.ReadFile(thepath); err != nil {
		log.Println(err)
		return &value
	}
	if err := json.Unmarshal(buff, &todoFileModel); err != nil {
		log.Println(err)
		return &value
	}

	value = resolvers.TodoFileResolver{
		T: &todoFileModel,
	}

	return &value
}
