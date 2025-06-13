# Starter Go Server

A boilerplate project for building a backend server with **Go** and **Fiber**, integrated with **PostgreSQL** and **Redis**. Includes middleware for **JWT-based authentication**.

---

## 🚀 Features

- ⚙️ Fiber & Gorm setup
- 🛡 JWT authentication middleware
- 🗃 PostgreSQL integration
- 🚀 Redis caching
- 📁 Environment-based configuration

---

## 📦 Prerequisites

Ensure the following are installed on your machine:

- [Go](https://go.dev/)
- [PostgreSQL](https://www.postgresql.org/)
- [Redis](https://redis.io/)

---

## 🛠 Getting Started

```bash
# 1. Clone the repository without history
git clone --depth=1 https://github.com/Maxime-Hrt/Fiber-Starter.git new-project

# 2. Navigate to the project folder
cd new-project

# 3. Remove the remote origin
rm -rf .git

# 4. Install dependencies
go mod tidy

# 5. Start the development server
go run ./cmd/main.go      
```

---

## ⚙️ Configuration
Set up your .env file in the root directory with the following structure:

```txt
# Database
DB_URL=your_postgres_url

# Redis
REDIS_URL=localhost:6379
REDIS_PASSWORD=

# JWT
JWT_SECRET=your_secret_key
```

--- 

## 🧪 Scripts

* `go run ./cmd/main.go` – Start server in development mode
* `go build -o main ./cmd/main.go` – Compile Go to binary
* `./main` – Run compiled app
