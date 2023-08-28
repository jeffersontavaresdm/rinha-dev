package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"rinha-golang/src/controller"
	"rinha-golang/src/database"
	"rinha-golang/src/repository"
)

func main() {
	app := fiber.New()
	db := database.OpenDB()
	pessoaRepository := repository.NewPessoaRepository(db)

	controller.Controller(app, pessoaRepository)

	panic(app.Listen(":3000"))
}
