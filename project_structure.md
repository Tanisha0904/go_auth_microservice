auth-service/
├── cmd/
│   └── main.go          # Entry point (Starts the server)
├── internal/
│   ├── auth/            # Core logic
│   │   ├── handler.go   # HTTP Request/Response handling
│   │   ├── service.go   # Business logic (JWT generation)
│   │   └── model.go     # Data structures (User, Claims)
│   └── middleware/
│       └── auth.go      # JWT Validation middleware
├── .env                 # Environment variables (Secrets)
├── go.mod               # Dependencies
└── go.sum