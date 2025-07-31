package todo

import (
    "errors"
)

func AddTask(title string, isDone bool) error {
    if title == "" {
        return errors.New("title is required")
    }

    db, err := getDB()
    if err != nil {
        return err
    }

    task := Task{
        Title: title,
        Done: isDone,
    }

    return db.Create(&task).Error
}

func ListTasks(showDone bool, showUndone bool) ([]Task, error) {
    db, err := getDB()
    if err != nil {
        return nil, err
    }

    var tasks []Task
    q := db.Order("created_at DESC")
    switch {
    case !showDone && !showUndone, showDone && showUndone:
    case showDone :
        q = q.Where("done = ?", true)
    case showUndone:
        q = q.Where("done = ?", false)
    }

    err = q.Find(&tasks).Error
    return tasks, err
}

func DoneTask(id string) error {
    db, err := getDB()
    
    if err != nil {
        return err
    }

    return db.Model(&Task{}).Where("id = ?", id).Update("done", true).Error
}

func DeleteTask(id string) error {
    db, err := getDB()
    if err != nil {
        return err
    }

    return db.Delete(&Task{}, id).Error
}

func ClearTasks() error {
    db, err := getDB()
    if err != nil {
        return err
    } 

    return db.Where("1 = 1").Delete(&Task{}).Error
}
