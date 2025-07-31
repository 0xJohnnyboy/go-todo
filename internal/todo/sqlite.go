package todo

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

const dbFile = "todo.db"

func getDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return nil, err
	}

	schema := `
    CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        done BOOLEAN NOT NULL,
        created_at DATETIME NOT NULL,
        updated_at DATETIME NOT NULL
    );
    `

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	return db, nil
}
