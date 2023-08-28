package service

import (
	"github.com/gofiber/fiber/v2"
	"rinha-golang/src/repository"
)

func FindAll(pessoaRepository *repository.PessoaRepository, c *fiber.Ctx) error {
	pessoas, err := pessoaRepository.GetPessoas()
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(pessoas)
}

func GetById(pessoaRepository *repository.PessoaRepository, c *fiber.Ctx) error {
	id := c.Params("id")
	pessoa, err := pessoaRepository.GetPessoaById(id)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(pessoa)
}

func Create(pessoaRepository *repository.PessoaRepository, c *fiber.Ctx) error {
	pessoa := repository.CreatePessoaRequest{}
	err := c.BodyParser(&pessoa)
	if err != nil {
		return c.SendStatus(400)
	}

	id, err := pessoaRepository.CreatePessoa(pessoa)

	if err != nil {
		return c.SendStatus(500)
	}

	return c.Status(201).SendString(id.String())
}
