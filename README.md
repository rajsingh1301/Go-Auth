<div align="center">

  <img src="assets/gopher.png" alt="Go Gopher" width="200" />

  <h1>🛡️ Go Auth API</h1>
  <p><b>A modern, robust, and secure authentication REST API built with Go, Gin, GORM, and PostgreSQL.</b></p>
  
  [![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)](https://go.dev/)
  [![Gin Framework](https://img.shields.io/badge/Gin-Web_Framework-00ADD8?style=for-the-badge)](https://gin-gonic.com/)
  [![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Supabase-336791?style=for-the-badge&logo=postgresql)](https://supabase.com/)
  [![JWT](https://img.shields.io/badge/JWT-JSON_Web_Tokens-000000?style=for-the-badge&logo=jsonwebtokens)](https://jwt.io/)
</div>

<br />

---

## ✨ Core Features

| Feature | Description |
| :--- | :--- |
| 🔐 **User Registration** | Secure password hashing out of the box using `bcrypt`. |
| 🔑 **Authentication** | Login system that issues secure JSON Web Tokens (JWT). |
| 🍪 **HttpOnly Cookies** | Prevents XSS attacks by storing JWTs securely in HttpOnly cookies with `SameSite=Lax` mode. |
| 🛡️ **Protected Routes** | Custom middleware to parse and validate JWT cookies to protect API endpoints. |
| 📦 **Auto-Migration** | Syncs your Go structs with your PostgreSQL database seamlessly via `GORM`. |

---

## 🛠️ Tech Stack & Architecture

This project is built on top of modern, battle-tested tools to ensure high performance and top-tier security:

- **Language:** Go (Golang) for blazingly fast execution.
- **Framework:** [Gin Web Framework](https://gin-gonic.com/) for routing and middleware.
- **ORM:** [GORM](https://gorm.io/) for developer-friendly database interactions.
- **Database:** PostgreSQL (Hosted natively on Supabase).
- **Security:** `bcrypt` for hashing, `golang-jwt/jwt` for tokens.
- **Hot Reloading:** `CompileDaemon` for an optimized development experience.

---

## 🚀 Getting Started

### Prerequisites
Before you begin, ensure you have the following installed:
- [Go (Golang)](https://go.dev/doc/install)
- A PostgreSQL database (e.g., [Supabase](https://supabase.com/), local Postgres).

### 1️⃣ Clone the repository
```bash
git clone https://github.com/rajsingh1301/Go-Auth.git
cd Go-Auth
```

### 2️⃣ Install Dependencies
```bash
go mod tidy
```

### 3️⃣ Setup Environment Variables
Create a `.env` file in the root of the project with your database and JWT secrets:
```env
PORT=3000
DB="host=db.your-supabase-host.com user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
SECRET="your-super-secret-jwt-key"
```

### 4️⃣ Run the Development Server
This project utilizes `CompileDaemon` for hot-reloading. You can start the app using the daemon:

```bash
CompileDaemon --command="./auth-go-test"
```

Or run it using the standard Go command:
```bash
go run main.go
```

---

## 📡 API Endpoints Reference

### 🌐 Public Routes
| Method | Endpoint | Description | Body Parameters |
| :---: | :--- | :--- | :--- |
| `POST` | `/signup` | Register a new user | `email`, `password` |
| `POST` | `/login` | Authenticate a user | `email`, `password` |

### 🔒 Protected Routes (Requires Auth Cookie)
| Method | Endpoint | Description |
| :---: | :--- | :--- |
| `GET` | `/validate` | Validates JWT & returns auth status |

---

## 🔒 Security Posture
- **No Plain Text Passwords:** All passwords are salted and hashed using standard `bcrypt` algorithms.
- **Token Expiry:** JWT tokens are strictly issued with a 30-day expiration timeframe.
- **Cookie Delivery:** Tokens are strictly delivered via `HttpOnly` cookies, fully preventing client-side JavaScript access (mitigating XSS).

---

<div align="center">
  <i>Engineered with ❤️ for secure web applications.</i>
</div>
