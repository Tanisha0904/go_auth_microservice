Go JWT Authentication MicroserviceA lightweight, high-performance authentication microservice built with Go and the Gin Gonic web framework. This service demonstrates how to implement secure user login and route protection using JSON Web Tokens (JWT).🚀 FeaturesJWT Generation: Securely signs tokens using the HMAC-SHA256 algorithm.Custom Middleware: A "Gatekeeper" that validates tokens for protected API routes.Route Grouping: Separates public endpoints (Login) from private ones (Dashboard).Type-Safe Claims: Uses custom JWT claims to store user-specific data.📂 Project StructurePlaintextauth-service/
├── cmd/
│   └── main.go          # Entry point (Wires up routes and starts server)
├── internal/
│   ├── database/        
│   │   └── database.go  # NEW: DB connection logic
│   ├── auth/            
│   │   ├── handler.go   # HTTP Login handler and request binding
│   │   ├── service.go   # JWT generation logic
│   │   └── model.go     # User and JWT claim structures
│   └── middleware/
│       └── auth.go      # JWT validation middleware logic
├── go.mod               # Go module dependencies
└── go.sum               # Dependency checksums
🛠️ Setup & Installation1. PrerequisitesGo 1.20 or higher installed.Postman (for testing).2. Clone and InstallBash# Clone the repository
git clone https://github.com/yourusername/auth-service.git
cd auth-service

# Install dependencies
go mod tidy
3. Run the ServiceBashgo run cmd/main.go
The server will start on http://localhost:8080.🧪 API Endpoints & Testing1. User Login (Public)Exchange credentials for a JWT token.URL: /loginMethod: POSTBody (JSON):JSON{
    "username": "admin",
    "password": "password123"
}
2. Dashboard (Protected)Access restricted data using the token received from login.URL: /api/dashboardMethod: GETHeader: Authorization: Bearer <YOUR_TOKEN_HERE>🔑 Technical Implementation DetailsProcessDescriptionAuthenticationValidates user credentials and issues a signed JWT string.AuthorizationMiddleware intercepts requests to /api/*, parsing the Authorization header.HMAC SigningUses a secret []byte key to sign and verify token integrity.Context SharingMiddleware extracts the username from the token and stores it in gin.Context for use by handlers.