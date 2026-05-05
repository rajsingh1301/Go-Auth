<div align="center">
  <h1>🛡️ Go Auth API</h1>
  <p>A modern, robust, and secure authentication REST API built with Go, Gin, GORM, and PostgreSQL.</p>
  
  [![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
  [![Gin Framework](https://img.shields.io/badge/Gin-Web_Framework-00ADD8?style=for-the-badge)](https://gin-gonic.com/)
  [![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Supabase-336791?style=for-the-badge&logo=postgresql)](https://supabase.com/)
  [![JWT](https://img.shields.io/badge/JWT-JSON_Web_Tokens-000000?style=for-the-badge&logo=jsonwebtokens)](https://jwt.io/)
</div>

<br />

## ✨ Features

- **User Registration**: Secure password hashing using `bcrypt`.
- **User Authentication**: Login system that issues secure JSON Web Tokens (JWT).
- **HttpOnly Cookies**: Prevents XSS attacks by storing JWTs securely in HttpOnly cookies with `SameSite=Lax` mode.
- **Protected Routes**: Custom middleware to parse and validate JWT cookies and protect API endpoints.
- **Database Auto-Migration**: Syncs your Go structs with your PostgreSQL database seamlessly via `GORM`.

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Framework:** [Gin](https://gin-gonic.com/)
- **ORM:** [GORM](https://gorm.io/)
- **Database:** PostgreSQL (Hosted on Supabase)
- **Security:** `bcrypt`, `golang-jwt/jwt`
- **Hot Reloading:** `CompileDaemon`

## 🚀 Getting Started

### Prerequisites
- Go installed on your machine.
- A PostgreSQL database (e.g., Supabase, local Postgres).

### 1. Clone the repository
```bash
git clone <repository-url>
cd auth-go-test
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Setup Environment Variables
Create a `.env` file in the root of the project:
```env
PORT=3000
DB="host=db.your-supabase-host.com user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
SECRET="your-super-secret-jwt-key"
```

### 4. Run the Development Server
This project uses `CompileDaemon` for hot-reloading. You can run the app using the daemon:

```bash
CompileDaemon --command="./auth-go-test"
```

Or using standard Go run:
```bash
go run main.go
```

## 📡 API Endpoints

### Public Routes
- `POST /signup` - Register a new user with `email` and `password` in the JSON body.
- `POST /login` - Authenticate a user and receive an HttpOnly JWT cookie.

### Protected Routes (Requires Login)
- `GET /validate` - Checks if the user is currently authenticated and returns a success message.

## 🔒 Security Measures
- Passwords are never stored in plain text.
- JWT tokens are issued with an expiry date (30 days).
- Tokens are delivered via `HttpOnly` cookies, preventing client-side JavaScript access and potential XSS attacks.

---

<div align="center">
  <i>Built with ❤️ for secure web applications.</i>
</div>
