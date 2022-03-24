package query

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/CodingProjects/Go/GoDo/inputs"
	"github.com/CodingProjects/Go/GoDo/models"
	"github.com/CodingProjects/Go/GoDo/resolvers"
)
type Query struct{}
func (q *Query) TodoFile(ctx context.Context, args struct{ Name string }) *resolvers.TodoFileResolver {
	value := resolvers.TodoFileResolver{}
	// dir := fmt.Sprint("/", args.Name)
	// file, err := os.Open(dir) // For read access.
	// if err != nil {
	// 	log.Printf("error: %s, while trying to open file: %s", err, args.Name)
	// 	return nil
	// }

	files, err := os.ReadDir(".")
	if err != nil {
		log.Println(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())

		if strings.Contains(file.Name(), "TODO") {
			thepath, err := filepath.Abs(filepath.Dir(file.Name()))
			if err != nil {
				log.Println(err)
			}
			log.Printf("path: %s", thepath)

			files2, err := os.ReadDir("TODO")
			if err != nil {
				log.Println(err)
			}

			for _, file2 := range files2 {
				if strings.Contains(file2.Name(), args.Name) {
					path, err := filepath.Abs(filepath.Dir(file2.Name()))
					if err != nil {
						log.Println(err)
					}
					fullpath := thepath + "/" + file.Name() + "/" + args.Name
					log.Printf("fullpath: %s", fullpath)

					text, err := os.ReadFile(fullpath)
					if err != nil {
						log.Println(err)
					}
					data := string(text)
					value = resolvers.TodoFileResolver{
						T: &models.TodoFile{
							Name: args.Name,
							Path: &path,
							Data: &data,
						},
					}
				}
			}
		}
	}

	return &value
}

func (m *Query) CreateTodoFile(args struct{ TodoFile *inputs.TodoFileInput }) *resolvers.TodoFileResolver {

	var thepath string

	fileInfo, err := os.Stat("./TODO")
	if err != nil {
		fmt.Println("Path does not exist!", err)
	}

	mode := fileInfo.Mode()
	if mode.IsDir() {
		fmt.Println("TODO", "is a directory!")
		dst, err := os.Create(filepath.Join("/Users/solocoder32/CodingProjects/Go/GoDo/TODO", filepath.Base(args.TodoFile.Name))) // dir is directory where you want to save file.
		if err != nil {
			log.Printf("error: %s, while trying to open file", err)
		}
		if _, err := dst.Write([]byte(*args.TodoFile.Data)); err != nil {
			dst.Close() // ignore error; Write error takes precedence
			log.Printf("error: %s, while trying to write to file", err)
		}
		if err := dst.Close(); err != nil {
			log.Printf("error: %s, while trying to close file", err)
		}

		defer dst.Close()

		thepath, err = filepath.Abs(filepath.Dir(dst.Name()))
		if err != nil {
			panic(err)
		}
	}

	value := resolvers.TodoFileResolver{
		T: &models.TodoFile{
			Name: args.TodoFile.Name,
			Path: &thepath,
			Data: args.TodoFile.Data,
		},
	}

	return &value
}
