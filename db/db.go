package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	// Init connection to DB
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Open pool only allows 10 connections to DB
	// Set 5 connection always ready for connection
	// if more than 10 connects, have to wait for other to connect again
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	// Create Users table for storing accounts
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

	// Create Events table for storing events
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

	// Create Registrations table for storing registration of an event by an account
	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		log.Fatalf("Could not create table: %v", err)
	}
}
