package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	dsn := "postgres://postgres:vardaanbajpai492@14@localhost:5432/todo_app?sslmode=disable"

	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("DB CONNECTION FAILED:", err)
	}

	DB = db
	log.Println("Postgres connected")
}


