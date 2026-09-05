# Contributing

## Running locally

1. Install Go 1.26+ and a PostgreSQL instance.
2. Copy `.env.example` to `.env` and fill in `PORT`, `SECRET` (JWT signing key), and `DB` (Postgres connection string).
3. Install dependencies and run:
   ```bash
   go mod download
   go run main.go
   ```
4. The API listens on `:8080` by default, exposing `POST /signup`, `POST /login`, and `GET /validate`.

## Running tests

```bash
go vet ./...
go test ./... -v
```

CI runs the same checks (`go vet`, `go build`, `go test`, and a Docker build) on every push/PR to `main`.

## Submitting a change

1. Fork the repo and create a branch off `main`.
2. Keep changes focused — one logical change per PR.
3. Make sure `go vet` and `go test ./...` pass locally before opening a PR.
4. Open a PR describing what changed and why.
