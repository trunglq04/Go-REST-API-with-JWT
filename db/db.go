package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Open pool only allows 10 connections to DB
	// if more than 10 connects, they have to wait for other to connect again
	DB.SetMaxOpenConns(10)
	// Set 5 connection always ready for connection
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		log.Fatalf("Could not create table: %v", err)
	}

	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events ( 
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL, 
        description TEXT NOT NULL,
        location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
    )
	`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("Could not create table: %v", err)
	}
}
