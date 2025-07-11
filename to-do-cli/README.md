# go-learning

# 📝 To-Do CLI App (Go)

A simple and efficient command-line to-do list application written in Go. It allows users to add, update, delete, and view their tasks, all saved locally in a JSON file.

---

## 📂 Folder Structure

to-do-cli/
├── cmd/
│   └── main.go              # Entry point of the CLI
├── internal/
│   └── todo/
│       └── todo.go          # Core logic for managing todos
├── go.mod                   # Go module file
├── todo.json                # JSON file to store todos
└── README.md                # Project documentation

---

## ⚙️ Setup Instructions

### 1. Clone the Repository

git clone https://github.com/codicecustode/to-do-cli.git  
cd to-do-cli

### 2. Build the Project

go build -o todo ./cmd

This compiles the project and creates an executable named:

- `todo` on Linux/macOS  
- `todo.exe` on Windows

---

## 🚀 Usage

> **Note:** Use `todo` on Linux/macOS terminals, and `todo.exe` on Windows command prompt or PowerShell.

### ➕ Add a Todo

Linux/macOS:  
todo add -title="Read book" -description="Finish Clean Code"

Windows:  
todo.exe add -title="Read book" -description="Finish Clean Code"

### 📝 Update a Todo

Linux/macOS:  
todo update -id=1 -title="Read updated" -description="Updated desc" -status="done"

Windows:  
todo.exe update -id=1 -title="Read updated" -description="Updated desc" -status="done"

### ❌ Delete a Todo

Linux/macOS:  
todo delete -id=1

Windows:  
todo.exe delete -id=1

### 📋 View All Todos

Linux/macOS:  
todo alltodo

Windows:  
todo.exe alltodo

---

## 🧪 Example Workflow

Linux/macOS:  
todo add -title="Buy groceries" -description="Milk, Bread, Eggs"  
todo alltodo  
todo update -id=1 -title="Buy essentials" -description="Milk, Bread, Eggs, Butter" -status="done"  
todo delete -id=1

Windows:  
todo.exe add -title="Buy groceries" -description="Milk, Bread, Eggs"  
todo.exe alltodo  
todo.exe update -id=1 -title="Buy essentials" -description="Milk, Bread, Eggs, Butter" -status="done"  
todo.exe delete -id=1

---

## ✅ Features

- Add tasks with title and description  
- Update title, description, and status by ID  
- Delete tasks by ID  
- View all tasks  
- All data stored locally in `todo.json`

---

## 📄 License

Licensed under the MIT License.

---

## 🙋‍♂️ Author

Made by **Aman Singh** — feel free to connect!
