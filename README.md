# 🔐 Go JWT Authentication Microservice

> A lightweight, high-performance authentication microservice built using **Go** and the **Gin Gonic** framework.  
> Implements secure login and route protection using **JSON Web Tokens (JWT)**.

---

## 🚀 Features

> Designed to be minimal, fast, and production-ready.

- 🔑 **JWT Generation** using HMAC-SHA256
- 🛡️ **Custom Middleware (Gatekeeper)** for token validation
- 🧭 **Route Grouping** (Public vs Protected APIs)
- 🧩 **Type-Safe JWT Claims**
- ⚡ **High-Performance Go Microservice**

---

## 📂 Project Structure

```plaintext
auth-service/
├── cmd/
│   └── main.go          # Application entry point
├── internal/
│   ├── auth/
│   │   ├── handler.go   # Login HTTP handler
│   │   ├── service.go   # JWT creation logic
│   │   └── model.go     # User & JWT claim models
│   └── middleware/
│       └── auth.go      # JWT validation middleware
├── go.mod               # Go module dependencies
└── go.sum               # Dependency checksums
```

### 🛠️ Setup & Installation
✅ Prerequisites: 

Ensure the following are installed:

> Go 1.20+

> Postman / cURL for API testing

## 📥 Clone the Repository

```bash
$ git clone https://github.com/yourusername/auth-service.git
$ cd auth-service
```

### Install Dependencies
```
$ go mod tidy
```
▶️ Run the Service
```
$ go run cmd/main.go
```

### 🚀 Server will start at:
```
http://localhost:8080
```

## 🧪 API Endpoints
**🔓 Login (Public Endpoint)**

Authenticate user and receive a JWT.

**URL:** /login

**Method:** POST

**Request Body**
````
{
  "username": "admin",
  "password": "password123"
}
````

**Response**
````
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
````
### 🔐 Dashboard (Protected Endpoint)

Accessible only with a valid JWT.

**URL:** /api/dashboard

**Method:** GET

**Headers**
            **Authorization:** Bearer <YOUR_JWT_TOKEN>

## 🔑 Technical Implementation Details
Component	    Description
Authentication	Validates credentials & issues JWT
Authorization	Middleware intercepts /api/* routes
Token Signing	HMAC-SHA256 secret-based signing
Context Sharing	Username extracted & injected into gin.Context

## 🔄 Authentication Flow

Step-by-step request lifecycle:

> User submits login credentials

> Server validates credentials

> JWT token is generated & signed

> Client stores the token

> Token sent via Authorization header

> Middleware validates token

> Protected handler executes

### ⚠️ Notes

- Important considerations:

- Credentials are hardcoded for demo

- Use database + hashed passwords in production

- Store JWT secrets in environment variables

- Enable token expiration & refresh logic if needed
