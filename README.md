# Taskoryx 🧠✅

A modern, self-hostable task management application built with Go (Fiber), PostgreSQL, Redis, and a sleek React frontend. Designed with clean architecture and best practices in mind, Taskoryx also supports Telegram integration and JWT-based authentication.

---

## ✨ Features

- User authentication with JWT
- Secure password hashing
- Task CRUD operations
- PostgreSQL via `sqlx`
- Redis caching
- Logging with Uber Zap
- Configuration management via Viper
- Telegram integration for notifications
- Modern React frontend (Vite + Tailwind)
- Clean architecture structure
- Fully dockerized setup for easy deployment

---

## 📁 Project Structure

See the tree below for a breakdown of the structure.

---

## 🚀 Getting Started

### Requirements

- Go >= 1.21
- PostgreSQL >= 14
- Redis
- Node.js >= 18 (for frontend)
- Docker (optional for dev)

---

### 🧰 Backend Setup

```bash
git clone https://github.com/yourusername/taskoryx.git
cd taskoryx

cp .env.example .env
go mod tidy

# Run the app
go run main.go
```

### 🖥️ Frontend Setup

```bash
cd frontend
npm install
npm run dev
```

### 🐋 Docker Setup

```bash
docker-compose up --build
```

---

## 📚 API Endpoints

### Auth

- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`

### Tasks

- `GET /api/v1/tasks`
- `POST /api/v1/tasks`
- `PUT /api/v1/tasks/:id`
- `DELETE /api/v1/tasks/:id`

---

## 🧱 Project Directory Tree

```bash
taskoryx/
├── cmd/
│   └── server/           # App entrypoint (main.go)
├── config/               # Configuration via Viper
├── internal/
│   ├── auth/             # Auth usecases, handlers
│   ├── task/             # Task usecases, handlers
│   ├── user/             # User logic
│   ├── middleware/       # JWT, logging, etc.
│   └── telegram/         # Telegram bot integration
├── pkg/
│   ├── db/               # PostgreSQL + Redis setup (sqlx)
│   ├── logger/           # Zap logger setup
│   └── utils/            # Utilities and helpers
├── migrations/           # SQL migration files
├── frontend/             # React + Tailwind frontend
├── .env.example          # Example env file
├── docker-compose.yml    # Docker configuration
└── README.md
```

---

## 📂 Directory Purpose Table

| Directory         | Purpose |
|------------------|---------|
| `cmd/server`      | Application entrypoint with `main.go` |
| `config/`         | Loads configuration using Viper |
| `internal/auth`   | Handles JWT auth, password logic |
| `internal/task`   | Core business logic for tasks |
| `internal/user`   | User registration, validation |
| `internal/middleware` | Middleware for Fiber (JWT, Logging) |
| `internal/telegram` | Telegram bot handler for notifications |
| `pkg/db`          | DB connection logic for PostgreSQL & Redis |
| `pkg/logger`      | Zap logger instance and configuration |
| `pkg/utils`       | Utility functions (e.g., hashing) |
| `migrations/`     | SQL migration files for DB schema |
| `frontend/`       | React app with Tailwind |
| `.env.example`    | Template for environment variables |

---

## 📬 Telegram Integration

To receive task notifications via Telegram:

1. Create a bot via [@BotFather](https://t.me/botfather)
2. Add your bot token and chat ID to `.env`
3. Enable notifications in your profile

---

## ✅ License

MIT — FOSS and free to self-host.
