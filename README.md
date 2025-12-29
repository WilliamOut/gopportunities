# gopportunities

API para gerenciar vagas de emprego(Go + Gin + GORM + PostgreSQL).

## Resumo
Projeto backend para criar/atualizar/remover e listar vagas de emprego.

## Requisitos
- Go 1.18+
- PostgreSQL (ou conexão via `DB_SOURCE`)
- Docker (Devcontainer)

## Variáveis de ambiente
- `DB_SOURCE` - string de conexão PostgreSQL (DSN). Exemplo:
  `postgresql://user:password@localhost:5432/gopportunities?sslmode=disable`

Se `DB_SOURCE` não estiver definida, o projeto usa por padrão:
`postgresql://user:password@db:5432/gopportunities?sslmode=disable`

## Instalação e execução
1. Baixe as dependências:

```bash
cd /workspace
go env -w GOSUMDB=off
go mod download
```

2. Rodar localmente (com `DB_SOURCE` configurada se necessário):

```bash
export DB_SOURCE="postgresql://user:password@localhost:5432/gopportunities?sslmode=disable"
go run main.go
```

3. Build:

```bash
go build ./...
```

## Docker
Build da imagem Docker:

```bash
docker build -t gopportunities:latest .
```

Rodar (exemplo com container do Postgres):

```bash
# executar Postgres (exemplo rápido)
docker run --name go-postgres -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=gopportunities -p 5432:5432 -d postgres

# executar a aplicação
docker run --rm -e DB_SOURCE="postgresql://user:password@host.docker.internal:5432/gopportunities?sslmode=disable" -p 8080:8080 gopportunities:latest
```

## Endpoints (API v1)
- `GET /api/v1/opening?id={id}` - obter vaga por ID
- `POST /api/v1/opening` - criar vaga (JSON)
- `PUT /api/v1/opening?id={id}` - atualizar vaga (JSON)
- `DELETE /api/v1/opening?id={id}` - deletar vaga
- `GET /api/v1/openings` - listar todas as vagas

## Swagger
A documentação Swagger está disponível em `/swagger/*any` quando o servidor estiver rodando.

## Estrutura do projeto
- `main.go` — ponto de entrada
- `handler/` — handlers HTTP
- `router/` — definição de rotas
- `schemas/` — modelos GORM
- `config/` — inicialização do DB e logger



## Licença
MIT
