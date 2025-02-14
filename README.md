# User CRUD

Implement user CRUD with Go, Gin, Gorm and PostgreSQL.

## Features

- Modular Project Structure: Organized code to maintain readability and scalability.
- Environment Management: .env configuration for managing environment variables.
- Dependency Management: Uses Go Modules for easy dependency tracking.
- Docker Support: Dockerized setup for easier deployment and consistency across environments.

## Getting Started

### Postman Collection

[Postman Collection](https://documenter.getpostman.com/view/6926243/2sAYXCkyc1)

### Prerequisites

- Go (version 1.16 or higher)
- Docker (optional for containerization)
- Make

### Installation

1. Clone the repo
   ```sh
    git clone https://github.com/feildrixliemdra/valorx-auth.git
    cd valorx-auth
   ```
2. Install dependencies
   ```sh
   go mod tidy
   ```
3. Set up environment variables:
   ```sh
    make env
    or
    cp config/config.yaml.example config/config.yaml  #fill in necessary values. #fill in necessary values.
   ```
4. Change git remote url to avoid accidental pushes to base project
   ```sh
   git remote set-url origin your_github_username/repo_name
   git remote -v # confirm the changes
   ```
5. Run the application:
   ```sh
   go run main.go serve-http
   ```
   Or Using Docker
   ```sh
   docker-compose up --build
   ```
   Access the application at http://localhost:<your_port>.

## Project Structure

```
.
├── cmd                # Main applications of the project
├── pkg                # Library code that's ok to use by external applications
├── internal           # Private application code (cannot be imported by other projects)
│   ├── config         # Application configuration and settings
│   ├── controllers    # Handle HTTP requests and responses, business logic for APIs
│   ├── models         # Data models representing entities, structs, and database schema
│   ├── repository     # Data access layer for managing database operations
│   ├── services       # Core business logic and reusable services for the application
│   ├── utils          # Utility functions and helper methods
requests and responses
├── configs            # Configuration files
├── scripts            # Scripts for various tasks
├── .env.example       # Environment variables example file
├── Dockerfile         # Docker configuration
├── docker-compose.yml # Docker Compose configuration
├── Makefile           # Automation commands
└── README.md

```

- `internal/config`

  Contains configuration files and setup functions for initializing configurations.
  It typically reads from environment variables and .env files to set up values
  for application settings, database connections, and API configurations.

- `internal/controllers`

  Contains handlers responsible for managing incoming HTTP requests and sending responses. Each controller usually corresponds to an endpoint or resource, delegating business logic to the appropriate services and handling the HTTP logic, such as parsing requests and formatting responses.

- `internal/services`

  Contains the core business logic of the application, defining reusable services and application flows. Services here interact with repositories to retrieve or update data, applying business rules and validations before returning results.

- `internal/models`

  Defines the data structures used across the application, representing entities like User, Product, or Order. These models often map directly to database tables and include struct tags to help with serialization, database ORM, or validation.

- `internal/repository`

  Acts as the Data Access Layer (DAL), managing database operations like querying, creating, updating, and deleting records. This layer encapsulates all direct interactions with the database, often leveraging interfaces to support dependency injection for testing.

- `internal/utils`

  Utility functions and helper methods that may not belong in any specific module. These could include functions for string manipulation, date formatting, error handling, or logging that are used throughout the application.

- `internal/config`

  Contains configuration files and setup functions for initializing configurations. It typically reads from environment variables and .env files to set up values for application settings, database connections, and API configurations.

## Contact

Feildrix Liemdra - feildrixliemdra@gmail.com
