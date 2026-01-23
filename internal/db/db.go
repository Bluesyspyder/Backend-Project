package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)


var DB *pgxpool.Pool


func ConnectDB(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == ""{
		log.Fatal("Database URL not found")
	}

	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(context.Background())
	err != nil {
		log.Fatal(err)
	}

	DB = db
	log.Println("Postgres connected succesfully")
}
