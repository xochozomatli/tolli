package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/xochozomatli/tolli/repo"
)

func main() {
	// Connect to db
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=tolli password=password dbname=tolli sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("error connecting to database: ", err)
	}
	fmt.Println("Connected to postgres")

	repo.CreateUsersTable(db)

	userStrArr := []any{"xoch@example.com", "xoch", "password"}
	repo.AddUser(db, userStrArr)
}
