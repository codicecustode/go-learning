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
	deleteId := deleteCmd.Int("id", -1, "Give the Id for deleting todo ")
	updateTitle := updateCmd.String("title", "", "Kindly give the title for updation")
	updateDescription := updateCmd.String("description", "", "description for updation")
	updateStatus := updateCmd.String("status", "", "Give status for updation")
	flag.Parse()

	todoList := todo.TodoList{}
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
			fmt.Print("Kindlr Provide the Todo Id for deletion")
		}
		todoList.DeleteTodo(*deleteId)

	case "update":
		updateCmd.Parse(flag.Args()[1:])
		if *updateDescription == "" || *updateTitle == "" {
			fmt.Println("Give the title and description fo update")
			return
		}
		todoList.UpdateTodo(*deleteId, *updateStatus, *updateTitle, *updateDescription)

	case "alltodo":
		todoList.PrintAllTodos()
	}
}
