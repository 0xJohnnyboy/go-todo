# go-todo

A minimal, idiomatic CLI and REST API Todo application written in Go, using GORM with SQLite for persistence and [Cobra](https://github.com/spf13/cobra) for CLI commands.

## ✨ Features

- ✅ Create, list, complete, and delete todos via CLI
- 📊 View daily and overall statistics
- 🌐 Run a REST API server
- 💾 SQLite database (via GORM)
- 🧪 Easily testable modular structure

## 🗂 Project Structure
```
├── cmd # CLI commands
│   ├── add.go
│   ├── clear.go
│   ├── delete.go
│   ├── done.go
│   ├── list.go
│   ├── root.go
│   ├── serve.go
│   ├── stats.go
│   └── version.go
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   └── router.go
│   ├── db
│   │   ├── gormdb.go
│   │   └── task-models.go
│   ├── task
│   │   ├── http.go
│   │   └── service.go
│   └── version
│       └── version.go
├── LICENSE
├── main.go # app entry point
├── Makefile # build commands
├── README.md
└── todo.db
``` 

## 🧱 Installation

```bash
git clone https://github.com/0xJohnnyboy/go-todo.git
cd go-todo
make [build|build-macos|build-windows|build-linux]
```

## 🚀 Usage

### CLI

Assuming your exe is named `todo` and is in your `$PATH`, you can run the following commands
```bash
# Show help
todo [command?] [-h|--help]

# Show version
todo version

# Add a task
todo add "Buy milk"

# List all tasks
todo list

# Mark a task as done
todo done 1

# Clear all tasks
todo clear

# Show stats
todo stats
```

### REST API

```bash
# Start the server
todo serve [--port|-p] 9876 &
# note that port is optional, default is 9876

# List all tasks
curl http://localhost:9876/tasks

# Get a task by id
curl http://localhost:9876/tasks/:id

# Create a task
curl -X POST http://localhost:9876/tasks -H 'Content-Type: application/json' -d '{"title": "Buy milk"}'

# Update a task
curl -X PUT http://localhost:9876/tasks/:id -H 'Content-Type: application/json' -d '{"title": "Buy milk", "done": true}'
# note that if the title field is empty, the task will be updated without changing the title

# Delete a task
curl -X DELETE http://localhost:9876/tasks/:id

# Complete a task
curl -X POST http://localhost:9876/tasks/:id/complete

# Get stats
curl -X GET http://localhost:9876/tasks/stats
``` 

## ⚙️ Requirements

- Go 1.24
- Make
- Git

## 📝 License

This is a learning project licensed under the MIT License - see the [LICENSE](LICENSE) file for details
