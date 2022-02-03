package query

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/CodingProjects/Go/GoDo/inputs"
	"github.com/CodingProjects/Go/GoDo/models"
	"github.com/CodingProjects/Go/GoDo/resolvers"
)

func (q *Query) TodoFile(ctx context.Context, args struct{ Name string }) *resolvers.TodoFileResolver {
	file, err := os.Open(args.Name) // For read access.
	if err != nil {
		log.Printf("error: %s, while trying to open file: %s", err, file)
		return nil
	}

	thepath, err := filepath.Abs(filepath.Dir(file.Name()))
	if err != nil {
		panic(err)
	}
	log.Printf("path: %s", thepath)

	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(args.Name)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)

	value := resolvers.TodoFileResolver{
		T: &models.TodoFile{
			Name: args.Name,
			Path: &thepath,
			Data: &text,
		},
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

		// If the file doesn't exist, create it, or append to the file
		// file, err := os.OpenFile(args.TodoFile.Name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		// if err != nil {
		// 	log.Printf("error: %s, while trying to open file: %s", err, file)
		// 	return nil
		// }
		// if _, err := file.Write([]byte(*args.TodoFile.Data)); err != nil {
		// 	file.Close() // ignore error; Write error takes precedence
		// 	log.Fatal(err)
		// }
		// if err := file.Close(); err != nil {
		// 	log.Fatal(err)
		// }

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
