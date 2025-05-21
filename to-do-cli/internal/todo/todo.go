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
	todo := Todo {
		ID:          len(t.Todos) + 1,
		Title:       title,
		Description: description,
		Status:      "pending",
	}
	t.Todos = append(t.Todos, todo)
	fmt.Println("Here is the added todo")
	fmt.Println("Todo added successfully!")
}

func (t *TodoList) PrintAllTodos() {
	fmt.Println("Here is All Todo............")
	for _, todo := range t.Todos {
		fmt.Println("ID:", todo.ID)
		fmt.Println("Title:", todo.Title)
		fmt.Println("Description:", todo.Description)
		fmt.Println("Status:", todo.Status)
		
	}
}

func (t *TodoList) UpdateTodo(id int, status, title, description string){

	isIdFound := false

	for i := range t.Todos {
		if t.Todos[i].ID == id {
			t.Todos[i].Status = status
			t.Todos[i].Title = title
			t.Todos[i].Description = description
			isIdFound = true
			break
		}

	}

	if isIdFound {
		fmt.Println("Todo are updated successfully.")
	} else {
		fmt.Println("Unable to find the specified id todo.")
	}

}

func (t *TodoList ) DeleteTodo(id int){
	deletedId := -1
	for i, todo := range t.Todos {
		if todo.ID == id {
			deletedId = i
			break
		}
	}
	
	if deletedId == -1 {
		fmt.Println("Deleted ID not Found")
		return
	}
	t.Todos = append(t.Todos[:deletedId], t.Todos[deletedId+1:]...)
	fmt.Println("Todo deleted succesfull")
}















