package service

import (
	"database/sql"
	"github.com/google/uuid"
)

type Pessoa struct {
	ID         string   `json:"id"`
	Nome       string   `json:"nome"`
	CpfCnpj    string   `json:"cpf_cnpj"`
	Nascimento string   `json:"nascimento"`
	Seguros    []string `json:"seguros"`
}

type CreatePessoaRequest struct {
	Nome       string   `json:"nome"`
	CpfCnpj    string   `json:"cpf_cnpj"`
	Nascimento string   `json:"nascimento"`
	Seguros    []string `json:"seguros"`
}

type PessoaService struct {
	db *sql.DB
}

func NewPessoaService(db *sql.DB) *PessoaService {
	return &PessoaService{db: db}
}

func (s *PessoaService) GetPessoa(id string) (*Pessoa, error) {

	query := `
SELECT id, nome, cpf_cnpj, nascimento
FROM pessoa
WHERE id = $1
`
	pessoa := Pessoa{}
	err := s.db.QueryRow(query, id).
		Scan(&pessoa.ID, &pessoa.Nome, &pessoa.CpfCnpj, &pessoa.Nascimento)

	if err != nil {
		return &Pessoa{}, err
	}

	return &pessoa, nil
}

func (s *PessoaService) CreatePessoa(pessoa CreatePessoaRequest) (uuid.UUID, error) {

	query := `
INSERT INTO pessoa (id, nome, cpf_cnpj, nascimento)
VALUES ($1, $2, $3, $4) RETURNING id
`
	id := uuid.Nil
	err := s.db.QueryRow(query,
		uuid.New(),
		pessoa.Nome,
		pessoa.CpfCnpj,
		pessoa.Nascimento).
		Scan(&id)

	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
