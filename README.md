Go Gin Project Template
A clean, structured template for building web applications with Go and the Gin framework.
Project Structure
.
├── .github/                 # GitHub specific files (workflows, templates)
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
└── README.md                # Project documentation
Getting Started
Prerequisites

Go 1.24+
Git

Installation

Clone this repository:
bashgit clone https://github.com/loiclaborderie/go-gin-template.git
cd go-gin-template

Install dependencies:
bashgo mod tidy

Run the application:
bashgo run cmd/main/main.go


Features

Clean, layered architecture
Configuration management
Routing with Gin
Example handlers and models
Test structure
GitHub Actions CI/CD

Usage as a Template
When creating a new repository on GitHub, you can select this repository as a template to start with the same directory structure and files.
Customization
Update the following files to customize the template for your project:

go.mod: Change the module name
README.md: Update project details
config/config.yaml: Set your application configuration

License
MIT License - See LICENSE file for details
Contributing
Contributions are welcome! Please feel free to submit a Pull Request.