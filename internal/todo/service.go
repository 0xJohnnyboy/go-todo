package todo

import (
    "errors"
    "time"
)

func AddTask(title string, isDone bool) error {
    if title == "" {
        return errors.New("title is required")
    }

    db, err := getDB()
    if err != nil {
        return err
    }
    defer db.Close()
    now := time.Now()
    _, err = db.Exec(`
    INSERT INTO tasks (title, done, created_at, updated_at) 
    VALUES (?, ?, ?, ?)
    `, title, isDone, now, now)

    return err
}

func ListTasks(showDone bool, showUndone bool) ([]Task, error) {
    db, err := getDB()
    if err != nil {
        return nil, err
    }
    defer db.Close()


    var whereClause string
    switch {
        case !showDone && !showUndone:
        case showDone && showUndone:
        case showDone :
            whereClause = "WHERE done = 1"
        case showUndone:
            whereClause = "WHERE done = 0"
    }
    rows, err := db.Query(`
        SELECT id, title, done, created_at, updated_at
        FROM tasks
        ` + whereClause + `
        ORDER BY created_at DESC
    `) 

    if err != nil {
        return nil, err
    }

    var tasks []Task

    for rows.Next() {
        var task Task
        err = rows.Scan(&task.Id, &task.Title, &task.Done, &task.CreatedAt, &task.UpdatedAt)
        if err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }

    return tasks, nil
}

func DoneTask(id string) error {
    db, err := getDB()
    
    if err != nil {
        return err
    }
    defer db.Close()

    result, err := db.Exec(`
        UPDATE tasks
        SET done = ?, updated_at = ?
        WHERE id = ?
    `, true, time.Now(), id)

    if err != nil {
        return err
    } 

    count, _ := result.RowsAffected()
    
    if count == 0 {
        return errors.New("task not found")
    }

    return nil
}

func DeleteTask(id string) error {
    db, err := getDB()
    if err != nil {
        return err
    }
    defer db.Close()

    result, err := db.Exec(`
        DELETE FROM tasks
        WHERE id = ?
    `, id)

    count, _ := result.RowsAffected()
    if count == 0 {
        return errors.New("task not found")
    }
    return nil
}

func ClearTasks() error {
    db, err := getDB()
    if err != nil {
        return err
    } 
    defer db.Close()

    result, err := db.Exec(`
        DELETE FROM tasks
    `)
    count, _ := result.RowsAffected()
    if count == 0 {
        return errors.New("no tasks to clear")
    }
    return nil
}
