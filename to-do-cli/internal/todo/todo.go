package todo


import "flag"

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
	todo := Todo {
		ID:          len(t.Todos) + 1,
		Title:       title,
		Description: description,
		Status:      "pending",
	}

	t.Todos = append(t.Todos, todo)
}

func (t *TodoList) UpdateTodo(id int, title, description string) {
	for i, todo := range t.Todos {
		if todo.ID == id {
			t.Todos[i].Title = title
			t.Todos[i].Description = description
			return
		}
	}
}

func (t *TodoList) DeleteTodo(id int) {
	for i, todo := range t.Todos {
		if(todo.ID == id) {
			t.Todos = append(t.Todos[:i], t.Todos[i+1:]...)
		}
	}
}

func (t *TodoList) ListTodos() {
	
}
func (t *TodoList) GetTodo(id int) *Todo {

}

func (t *TodoList) GetTodos() []Todo {

}


func (t *TodoList) getTodosByDate() []Todo {
	
}








