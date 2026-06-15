Caesar cipher API

Simple Go API implementing a Caesar-style shift over a custom alphabet.

Prerequisites
- Go 1.20+ (module-aware)

Run

```bash
go mod tidy
go run ./cmd/app
```

API
- POST `/cipher/:shift` — JSON body: `{"message":"your text"}`

Example

```bash
curl -s -X POST 'http://localhost:8080/cipher/3' \
	-H 'Content-Type: application/json' \
	-d '{"message":"abc123"}'
```

Response example: `{"message":"def456"}`

Notes
- Alphabet: `abcdefghijklmnopqrstuvwxyz1234567890` (lowercase + digits)
- Characters not in the alphabet are returned unchanged.

Docker
- Build and run with Docker:
  ```bash
  docker build -t caesar-cipher .
  docker run -p 8080:8080 --env-file .env caesar-cipher
  ```
- Or start the full app + PostgreSQL stack with Docker Compose:
  ```bash
  docker compose up --build
  ```
- The API is available at `http://localhost:8080`.
- PostgreSQL is configured by default with:
  - `POSTGRES_USERNAME=postgres`
  - `POSTGRES_PASSWORD=sar58yeaf`
  - `POSTGRES_DATABASE=postgres`
  - `POSTGRES_SSLMODE=disable`
  - `POSTGRES_HOST=postgres`
  - `POSTGRES_PORT=5432`
