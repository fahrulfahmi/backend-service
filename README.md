# SnapPOS Backend Service

Backend service for **SnapPOS (Point of Sale Application)**, built as a learning and portfolio project using **Golang** with **Clean Architecture** principles.

---

## 🚀 Technology Stack

* **Go (Golang)** – Backend programming language
* **Gin** – HTTP web framework
* **GORM** – ORM for database access
* **PostgreSQL** – Relational database
* **JWT** – Authentication & authorization
* **Docker** – Containerization (planned)
* **Jenkins** – CI/CD automation (planned)

---

## 📂 Project Structure

```
backendService/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/        # environment & database configuration
│   ├── routes/        # HTTP route definitions
│   ├── auth/          # authentication & JWT middleware
│   ├── user/          # user module (entity, repository, service, handler)
│   ├── product/       # product module
│   └── order/         # order / transaction module
├── pkg/               # shared utilities
├── .env               # environment variables
├── go.mod
├── go.sum
└── README.md
```

The structure follows **Clean Architecture**:

* **Handler** – HTTP layer
* **Service / Usecase** – business logic
* **Repository** – database access
* **Entity / Model** – data representation

---

## ⚙️ Environment Configuration

Create a `.env` file inside `backendService`:

```env
PORT=8080
DB_HOST=localhost
DB_PORT=5433
DB_USER=posuser
DB_PASSWORD=pospass123
DB_NAME=posdb
SECRET_KEY=your_jwt_secret
```

Make sure PostgreSQL is running and the database already exists.

---

## ▶️ Run Application (Local)

1. Go to backend directory

```bash
cd backendService
```

2. Install dependencies

```bash
go mod tidy
```

3. Run the server

```bash
go run cmd/main.go
```

The server will start at:

```
http://localhost:8080
```

---

## 🔐 Authentication & Roles

Authentication uses **JWT** with role-based access control:

* **Admin**

  * Manage users
  * Manage products
  * Access dashboard

* **Owner**

  * View sales reports
  * Monitor stock

* **Cashier**

  * Handle transactions

Roles are validated using JWT middleware.

---

## 📌 API Modules (Progress)

* [x] Authentication (Register, Login, JWT)
* [x] User Management
* [ ] Product Management
* [ ] Order / Transaction
* [ ] Digital Split Bill
* [ ] Dashboard & Reports

---

## 🧪 Database Migration

Database migration is executed automatically on application startup:

```go
db.AutoMigrate(&user.User{})
```

In the future, migrations will be separated into a dedicated migration tool.

---

## 🐳 Docker (Planned)

Planned setup:

* `Dockerfile`
* `docker-compose.yml`
* PostgreSQL container

---

## 🔄 CI/CD (Planned)

* Jenkins (local setup)
* Automatic deployment on GitHub push

---

## 👨‍💻 Author

**Fahrul Fahmi**
Backend Engineer | Learning Fullstack & DevOps

---

## 📝 Notes

This project is intended for:

* Practicing Clean Architecture in Golang
* Implementing JWT & RBAC
* Learning Docker and CI/CD pipelines

Feel free to explore and extend this project 🚀
