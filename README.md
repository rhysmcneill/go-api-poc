
# go-api-poc — Books API (PoC)

Lightweight Proof-of-Concept REST API in Go for learning HTTP handlers, routing, and simple in-memory data management. This repo implements a minimal books service (in-memory) with basic CRUD endpoints so you can experiment with handlers, JSON, and testing.

## What this PoC demonstrates

- Project layout using `cmd/` and `internal/`
- HTTP handlers with `chi` routes
- JSON request/response handling
- Simple in-memory store and service layer
- Example curl commands to exercise the API

## Quick overview

- Server listens on port `3000` (see `cmd/server/main.go`).
- Primary resource: `books` with fields `id`, `Title`, and `Author`.
- JSON field names in this project are case-sensitive and currently use `id`, `Title`, and `Author` to match the model tags.

## Endpoints

- `GET /health` — basic health check
- `GET /books` — list all books
- `GET /books/{id}` — get a single book by numeric id
- `POST /books` — create a new book (JSON body)
- `PUT /books/{id}` — replace/update a book (JSON body)
- `DELETE /books/{id}` — delete a book by id

## Book model

Fields in the `models.Books` struct:

- `id` (int) — JSON key `id`
- `Title` (string) — JSON key `Title`
- `Author` (string) — JSON key `Author`

Example JSON body for creating a book (note capital `Title`/`Author` keys):

```json
{
	"id": 4,
	"Title": "Night of the Museum",
	"Author": "Rhys McNeill"
}
```

## Run locally

Build and run the server from the repository root:

```bash
cd /home/rmcneill/Desktop/go-api-poc
go run ./cmd/server
```

The server listens on `:3000` by default.

## Useful curl examples

- List all books:

```bash
curl -i http://localhost:3000/books
```

- Get one book (id = 2):

```bash
curl -i http://localhost:3000/books/2
```

- Create a new book (returns 201 on success):

```bash
curl -i -X POST http://localhost:3000/books \
	-H "Content-Type: application/json" \
	-d '{"id":5,"Title":"Example Book","Author":"Jane Doe"}'
```

- Update (replace) a book (id = 5):

```bash
curl -i -X PUT http://localhost:3000/books/5 \
	-H "Content-Type: application/json" \
	-d '{"id":5,"Title":"Updated","Author":"Jane"}'
```

- Delete a book (id = 3):

```bash
curl -i -X DELETE http://localhost:3000/books/3
```

## Tests

Run the test suite (if present) with:

```bash
go test ./...
```

---
Updated README to reflect this repo's current Books API PoC.
