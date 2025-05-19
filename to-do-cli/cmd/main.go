package main

import (
	"flag"
	"fmt"
	"os"
	"todocli/internal/todo"
)

func main() {

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	//updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	//deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError	)

	addTitle := addCmd.String("title", "", "Title of the todo")
	addDescription := addCmd.String("description", "", "Description of the todo")
	flag.Parse()

	switch flag.Args()[0] {
	case "add":
		addCmd.Parse(flag.Args()[1:])
		if *addTitle == "" || *addDescription == "" {
			fmt.Println("Please provide both title and description")
			os.Exit(1)
		}

		todoList := todo.TodoList{}
		todoList.AddTodo(*addTitle, *addDescription)

	case "allTodos":
		todoList := todo.TodoList{}
		todoList.PrintAllTodos()
	}
}
