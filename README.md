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
