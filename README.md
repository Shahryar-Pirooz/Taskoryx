# 📋 Tasks – Clean Architecture Task Manager

A full-stack, self-hostable task management system built with Go (Fiber, SQLX, Redis) and React (TailwindCSS), following **Hexagonal Architecture**. Designed for high performance, code maintainability, and production readiness.
---

## 🚀 Features

- ✅ Hexagonal (Ports & Adapters) Architecture
- 🔐 JWT Authentication & Role-based access
- 🗂 Manage Tasks with Status & Due Dates
- ⚡ Fast Redis-based Caching
- 🐘 PostgreSQL for persistent storage
- 🧾 Structured logging using Zap
- 📦 Config management using Viper
- 🤖 Telegram bot integration for alerts
- 🎨 React + TailwindCSS frontend
- 🐳 Docker + Docker Compose support
---

## 🧱 Tech Stack
| Layer | Tech Used |
|------------|----------------------------------------|
| Frontend | React, TailwindCSS |
| Backend | Go, Fiber, SQLX, Redis, Zap, Viper |
| Database | PostgreSQL |
| Auth | JWT |
| Bot | Telegram |
| DevOps | Docker, Docker Compose |
---

## 📁 Project Structure
See the [Directory Tree](#-directory-tree-backend--frontend---hexagonal) above for details.
---

## 🧰 Installation
### 🚧 Prerequisites
- Go ≥ 1.21
- Node.js ≥ 18
- Docker + Docker Compose
- PostgreSQL + Redis (can be local or via Docker)
### 🔧 Backend Setup
```bash
git clone https://github.com/your-username/tasks.git
cd tasks/backend
# Copy and configure environment
cp .env.example .env
# Run using Docker
docker-compose up --build
```
### 💻 Frontend Setup
```bash
cd ../frontend
npm install
npm run dev
```
Access frontend at: `http://localhost:5173`
Backend API runs at: `http://localhost:8080`
---
## 🧪 Running Tests
```bash
cd backend
go test ./...
```
---
## ✍️ Contribution Guide
1. Fork the repository 🍴
2. Create a new branch: `git checkout -b feature-name`
3. Make your changes ✨
4. Commit: `git commit -m "Add some feature"`
5. Push: `git push origin feature-name`
6. Open a Pull Request 📬
---
## 📜 License
This project is licensed under the **MIT License**. See [LICENSE](./LICENSE) for details.
---
## 💬 Contact
Maintained by [Shahryar Pirooz](https://github.com/Shahryar-Pirooz)
Telegram Bot support coming soon 🤖 
