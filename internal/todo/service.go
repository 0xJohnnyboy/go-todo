package todo

import (
    "errors"
    "time"
    "fmt"
    "strconv"
)

func AddTask(title string) error {
    if title == "" {
        return errors.New("title is required")
    }

    tasks, err := LoadTasks()

    if err != nil {
        return err
    }

    newId := getNextId(tasks)

    task := Task{
        Id: newId,
        Title: title,
        Done: false,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    tasks = append(tasks, task)

    return SaveTasks(tasks)
}

func ListTasks() error {
    tasks, err := LoadTasks()
    if err != nil {
        return err
    }

    if len(tasks) == 0 {
        fmt.Println("No tasks")
        return nil
    }

    for _, task := range tasks {
        fmt.Printf("%d: %s\n", task.Id, task.Title)
        fmt.Printf("  Done: %t\n", task.Done)
        fmt.Printf("  Created at: %s\n", task.CreatedAt.Format(time.RFC3339))
        fmt.Printf("  Updated at: %s\n", task.UpdatedAt.Format(time.RFC3339))
    }

    return nil
}

func DoneTask(id string) error {
    tasks, err := LoadTasks()

    if err != nil {
        return err
    }

    intId, err := strconv.Atoi(id)
    if err != nil {
        return err
    }

    for i, task := range tasks {
        if task.Id == intId {
            tasks[i].Done = true
            tasks[i].UpdatedAt = time.Now()
            return SaveTasks(tasks)
        }
    }
    
    err = errors.New("task not found")
    return err
}

func DeleteTask(id string) error {
    tasks, err := LoadTasks()

    if err != nil {
        return err
    }

    intId, err := strconv.Atoi(id)
    if err != nil {
        return err
    }

    var newTasks []Task

    for _, task := range tasks {
        if task.Id == intId {
            continue
        }
        newTasks = append(newTasks, task)
    }

    return SaveTasks(newTasks)
}

func ClearTasks() error {
    _, err := LoadTasks()

    if err != nil {
        return err
    }

    return SaveTasks([]Task{})
}


func getNextId(tasks []Task) int {
    maxId := 0
    for _, task := range tasks {
        if task.Id > maxId {
            maxId = task.Id
        }
    }
    return maxId + 1
}
