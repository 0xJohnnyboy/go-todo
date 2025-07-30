package todo

import (
    "encoding/json"
    "errors"
    "os"
)

const dataFile = "tasks.json"

func LoadTasks() ([]Task, error) {
    var tasks []Task

    if _, err:= os.Stat(dataFile); errors.Is(err, os.ErrNotExist) {
        return tasks, nil
    }

    bytes, err := os.ReadFile(dataFile)

    if err != nil {
        return nil, err
    }

    if len(bytes) == 0 {
        return tasks, nil
    }

    err = json.Unmarshal(bytes, &tasks)

    if err != nil {
        return nil, err
    }

    return tasks, nil
}

func SaveTasks(tasks []Task) error {
    bytes, err := json.MarshalIndent(tasks, "", "  ")
    
    if err != nil {
        return err
    }

    return os.WriteFile(dataFile, bytes, 0644)
}
