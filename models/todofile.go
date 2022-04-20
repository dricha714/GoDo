package models

type TodoFile struct {
	Name  string
	Path  *string
	Items TodoItem
}
