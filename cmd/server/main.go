package main

import (
	"context"
	"log"

	"github.com/Bluesyspyder/Backend-Project/internal/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	ID        int    `json:"id"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}

func main() {
	db.ConnectDB()

	app := fiber.New()


	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,PATCH,DELETE",
	}))

	app.Head("/api/todos", func(c *fiber.Ctx) error{
		return c.SendStatus(200)
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		rows, err := db.DB.Query(
			context.Background(),
			"SELECT id,body,completed FROM todos ORDER BY id DESC",
		)

		if err != nil {
			return err
		}

		defer rows.Close()

		todos := []Todo{}

		for rows.Next() {
			var t Todo
			if err := rows.Scan(&t.ID, &t.Body, &t.Completed); err != nil {
				return err
			}
			todos = append(todos, t)
		}

		return c.Status(200).JSON(todos)
	})

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

	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var todo Todo
		err := db.DB.QueryRow(
			context.Background(),
			`
			UPDATE todos
			SET completed = NOT completed
			WHERE id = $1
			RETURNING id, body, completed
			`,
			id,
		).Scan(&todo.ID, &todo.Body, &todo.Completed)

		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
		}

		return c.Status(201).JSON(todo)
	})

	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		cmd, err := db.DB.Exec(
			context.Background(),
			"DELETE FROM todos WHERE id=$1",
			id,
		)
		if err != nil {
			return err
		}

		if cmd.RowsAffected() == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
		}

		return c.Status(200).JSON(fiber.Map{"msg": "Deleted"})
	})

	log.Fatal(app.Listen(":4000"))
}
