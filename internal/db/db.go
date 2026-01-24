package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func ConnectDB() {
	// load .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not found")
	}

	var err error
	DB, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	log.Println("Postgres connected successfully")
}
