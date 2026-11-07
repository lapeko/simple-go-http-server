# Runbook: Counter Service

## Health Check

Check current counter value:
```bash
curl http://localhost:3000/counter
```

Increment counter:
```bash
curl -X POST http://localhost:3000/counter
```

---

## Troubleshooting

### 1. Port Conflict (`bind: address already in use`)
Port `3000` is already in use by another process (e.g., Backstage frontend).

Start on a different port:
```bash
PORT=8080 go run cmd/api/main.go
```

---

### 2. Counter Resets to 0
The counter is stored in-memory using `sync/atomic.Uint64`. Any process or container restart will reset the counter to zero by design.

---

### 3. Swagger UI Returns 404 (`/swagger/index.html`)
Regenerate swagger documentation files:
```bash
cd src
go generate ./...
```

---

## Service Restart

Stop the running process:
```bash
pkill -f "cmd/api/main"
```

Start the service:
```bash
go run cmd/api/main.go
```
