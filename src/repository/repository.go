package repository

import (
	"database/sql"
	"encoding/json"
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

type PessoaRepository struct {
	db *sql.DB
}

func NewPessoaRepository(db *sql.DB) *PessoaRepository {
	return &PessoaRepository{db: db}
}

func (s *PessoaRepository) GetPessoas() (*[]Pessoa, error) {
	query := `
SELECT id, nome, cpf_cnpj, nascimento, seguros
FROM pessoa
`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var pessoas []Pessoa
	for rows.Next() {
		var pessoa Pessoa
		var segurosJSON []byte

		err := rows.Scan(&pessoa.ID, &pessoa.Nome, &pessoa.CpfCnpj, &pessoa.Nascimento, &segurosJSON)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(segurosJSON, &pessoa.Seguros)
		if err != nil {
			return nil, err
		}

		pessoas = append(pessoas, pessoa)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &pessoas, nil
}

func (s *PessoaRepository) GetPessoaById(id string) (*Pessoa, error) {
	query := `
SELECT id, nome, cpf_cnpj, nascimento, seguros
FROM pessoa
WHERE id = $1
`
	pessoa := Pessoa{}
	var segurosJSON []byte
	err := s.db.QueryRow(query, id).Scan(&pessoa.ID, &pessoa.Nome, &pessoa.CpfCnpj, &pessoa.Nascimento, &segurosJSON)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(segurosJSON, &pessoa.Seguros)
	if err != nil {
		return nil, err
	}

	return &pessoa, nil
}

func (s *PessoaRepository) CreatePessoa(pessoa CreatePessoaRequest) (uuid.UUID, error) {
	query := `
INSERT INTO pessoa (id, nome, cpf_cnpj, nascimento, seguros)
VALUES ($1, $2, $3, $4, $5) RETURNING id
`
	segurosJSON, err := json.Marshal(pessoa.Seguros)
	if err != nil {
		return uuid.Nil, err
	}

	id := uuid.Nil
	err = s.db.QueryRow(
		query,
		uuid.New(),
		pessoa.Nome,
		pessoa.CpfCnpj,
		pessoa.Nascimento,
		segurosJSON,
	).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
