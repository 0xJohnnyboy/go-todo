package main

import "fmt"
import "os"
import "todo/internal/todo"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo [add|list|done|delete] [args...]")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo add [title]")
			os.Exit(1)
		}
		title := os.Args[2]

		err := todo.AddTask(title)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "list":
        err := todo.ListTasks()
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Use: todo done [task id]")
			os.Exit(1)
		}
        err := todo.DoneTask(os.Args[2])
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
		fmt.Printf("Done task: %s\n", os.Args[2])
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Use: todo delete [task id]")
			os.Exit(1)
		}
        err := todo.DeleteTask(os.Args[2])
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
		fmt.Printf("Delete task: %s\n", os.Args[2])
    case "clear":
        err := todo.ClearTasks()
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
	default:
		fmt.Println("Unknown command:", cmd)
	}
}
