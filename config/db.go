package config

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	db *sql.DB
)

func Connect() {
	d, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	db = d
}

func GetDB() *sql.DB {
	return db
}

func InitDB() {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS Users (
        ID INTEGER PRIMARY KEY AUTOINCREMENT,
		UserName VARCHAR(32) NOT NULL,
		Password VARCHAR(32) NOT NULL
    )`)

	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}
}
