package todo
import (
	"fmt"
)



type Todo struct {
	ID          int
	Title       string
	Description string
	Status      string
}


type TodoList struct {
	Todos []Todo
}

func (t *TodoList) AddTodo(title, description string) {
	fmt.Println("Adding todo...")
	fmt.Println("Title:", title)
	fmt.Println("Description:", description)
	todo := Todo {
		ID:          len(t.Todos) + 1,
		Title:       title,
		Description: description,
		Status:      "pending",
	}

	t.Todos = append(t.Todos, todo)
	fmt.Println("Here is the added todo")
	fmt.Println("Todo added successfully!")
	for _, todo := range t.Todos {
		fmt.Println("ID:", todo.ID)
		fmt.Println("Title:", todo.Title)
		fmt.Println("Description:", todo.Description)
		fmt.Println("Status:", todo.Status)
	}
}

func (t *TodoList) PrintAllTodos() {
	fmt.Println("All Todos:")
	for _, todo := range t.Todos {
		fmt.Println("Yes todo are there")
		fmt.Println("ID:", todo.ID)
		fmt.Println("Title:", todo.Title)
		fmt.Println("Description:", todo.Description)
		fmt.Println("Status:", todo.Status)
		
	}
}

















