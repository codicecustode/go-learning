package main

import (
	"flag"
	"fmt"
	"os"
	"todocli/internal/todo"
)

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	addTitle := addCmd.String("title", "", "Title of the todo")
	addDescription := addCmd.String("description", "", "Description of the todo")

	deleteId := deleteCmd.Int("id", -1, "ID of the todo to delete")

	updateTitle := updateCmd.String("title", "", "Title for update")
	updateDescription := updateCmd.String("description", "", "Description for update")
	updateStatus := updateCmd.String("status", "", "Status for update")

	flag.Parse()

	todoList := todo.TodoList{}

	if len(flag.Args()) < 1 {
		fmt.Println("Expected 'add', 'delete', 'update', or 'alltodo' subcommands")
		os.Exit(1)
	}

	switch flag.Args()[0] {
	case "add":
		addCmd.Parse(flag.Args()[1:])
		if *addTitle == "" || *addDescription == "" {
			fmt.Println("Please provide both title and description")
			os.Exit(1)
		}
		todoList.AddTodo(*addTitle, *addDescription)

	case "delete":
		deleteCmd.Parse(flag.Args()[1:])
		if *deleteId == -1 {
			fmt.Println("Please provide the Todo ID for deletion")
			os.Exit(1)
		}
		todoList.DeleteTodo(*deleteId)

	case "update":
		updateCmd.Parse(flag.Args()[1:])
		if *updateTitle == "" || *updateDescription == "" {
			fmt.Println("Please provide both title and description for update")
			os.Exit(1)
		}
		todoList.UpdateTodo(*deleteId, *updateStatus, *updateTitle, *updateDescription)

	case "alltodo":
		todoList.PrintAllTodos()

	default:
		fmt.Println("Unknown command. Available commands: add, delete, update, alltodo")
	}
}
