# 🧩 Go Tasks API

A beginner-friendly **REST API** built in **Go**, designed to practice real-world backend fundamentals — HTTP routing, JSON handling, validation, persistence, and testing — using idiomatic Go and the standard library first.

---

## 🎯 Goal

Build a small, maintainable REST API for managing tasks.  
Learn how to:

- Structure a Go project cleanly (`cmd/`, `internal/`)
- Implement CRUD routes using `net/http`
- Handle JSON requests and responses
- Add validation and consistent error handling
- Swap between an in-memory and SQLite data store
- Write tests for handlers and data layers

---

## ⚙️ Tech Stack

| Component | Choice | Notes |
|------------|---------|-------|
| Language | Go 1.21+ | No frameworks — focus on standard library |
| Routing | `net/http` | Learn fundamentals before using `chi` |
| Storage | In-memory → SQLite | Same interface, easy swap |
| Config | Environment vars | Keep it simple |
| Testing | `testing` + `httptest` | Standard Go tooling |

---

## 📁 Project Structure

``` bash
myapi/
cmd/api/ # main entrypoint (main.go)
internal/http/ # handlers, routing, middleware
internal/store/ # in-memory + SQLite stores
internal/core/ # validation, error helpers
testdata/ # example JSON or fixtures
```


---

## 🧱 Development Phases

| Phase | Goal |
|-------|------|
| **0** | Define API contract (README + curl tests) ✅ |
| **1** | Implement server, routes, and in-memory store |
| **2** | Add SQLite persistence and swapable store |
| **3** | Add validation, error helpers, and logging |
| **4** | Write tests for handlers and stores |
| **5** *(optional)* | Add pagination, filtering, and CORS |

---

## 🧪 Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/healthz` | Health check |
| `GET` | `/tasks` | List tasks |
| `POST` | `/tasks` | Create task |
| `GET` | `/tasks/{id}` | Get one task |
| `PATCH` | `/tasks/{id}` | Update task |
| `DELETE` | `/tasks/{id}` | Delete task |

---

## 🧠 Key Learning Outcomes

- Structuring Go projects for clarity and growth  
- Using `context` and `sync` safely  
- Writing clean, testable handlers  
- Understanding REST design and validation patterns  

---

## 🚀 Getting Started

```bash
go run ./cmd/api
curl http://localhost:8080/healthz
