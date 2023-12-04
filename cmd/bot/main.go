package main

import (
	"database/sql"
	"fmt"
	"log"
	"taskmanager/internal/data/postgres"

	_ "github.com/lib/pq"
)

func main() {
	dsn := `user=postgres password=admin dbname=taskmanager sslmode=disable`
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	repo := postgres.New(db)

	tasks, err := repo.GetAllTasks()
	if err != nil {
		panic(err)
	}
	for _, t := range tasks {
		fmt.Printf("%#v\n", t)
	}

	user, err := repo.GetUserByID(3)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", user)
}
