package inputs

type TodoFileInput struct {
	Name  string
	Path  *string
	Items *[]*TodoItemInput
}

type TodoItemInput struct {
	Completed *bool
	Label     string
}
