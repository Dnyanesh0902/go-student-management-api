# 🚀 Student Management API - Golang

A professional backend REST API project built using **Go**, **Gin**, **GORM**, **SQL Server**, and **JWT Authentication** following **Clean Architecture** and **Repository Pattern**.

This project demonstrates enterprise-level backend development concepts such as:

- JWT Authentication
- Role-Based Authorization
- Middleware
- Repository Pattern
- Service Layer
- DTO Validation
- SQL Server Integration
- Clean Architecture
- Logging Middleware

---

# 🛠 Tech Stack

## Backend

- Golang
- Gin Framework
- GORM ORM
- SQL Server
- JWT Authentication
- bcrypt Password Hashing

---

# 📂 Project Structure

```text
student-management-api/
│
├── cmd/
│   └── main.go
│
├── controllers/
│   ├── auth_controller.go
│   └── student_controller.go
│
├── database/
│   └── db.go
│
├── dto/
│   └── student_dto.go
│
├── middleware/
│   ├── jwt_middleware.go
│   ├── role_middleware.go
│   └── logger_middleware.go
│
├── models/
│   ├── student.go
│   └── user.go
│
├── repositories/
│   └── student_repository.go
│
├── routes/
│   └── routes.go
│
├── services/
│   └── student_service.go
│
├── utils/
│   ├── jwt.go
│   ├── password.go
│   └── validation.go
│
├── .env
├── go.mod
├── go.sum
└── README.md
```

---

# ⚙️ Features

## ✅ Authentication

- User Registration
- User Login
- JWT Token Generation
- Password Hashing using bcrypt

---

## ✅ Authorization

- Role-Based Access Control
- Admin Protected APIs
- JWT Middleware Protection

---

## ✅ Student CRUD APIs

- Create Student
- Get All Students
- Get Student By ID
- Update Student
- Delete Student

---

## ✅ Middleware

- JWT Authentication Middleware
- Role Middleware
- Logger Middleware

---

## ✅ Validation

- Request DTO Validation
- User-Friendly Validation Errors

---

# 🔐 Authentication APIs

## Register User

```http
POST /api/auth/register
```

### Request Body

```json
{
  "name": "Dnyanesh",
  "email": "dnyanesh@gmail.com",
  "password": "1234",
  "role": "Admin"
}
```

---

## Login User

```http
POST /api/auth/login
```

### Request Body

```json
{
  "email": "dnyanesh@gmail.com",
  "password": "1234"
}
```

### Response

```json
{
  "message": "Login successful",
  "token": "JWT_TOKEN"
}
```

---

# 📚 Student APIs

## Create Student

```http
POST /api/students
```

### Headers

```text
Authorization: Bearer JWT_TOKEN
```

### Request Body

```json
{
  "name": "Kiran",
  "email": "kiran@gmail.com",
  "age": 24,
  "course": "Go Backend",
  "city": "Pune"
}
```

---

## Get All Students

```http
GET /api/students
```

---

## Get Student By ID

```http
GET /api/students/1
```

---

## Update Student

```http
PUT /api/students/1
```

---

## Delete Student

```http
DELETE /api/students/1
```

---

# 🔒 Protected Routes

These APIs require JWT Token:

- Create Student
- Update Student
- Delete Student
- Get Students

---

# 👨‍💻 Role-Based Authorization

| Role | Access |
|------|--------|
| Admin | Create/Delete Student |
| User | Read Only |

---

# 🗄 Database

## SQL Server using Docker

### Run SQL Server Container

```bash
docker run --platform linux/amd64 \
-e "ACCEPT_EULA=Y" \
-e "MSSQL_SA_PASSWORD=StrongPass@123" \
-p 1433:1433 \
--name sqlserver \
-d mcr.microsoft.com/mssql/server:2022-latest
```

---

# ▶️ Run Project

## Install Dependencies

```bash
go mod tidy
```

---

## Start Project

```bash
go run cmd/main.go
```

---

# 🌐 Server URL

```text
http://localhost:8080
```

---

# 🧪 API Testing

This project APIs are tested using:

- Postman

---

# 📌 Important Golang Concepts Used

- Structs
- Interfaces
- Packages
- Middleware
- JWT Authentication
- bcrypt Password Hashing
- Clean Architecture
- Repository Pattern
- Service Layer
- DTO Validation
- SQL Server Integration

---

# 🎯 Interview Questions & Answers

## 1. Why Gin Framework?

### Answer:
Gin is a lightweight and high-performance HTTP web framework in Go. It provides routing, middleware support, JSON validation, and better performance than the standard net/http package.

---

## 2. Why GORM?

### Answer:
GORM is an ORM library for Golang used to interact with databases easily using structs and methods instead of writing raw SQL queries.

---

## 3. What is JWT Authentication?

### Answer:
JWT (JSON Web Token) is a secure token-based authentication mechanism used to verify users without storing session data on the server.

---

## 4. Why use bcrypt?

### Answer:
bcrypt is used for password hashing to securely store passwords in the database.

---

## 5. What is Middleware in Gin?

### Answer:
Middleware is a function executed before or after the request handler. It is used for authentication, logging, error handling, etc.

---

## 6. What is Repository Pattern?

### Answer:
Repository Pattern separates database logic from business logic and improves maintainability and scalability.

---

## 7. What is Service Layer?

### Answer:
Service Layer contains business logic and communicates between controllers and repositories.

---

## 8. What is DTO?

### Answer:
DTO (Data Transfer Object) is used to validate and transfer request/response data securely.

---

## 9. Difference Between Authentication and Authorization?

| Authentication | Authorization |
|---------------|---------------|
| Verifies user identity | Checks user permissions |
| Login process | Access control |

---

## 10. Why use Clean Architecture?

### Answer:
Clean Architecture improves code maintainability, scalability, and testability by separating concerns into layers.

---

# 🚀 Future Improvements

- Pagination
- Search & Filter APIs
- Redis Caching
- Docker Compose
- Swagger Documentation
- CI/CD Pipeline
- Unit Testing
- Microservices

---

# 👨‍💻 Author

## Dnyaneshwar Kokate

- Golang Backend Developer
- .NET Full Stack Developer

### GitHub

https://github.com/Dnyanesh0902

### LinkedIn

https://www.linkedin.com/in/dnyaneshwar-kokate-04a12b258/

### Portfolio

https://dnyanesh.miracledevelopers.in/

---

⭐ If you like this project, give it a star on GitHub.