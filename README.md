# Go Gin Project Template

A clean, structured template for building web applications with Go and the Gin framework.

## 📂 Project Structure

```
.
├── .github/                 # GitHub workflows for CI/CD
├── .gitignore               # Git ignore file
├── cmd/                     # Main applications
│   └── main/                # The main application entry point
│       └── main.go          # Main function
├── config/                  # Configuration files
│   └── config.yaml          # Application configuration
├── handlers/                # HTTP request handlers
│   └── hello.go             # Example handler
├── internal/                # Private application code
│   ├── middleware/          # HTTP middleware
│   └── utils/               # Utility functions
├── models/                  # Data models
│   └── user.go              # User model
├── routes/                  # Route definitions
│   └── routes.go            # Route setup
├── test/                    # Test files
│   └── unit/                # Unit tests
│       └── hello_test.go    # Tests for hello handler
├── go.mod                   # Go module definition
├── go.sum                   # Go module checksums
├── Makefile                 # Build automation
├── Dockerfile               # Docker container definition
└── README.md                # Project documentation
```

## ✨ Features

- **Clean, Layered Architecture**: Separation of concerns for better maintainability
- **Configuration Management**: YAML-based application configuration
- **Routing with Gin**: Fast and flexible HTTP router
- **Example Handlers and Models**: Ready-to-use examples for quick development
- **Test Structure**: Set up for unit and integration testing
- **GitHub Actions CI/CD**: Automated testing and deployment
- **Docker Support**: Easy containerization for deployment

## 🚀 Getting Started

### Prerequisites

- Go 1.24+
- Git

### Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/loiclaborderie/bahasa-project.git
   cd go-gin-template
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run cmd/main/main.go
   ```
   
   Alternatively, use the Makefile:
   ```bash
   make run
   ```

## 🔧 Development

### Build the Application

```bash
make build
```

### Run Tests

```bash
make test
```

### Run Linters

```bash
make lint
```

### Build Docker Image

```bash
make docker
```

## 🧩 Usage as a Template

When creating a new repository on GitHub, you can select this repository as a template to start with the same directory structure and files.

### Customization

Update the following files to customize the template for your project:

1. **`go.mod`**: Change the module name
   ```go
   module github.com/your-username/your-project
   ```

2. **`README.md`**: Update project details

3. **`config/config.yaml`**: Set your application configuration

## 📝 API Endpoints

The template comes with several pre-configured endpoints:

- **GET `/health`**: Health check
- **GET `/api/v1/hello`**: Hello world example
- **GET `/api/v1/users`**: Get all users
- **GET `/api/v1/users/:id`**: Get user by ID
- **POST `/api/v1/users`**: Create new user
- **PUT `/api/v1/users/:id`**: Update user
- **DELETE `/api/v1/users/:id`**: Delete user

## 📄 License

MIT License - See License file file for details

## 👥 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
