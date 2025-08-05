package main

import (
	"todo/cmd"

	d "todo/internal/db"
	t "todo/internal/task"
	u "todo/internal/user"
)

func main() {
	db, err := d.GetDB()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&u.User{}, &t.Task{})
	cmd.Execute()
}
