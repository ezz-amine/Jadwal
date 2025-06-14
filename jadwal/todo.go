package jadwal

import "fmt"

type Todo struct {
	Content string
	IsDone  bool
}

func (t Todo) DoneMarker() string {
	if t.IsDone {
		return "×"
	}

	return "-"
}

func (t *Todo) Done() error {
	if t.IsDone {
		panic(fmt.Sprintf("'%v' is already marked as done", t.Content))
	}
	t.IsDone = true
	return nil
}

func NewTodo(Content string) Todo {
	return Todo{Content: Content}
}
