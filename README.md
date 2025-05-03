[![Go](https://img.shields.io/badge/Go-1.24-00ADD8.svg)](https://golang.org)  [![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13-blue.svg)](https://www.postgresql.org)  [![Docker](https://img.shields.io/badge/Docker-20.10-blue.svg)](https://docker.com)  [![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

# Taskoryx  
_A Task Management Backend API_

Welcome to **Taskoryx**, the robust backend service powering your task management needs! 🚀

---

## 🔍 Description
This is the **backend** of the Taskoryx application, built in Go 1.24 with a focus on scalability and clean architecture. A frontend will be developed soon—stay tuned! 🎉

Key highlights:
- **RESTful API** for users and tasks
- **Scalable** design with Fiber, GORM, and Viper
- **Structured logging** via Zap
- **Configurable** through YAML
- **Dockerized** for easy deployment
- **PostgreSQL** as the data store

---

## 🛠 Tech & Packages
We leverage the following technologies and packages:
- **Language**: Go 1.24
- **Web Framework**: Fiber (github.com/gofiber/fiber/v3)
- **ORM**: GORM (gorm.io/gorm)
- **Config**: Viper (github.com/spf13/viper)
- **Logging**: Zap (go.uber.org/zap)
- **Database**: PostgreSQL
- **Containerization**: Docker & Docker Compose
- **Utilities**: UUID (github.com/google/uuid), fsnotify, schema binding, and more

---

## 📁 Project Structure
```bash
shahryar-pirooz-taskoryx/
├── README.md
├── config.sample.yml
├── go.mod
├── go.sum
├── LICENSE
├── api/
│   └── http/
│       ├── setup.go
│       └── handlers/
│           ├── response.go
│           └── user-handler.go
├── app/
│   └── app.go
├── cmd/
│   └── main.go
├── config/
│   ├── config.go
│   └── read.go
├── internal/
│   ├── task/
│   │   ├── service.go
│   │   ├── domain/
│   │   │   └── task-domain.go
│   │   └── port/
│   │       ├── service.go
│   │       └── task.go
│   └── user/
│       ├── service.go
│       ├── domain/
│       │   └── user-domain.go
│       └── port/
│           ├── service.go
│           └── user.go
└── pkg/
    ├── adapters/
    │   └── storage/
    │       ├── task-repo.go
    │       ├── user-repo.go
    │       ├── mapper/
    │       │   ├── task.go
    │       │   └── user.go
    │       └── types/
    │           ├── task-repo-type.go
    │           └── user-repo-type.go
    ├── context/
    │   └── app_context.go
    ├── db/
    │   └── gorm.go
    ├── fp/
    │   └── mappers.go
    └── logger/
        └── zap.go
```

---

## 🏁 Getting Started
### 1. Prerequisites
- Go 1.24+
- PostgreSQL 13+
- Docker & Docker Compose (optional)

### 2. Installation
```bash
git clone https://github.com/yourusername/taskoryx.git
cd taskoryx
cp config.sample.yml config.yml
go mod tidy
```

### 3. Configuration
Edit `config.yml` with your settings:
```yaml
http:
  host: localhost
  port: "8080"

database:
  host: localhost
  port: "5432"
  name: tskrx
  user: ess
  password: 123

production: false
```

### 4. Run Locally
```bash
go run cmd/main.go
```
Visit: `http://localhost:8080/api/v1`

---

## 🐳 Docker & Deployment
Use Docker Compose to spin up the database and service:
```yaml
version: '3'
services:
  postgres:
    image: postgres:latest
    environment:
      POSTGRES_USER: ess
      POSTGRES_PASSWORD: 123
      POSTGRES_DB: tskrx
    ports:
      - "5432:5432"

  taskoryx:
    build: .
    command: go run cmd/main.go
    volumes:
      - ./:/app
    ports:
      - "8080:8080"
    depends_on:
      - postgres
```
```bash
docker-compose up -d
```

---

## 📋 API Endpoints
### Users
| Method | Endpoint            | Description    |
| ------ | ------------------- | -------------- |
| GET    | `/api/v1/tasks`     | List all users |
| GET    | `/api/v1/tasks/:id` | Get user by ID |

### Tasks  *(TBD)*
- Endpoints coming soon

---

## 🤝 Contributing
Feel free to pull a request!✨

All contributions are under the MIT License. Please follow Go community best practices.

---

## 📄 License
MIT License © 2025 DevEss. See [LICENSE](LICENSE) for details.

---

Thanks for checking out Taskoryx! Happy coding! 😊

