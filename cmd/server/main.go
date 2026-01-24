package main

import (
	"context"
	"log"

	"github.com/Bluesyspyder/Backend-Project/internal/db"
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}

func main() {
	db.ConnectDB()

	app := fiber.New()

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := new(Todo)

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		err := db.DB.QueryRow(
			context.Background(),
			"INSERT INTO todos (body) VALUES ($1) RETURNING id, body, completed",
			todo.Body,
		).Scan(&todo.ID, &todo.Body, &todo.Completed)

		if err != nil {
			return err
		}

		return c.Status(201).JSON(todo)
	})

	log.Fatal(app.Listen(":4000"))
}
