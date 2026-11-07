# Simple Go HTTP Server

A lightweight HTTP REST API server written in Go using standard library `net/http` and documented with OpenAPI / Swagger UI.

## Features

- **In-Memory Atomic Counter**: Thread-safe counter using `sync/atomic.Uint64`.
- **Standard Library Routing**: Zero heavy framework dependencies, clean standard `net/http`.
- **Interactive Swagger UI**: Interactive API documentation served directly at `/swagger/index.html`.
- **Automated Docs Generation**: Easily generate and update Swagger specs via `go generate`.

## Project Structure

```text
.
├── catalog-info.yaml               # Backstage service catalog metadata
├── README.md
└── src
    ├── cmd
    │   └── api
    │       └── main.go             # Application entrypoint and Swagger general annotations
    ├── docs                        # Auto-generated Swagger specification files
    │   ├── docs.go
    │   ├── swagger.json
    │   └── swagger.yaml
    ├── go.mod
    ├── go.sum
    └── internal
        └── api
            ├── handlers
            │   └── counter.go      # Counter HTTP handler with Swagger endpoint annotations
            └── main.go             # Server initialization and route definitions
```

## Prerequisites

- [Go](https://go.dev/dl/) 1.25 or newer
- [Swag CLI](https://github.com/swaggo/swag) for generating Swagger docs:
  ```bash
  go install github.com/swaggo/swag/cmd/swag@latest
  ```

## Getting Started

1. Navigate to the source directory:
   ```bash
   cd src
   ```

2. Download and verify dependencies:
   ```bash
   go mod tidy
   ```

3. (Optional) Regenerate Swagger documentation:
   ```bash
   go generate ./...
   ```

4. Run the server:
   ```bash
   go run cmd/api/main.go
   ```

By default, the server listens on port `3000`. You can configure a custom port using the `PORT` environment variable:
```bash
PORT=8080 go run cmd/api/main.go
```

## API Endpoints

| Method | Path | Description | Response Status | Content-Type |
|---|---|---|---|---|
| `GET` | `/counter` | Get current counter value | `200 OK` | `text/plain` |
| `POST` | `/counter` | Increment counter by 1 | `204 No Content` | - |
| `GET` | `/swagger/index.html` | Swagger UI documentation | `200 OK` | `text/html` |

### Examples

**Get counter:**
```bash
curl -i http://localhost:3000/counter
```

**Increment counter:**
```bash
curl -i -X POST http://localhost:3000/counter
```

**Swagger UI:**
Open [http://localhost:3000/swagger/index.html](http://localhost:3000/swagger/index.html) in your browser to interact with the API.

## Documentation Generation

Swagger specs are generated from annotations using `swag`. To update docs after modifying handler annotations:

```bash
cd src
go generate ./...
```

