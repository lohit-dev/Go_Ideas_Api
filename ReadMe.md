# 🚀 Go Ideas API

A RESTful API for managing project ideas, built with Go.

---

## ✨ Features

- CRUD operations for project ideas
- Tag ideas with technologies and categories
- Track idea status: `requested`, `reviewing`, `planned`, `in-progress`, `published`, `rejected`
- Voting system for ideas
- PostgreSQL for persistent storage
- Interactive API documentation with Swagger

---

## 📚 API Documentation

Once the server is running, you can access Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

---

## ⚙️ Prerequisites

- Go `v1.21+`
- PostgreSQL `v14+`
- Docker (optional, for containerized setup)

---

## 🔐 Environment Variables

| Variable      | Description              | Default     |
| ------------- | ------------------------ | ----------- |
| `PORT`        | Server port              | `8080`      |
| `DB_HOST`     | PostgreSQL host          | `localhost` |
| `DB_PORT`     | PostgreSQL port          | `5433`      |
| `DB_USER`     | PostgreSQL username      | `postgres`  |
| `DB_PASSWORD` | PostgreSQL password      | `postgres`  |
| `DB_NAME`     | PostgreSQL database name | `ideadb`    |
| `DB_SSLMODE`  | PostgreSQL SSL mode      | `disable`   |

---

## 💻 Running Locally

1. Clone the repository:

   ```bash
   git clone https://github.com/lohit-dev/go_ideas_api.git
   cd go_ideas_api
   ```

2. Download dependencies:

   ```bash
   go mod download
   ```

3. Set up a PostgreSQL instance (or use Docker Compose)

4. Run the API server:

   ```bash
   go run cmd/server/main.go
   ```

5. The API will be available at:

   ```
   http://localhost:8080
   ```

---

## 🐳 Running with Docker

### Build the Docker image

```bash
docker build -t go-ideas-api .
```

### Run the container

```bash
docker run -p 8080:8080 go-ideas-api
```

### Or use Docker Compose (API + PostgreSQL)

```bash
docker-compose up
```

Make sure your `docker-compose.yml` and `Dockerfile` are correctly set up before running.

---

## 📦 API Endpoints

| Method | Endpoint        | Description             |
| ------ | --------------- | ----------------------- |
| GET    | `/v1/idea`      | Get all ideas           |
| GET    | `/v1/idea/{id}` | Get a specific idea     |
| POST   | `/v1/idea`      | Create a new idea       |
| POST   | `/v1/idea/{id}` | Update an existing idea |
| DELETE | `/v1/idea/{id}` | Delete an idea          |

---

## 🧱 Project Structure

```
go_ideas_api/
├── cmd/
│   └── server/
│       └── main.go          # Entry point
├── docs/                   # Swagger docs
├── internal/
│   ├── config/             # Configuration handling
│   ├── handler/            # HTTP handlers
│   ├── middleware/         # Middleware (e.g., logging, auth)
│   ├── model/              # Data models
│   ├── router/             # Route definitions
│   ├── service/            # Business logic
│   └── storage/            # DB logic
└── pkg/                    # Shared utilities
```

---

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch:

   ```bash
   git checkout -b feature/my-awesome-feature
   ```

3. Commit your changes:

   ```bash
   git commit -m 'Add my awesome feature'
   ```

4. Push to GitHub:

   ```bash
   git push origin feature/my-awesome-feature
   ```

5. Open a Pull Request 🚀

---

## 📄 License

Licensed under the [MIT License](LICENSE).

---
