package common

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Utils
func StrPtr(str string) *string {
	return &str
}

func BaseTodoDirectory() string {
	p, err := filepath.Abs(".")
	if err != nil {
		panic(fmt.Sprintf("getting dir at path: %s", p))
	}
	root := filepath.Join(p, "TODO")
	if _, err := os.Stat(root); err != nil {
		if err := CreateTodoDirectory(); err != nil {
			panic(err)
		}
	}
	return root
}

func CreateTodoDirectory() error {
	path, err := filepath.Abs(".")
	if err != nil {
		return err
	}
	log.Printf("creating dir at path: %s", path)
	if _, err := os.Stat("TODO"); os.IsNotExist(err) {
		// path/to/whatever does not exist
		err = os.Mkdir(filepath.Join(path, "TODO"), 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
