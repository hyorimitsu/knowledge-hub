# Knowledge Hub

Knowledge Hub is a multi-tenant knowledge sharing platform built with modern technologies.

This project is primarily generated using [Roo Code](https://github.com/RooInc/roo-code) (previously known as Roo Cline), an AI-powered development tool.

## Project Structure

```
.
├── frontend/          # Next.js frontend application
├── backend/           # Go/Echo backend application
├── docker/            # Docker related files
├── docs/             # Project documentation
├── .clinerules       # Project coding rules
└── clinelogs/        # Development process logs
```

## Technology Stack

### Frontend
- Next.js (App Router)
- TypeScript
- ESLint
- Prettier

### Backend
- Go
- Echo Framework
- OpenAPI 3.1
- Gorm
- golangci-lint

### Database
- PostgreSQL (Multi-tenant architecture)

### Infrastructure
- Docker
- Docker Compose

## Getting Started

### Prerequisites
- Docker
- Docker Compose
- Node.js (for local frontend development)
- Go (for local backend development)

### Development Setup

1. Clone the repository
```bash
git clone https://github.com/hyorimitsu/knowledge-hub.git
cd knowledge-hub
```

2. Start the development environment
```bash
# Start all services
docker compose up -d

# View logs
docker compose logs -f
```

3. Database Operations
```bash
# Run database migrations
docker compose -f compose-db-tools.yaml run --rm migrate up

# Rollback migrations
docker compose -f compose-db-tools.yaml run --rm migrate down

# Check migration version
docker compose -f compose-db-tools.yaml run --rm migrate version
```

4. Development Commands
```bash
# Run tests
docker compose -f compose-tools.yaml run --rm test

# Run linter
docker compose -f compose-tools.yaml run --rm lint

# Generate OpenAPI documentation
docker compose -f compose-tools.yaml run --rm swag

# Update Go dependencies
docker compose -f compose-tools.yaml run --rm tidy

# Vendor dependencies
docker compose -f compose-tools.yaml run --rm vendor

# Format Go code
docker compose -f compose-tools.yaml run --rm fmt
```

5. Access the applications
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080
- API Documentation: http://localhost:8080/swagger/

## Project Status
🚧 Currently under development

## License
This project is open source and available under the MIT License.