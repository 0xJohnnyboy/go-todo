package db

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var dbName = "todo.db"

func GetDB() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    err = db.AutoMigrate(&Task{})
    if err != nil {
        return nil, err
    }

    return db, nil
}
