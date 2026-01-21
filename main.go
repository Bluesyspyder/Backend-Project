package main

import (
	"context"
	// "fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID int `json:"id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

func main() {
	ConnectDB()
	if DB == nil {
	log.Fatal("DB is nil — connection failed")
}
	if err := DB.Ping(context.Background()); err != nil {
		log.Fatal("DB ping failed:", err)
	}
	app:=fiber.New()


	app.Get("/",func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"msg":"API is working"})

	})

	//Create a TODO

	app.Post("/api/todos", func(c *fiber.Ctx)error {
		todo:= new(Todo)

		if err:= c.BodyParser(todo); err != nil{
			return err
		}

		err := DB.QueryRow(
			context.Background(),
			"INSERT INTO todos (body) VALUES ($1) RETURNING id, completed",
			todo.Body,
		).Scan(&todo.ID,&todo.Completed)

		if err!= nil{
			return err
		}

		return c.Status(201).JSON(todo)

	})

	//UPDATE a TODO

app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
	id := c.Params("id")

	var req struct {
		Body *string `json:"body"`
	}
	c.BodyParser(&req)

	var todo Todo
	err := DB.QueryRow(
		context.Background(),
		`
		UPDATE todos
		SET
			body = COALESCE($1, body),
			completed = NOT completed
		WHERE id = $2
		RETURNING id, body, completed
		`,
		req.Body, id, 
	).Scan(&todo.ID, &todo.Body, &todo.Completed)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	}

	return c.JSON(todo)
})

//DELETE a todo

app.Delete("/api/todos/:id", func(c *fiber.Ctx) error{

	id := c.Params("id")

	cmdTag, err := DB.Exec(
		context.Background(),
		"DELETE FROM todos WHERE id=$1",
		id,
	)

	if err != nil{
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error":"Todo not found",
		})
	}

	return c.JSON(fiber.Map{
		"Success":"Todo Deleted",
	})

})


log.Fatal(app.Listen(":4000"))

}