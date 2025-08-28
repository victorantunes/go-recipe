# Go Ping API

A simple REST API in Go with PostgreSQL, Docker, API versioning, and hot reload for development.

---

## Features

- Basic `/api/v1/ping` route returning JSON `{ "message": "pong" }`
- PostgreSQL integration with Docker Compose
- API versioning (`/api/v1`)
- Docker multi-stage build for production
- Hot reload in development with [air](https://github.com/cosmtrek/air)

---

## Folder Structure
```
go-recipe/
├── cmd
│ ├── main.go
│ └── pkg
│     └── validator
│         └── validate.go
├── config
│ └── config.go
├── db
│ ├── postgres.go
│ ├── repositories
│ │ └── user_repository.go
│ └── store.go
├── docker-compose.override.yml
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── handlers
│ └── v1
│     ├── ping_handler.go
│     └── user_handler.go
├── models
│ └── user.go
├── README.md
├── routes
│ └── v1
│     ├── ping_routes.go
│     ├── router.go
│     └── user_routes.go
├── services
│ └── user_service.go

```


## Getting Started

### Prerequisites

- Docker & Docker Compose installed
- (Optional) Go installed locally for development without Docker

---

### Environment Variables

Set these variables in your environment or via Docker Compose:

| Variable      | Description          | Example    |
| ------------- | -------------------- | ---------- |
| `DB_HOST`     | PostgreSQL hostname  | `db`       |
| `DB_PORT`     | PostgreSQL port      | `5432`     |
| `DB_USER`     | PostgreSQL user      | `postgres` |
| `DB_PASSWORD` | PostgreSQL password  | `postgres` |
| `DB_NAME`     | Database name        | `pingdb`   |

---

### Run with Docker (production)

```bash
  docker-compose up --build
```

- API available at: `http://localhost:8080/api/v1/ping`
- PostgreSQL accessible on port 5432

---

### Run with Docker + Hot Reload (development)

```bash
  docker-compose -f docker-compose.yml -f docker-compose.override.yml up --build
```

- Code changes auto-rebuild and restart server inside container
- Avoids polluting local directory with build artifacts (`main`, `.fuse*`)

---

### Test API
```bash
  curl http://localhost:8080/api/v1/ping
```

Expected JSON:
```json
{
  "message": "pong"
}
```

---

## Development

If you want to run locally without Docker:
```bash
  go run ./cmd/main.go
```

---

## Notes

- The `.air.toml` configures the `air` tool for hot reload
- The `Dockerfile` is multi-stage: builds in a Go environment, runs in Alpine
- Volume mounting in `docker-compose.override.yml` enables live code updates

---

## License

MIT License © Núcleo de Tecnologia do MTST
