package main

import (
	"fmt"

	"github.com/ezz-amine/Jadwal/jadwal"
)

func main() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()

	todos := []jadwal.Todo{}
	todos = append(todos, jadwal.NewTodo("Some task"), jadwal.NewTodo("Some other taks"))
	for idx, todo := range todos {
		c := "-"
		if todo.IsDone {
			c = "x"
		}

		fmt.Printf("[%s] %s (%d)\n", c, todo.Content, idx)
	}
}
