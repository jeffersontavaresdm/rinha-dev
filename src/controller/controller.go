package controller

import (
	"github.com/gofiber/fiber/v2"
	"rinha-golang/src/repository"
	"rinha-golang/src/service"
)

func Controller(app *fiber.App, pessoaRepository *repository.PessoaRepository) {

	app.Get("/pessoa", findAll(pessoaRepository))

	app.Get("/pessoa/:id", getById(pessoaRepository))

	app.Post("/pessoa", create(pessoaRepository))
}

func findAll(pessoaRepository *repository.PessoaRepository) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return service.FindAll(pessoaRepository, c)
	}
}

func getById(pessoaRepository *repository.PessoaRepository) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return service.GetById(pessoaRepository, c)
	}
}

func create(pessoaRepository *repository.PessoaRepository) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return service.Create(pessoaRepository, c)
	}
}
