package main

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"os"
	"rinha-golang/service"
	"time"
)

func openDB() *sql.DB {

	// Format: postgres://user:password@host:port/dbname?connect_timeout=5&sslmode=disable
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		panic("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(30)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	return db

}

func main() {

	db := openDB()

	pessoaSvc := service.NewPessoaService(db)

	app := fiber.New()

	app.Post("/pessoa", func(c *fiber.Ctx) error {

		pessoa := service.CreatePessoaRequest{}
		err := c.BodyParser(&pessoa)
		if err != nil {
			return c.SendStatus(400)
		}

		id, err := pessoaSvc.CreatePessoa(pessoa)

		if err != nil {
			return c.SendStatus(500)
		}

		return c.Status(201).SendString(id.String())
	})

	app.Get("/pessoa/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		pessoa, err := pessoaSvc.GetPessoa(id)
		if err != nil {
			return c.SendStatus(500)
		}

		return c.JSON(pessoa)
	})

	app.Get("/pessoa", func(c *fiber.Ctx) error {
		termo := c.Params("t")
		return c.SendString(termo)
	})

	app.Listen(":3000")
}
