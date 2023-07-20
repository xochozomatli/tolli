package repo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func CreateUsersTable(db *sql.DB) {
	// Create users table if none exists
	sqlString := `
	CREATE TABLE IF NOT EXISTS users (
		id serial PRIMARY KEY,
		email VARCHAR (255) UNIQUE NOT NULL,
		username VARCHAR (50) UNIQUE NOT NULL,
		password VARCHAR (50) NOT NULL
	);`
	if _, err := db.Exec(sqlString); err != nil {
		log.Fatal("error creating table: ", err)
	}
	fmt.Println("Created or found table: users'")
}

func AddUser(db *sql.DB, userStrArr []any) {
	// Insert user
	insertUser := `
		INSERT INTO users (email, username, password)
		VALUES ($1, $2, $3);`
	if _, err := db.Exec(insertUser, userStrArr...); err != nil {
		log.Fatal("error inserting user: ", err)
	}
	fmt.Printf("Inserted user '%s'\n", userStrArr[1])
}
