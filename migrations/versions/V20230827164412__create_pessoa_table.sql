CREATE TABLE pessoa
(
    id         VARCHAR,
    cpf_cnpj   VARCHAR,
    nome       VARCHAR,
    nascimento TIMESTAMPTZ,
    seguros    JSONB
);

CREATE UNIQUE INDEX idx_id ON pessoa (id);
CREATE UNIQUE INDEX idx_pessoa_cpf_cnpj ON pessoa (cpf_cnpj);
