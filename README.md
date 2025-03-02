# Knowledge Hub

Knowledge Hub is a multi-tenant knowledge sharing platform built with modern technologies. This project demonstrates the capabilities of [Roo Code](https://github.com/RooInc/roo-code) (previously known as Roo Cline), an AI-powered development tool.

## Repository Purpose

This repository was created with the primary goal of **building a functional web application using Roo Code**. While the programming languages and Docker usage were specified, other aspects (libraries, architecture design, coding style, etc.) are **100% AI-generated outputs adopted without modification**.

Please note the following important points:
- The application may contain numerous bugs
- The architecture, code structure, and coding style are **not intended to serve as reference material**
- Ideally, **human guidance should provide detailed instructions and corrections**

The content of this repository is intended for testing and experimental use of Roo Code and does not guarantee production-level quality or best practices.

## Project Structure

```
.
├── frontend/          # Next.js frontend application
├── backend/           # Go/Echo backend application
├── docker/            # Docker related files
├── docs/              # Project documentation
├── .clinerules        # Project coding rules
└── clinelogs/         # Development process logs
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
docker compose -f compose-db-tools.yaml run --rm migrate /go/src/github.com/hyorimitsu/knowledge-hub/backend/scripts/migrate.sh up

# Rollback migrations
docker compose -f compose-db-tools.yaml run --rm migrate /go/src/github.com/hyorimitsu/knowledge-hub/backend/scripts/migrate.sh down

# Check migration version
docker compose -f compose-db-tools.yaml run --rm migrate /go/src/github.com/hyorimitsu/knowledge-hub/backend/scripts/migrate.sh version
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

# Frontend development commands
# Install npm packages
docker compose -f compose-tools.yaml run --rm node npm install

# Add a specific npm package
docker compose -f compose-tools.yaml run --rm node npm install package-name

# Run npm scripts
docker compose -f compose-tools.yaml run --rm node npm run script-name
```

5. Access the applications
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080
- API Documentation: http://localhost:8080/swagger/

### Troubleshooting

#### API Connectivity Issues

If you encounter 404 errors when the frontend tries to connect to the backend API:

1. Make sure both the frontend and backend servers are running
2. Check that the Next.js proxy configuration in `frontend/next.config.ts` is correctly set up:
   ```typescript
   async rewrites() {
     return [
       {
         source: '/api/:path*',
         destination: 'http://localhost:8080/api/:path*',
       },
     ];
   }
   ```
3. Ensure the backend CORS configuration in `backend/cmd/api/main.go` allows requests from the frontend:
   ```go
   e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
     AllowOrigins: []string{"http://localhost:3000"},
     AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
     AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
     AllowCredentials: true,
   }))
   ```

## Project Status
🚧 Currently under development

## License
This project is open source and available under the MIT License.