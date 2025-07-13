# Sydney Health Clone - Development Guide

## Prerequisites

### Required Software

- **Go** 1.21+ ([Download](https://go.dev/dl/))
- **Node.js** 16+ and npm ([Download](https://nodejs.org/))
- **Docker** & Docker Compose ([Download](https://www.docker.com/))
- **Kubernetes** (minikube or Docker Desktop)
- **Protocol Buffers** compiler ([Download](https://github.com/protocolbuffers/protobuf/releases))
- **Git** ([Download](https://git-scm.com/))

### Platform-Specific Requirements

#### iOS Development
- macOS 12+ 
- Xcode 15+
- CocoaPods 1.12+

#### Android Development
- Android Studio Hedgehog (2023.1.1)+
- JDK 17
- Android SDK 34

## Initial Setup

### 1. Clone the Repository
```bash
git clone https://github.com/sydney-health/sydney-health-clone.git
cd sydney-health-clone
```

### 2. Install Dependencies

#### Backend Dependencies
```bash
# Install Go dependencies
cd backend
go mod download

# Install protoc plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

#### Web Dependencies
```bash
cd web
npm install
```

#### iOS Dependencies
```bash
cd ios
pod install
```

#### Android Dependencies
```bash
cd android
./gradlew build
```

### 3. Generate Protocol Buffers
```bash
cd shared/proto
./build.sh
```

### 4. Set Up Local Infrastructure

#### Start Core Services
```bash
# Start databases and message queue
docker-compose -f docker-compose.infrastructure.yml up -d

# Start Kafka
docker-compose -f backend/docker-compose.kafka.yml up -d

# Start observability stack
docker-compose -f devops/observability/docker-compose.observability.yml up -d
```

#### Create Database Schema
```bash
# MySQL
mysql -h localhost -u root -p < backend/migrations/001_initial_schema.sql

# PostgreSQL
psql -h localhost -U postgres -f backend/migrations/001_initial_schema.sql
```

## Development Workflow

### Backend Development

#### Running Services Locally

1. **Start the Gateway Service**
```bash
cd backend/services/gateway
go run cmd/main.go -config ../../config/gateway.yaml
```

2. **Start Individual Microservices**
```bash
# Member Service
cd backend/services/member
go run cmd/main.go -port 50051

# Benefits Service
cd backend/services/benefits
go run cmd/main.go -port 50052

# Claims Service
cd backend/services/claims
go run cmd/main.go -port 50054

# Provider Service
cd backend/services/provider
go run cmd/main.go -port 50053

# Messaging Service
cd backend/services/messaging
go run cmd/main.go -port 50055
```

#### Using Docker Compose
```bash
cd backend
docker-compose up
```

#### Running Tests
```bash
# Run all backend tests
cd backend
go test ./...

# Run with coverage
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Web Development

#### Development Server
```bash
cd web
npm run dev
# Access at http://localhost:3000
```

#### Building for Production
```bash
npm run build
npm start
```

#### Running Tests
```bash
npm test
npm run test:coverage
```

### iOS Development

1. Open the workspace in Xcode:
```bash
cd ios
open SydneyHealth.xcworkspace
```

2. Select your development team in project settings
3. Choose a simulator or device
4. Press Cmd+R to build and run

#### Running Tests
```bash
xcodebuild test -workspace SydneyHealth.xcworkspace \
  -scheme SydneyHealth \
  -destination 'platform=iOS Simulator,name=iPhone 15'
```

### Android Development

1. Open the project in Android Studio:
```bash
cd android
studio .
```

2. Sync Gradle files
3. Select an emulator or device
4. Click Run or press Shift+F10

#### Running Tests
```bash
./gradlew test
./gradlew connectedAndroidTest
```

## Configuration

### Environment Variables

Create `.env` files for local development:

#### Backend (.env)
```env
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=health_user
DB_PASSWORD=health_pass
DB_NAME=sydney_health

# Kafka
KAFKA_BROKERS=localhost:9092
KAFKA_GROUP_ID=health-dev

# JWT
JWT_SECRET=your-dev-secret-key
JWT_DURATION=3600

# Service Ports
GATEWAY_PORT=8080
MEMBER_SERVICE_PORT=50051
BENEFITS_SERVICE_PORT=50052
CLAIMS_SERVICE_PORT=50054
PROVIDER_SERVICE_PORT=50053
MESSAGING_SERVICE_PORT=50055
```

#### Web (.env.local)
```env
FUSION_API_URL=http://localhost:8080/api/v1
FUSION_WS_URL=ws://localhost:8080/ws
```

### Service Configuration

Backend services use YAML configuration files:

```yaml
# backend/config/gateway.yaml
server:
  port: 8080
  environment: development
  log_level: debug

database:
  driver: postgres
  host: localhost
  port: 5432
  # ... other settings
```

## Common Development Tasks

### Adding a New API Endpoint

1. **Define the protocol buffer**
```protobuf
// shared/proto/new_service.proto
service NewService {
  rpc NewMethod(NewRequest) returns (NewResponse);
}
```

2. **Generate code**
```bash
cd shared/proto
./build.sh
```

3. **Implement the service**
```go
// backend/services/new/internal/service/service.go
func (s *Service) NewMethod(ctx context.Context, req *pb.NewRequest) (*pb.NewResponse, error) {
    // Implementation
}
```

4. **Add gateway route**
```go
// backend/services/gateway/internal/proxy/proxy.go
api.HandleFunc("/new-endpoint", proxy.NewMethod).Methods("POST")
```

### Adding a New RIB (Mobile)

#### iOS
```swift
// Create Builder, Interactor, Router, and optionally View
// Following the RIBs pattern
```

#### Android
```kotlin
// Create Builder, Interactor, Router, and optionally View
// Following the RIBs pattern
```

### Database Migrations

1. Create migration file:
```sql
-- backend/migrations/002_add_new_table.sql
CREATE TABLE new_table (
    id INT PRIMARY KEY,
    -- columns
);
```

2. Apply migration:
```bash
migrate -path backend/migrations -database "postgres://..." up
```

## Debugging

### Backend Debugging

#### Using Delve
```bash
# Install Delve
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug a service
cd backend/services/gateway
dlv debug ./cmd/main.go -- -config ../../config/gateway.yaml
```

#### Logging
```go
import "github.com/sydney-health-clone/backend/shared/logger"

logger.Info("Processing request", 
    zap.String("member_id", memberID),
    zap.Int("count", count),
)
```

### Web Debugging

- Use Chrome DevTools
- Enable React Developer Tools
- Check Fusion.js debug mode: `fusion dev --debug`

### Mobile Debugging

#### iOS
- Use Xcode debugger
- Enable View Debugging
- Check Charles Proxy for network requests

#### Android
- Use Android Studio debugger
- Enable Layout Inspector
- Use Stetho for network debugging

## Performance Profiling

### Backend Profiling
```go
import _ "net/http/pprof"

// In main.go
go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
}()
```

Access profiles at:
- CPU: http://localhost:6060/debug/pprof/profile
- Memory: http://localhost:6060/debug/pprof/heap
- Goroutines: http://localhost:6060/debug/pprof/goroutine

### Web Performance
```bash
# Run Lighthouse
npm run lighthouse

# Bundle analysis
npm run analyze
```

## Troubleshooting

### Common Issues

#### Port Already in Use
```bash
# Find process using port
lsof -i :8080

# Kill process
kill -9 <PID>
```

#### Database Connection Failed
- Check if Docker containers are running
- Verify credentials in configuration
- Check firewall settings

#### Proto Generation Failed
- Ensure protoc is installed and in PATH
- Check proto syntax errors
- Verify all imports exist

#### Build Failures

##### Go
```bash
# Clear module cache
go clean -modcache

# Update dependencies
go mod tidy
```

##### Node.js
```bash
# Clear npm cache
npm cache clean --force

# Reinstall dependencies
rm -rf node_modules package-lock.json
npm install
```

### Getting Help

1. Check existing issues on GitHub
2. Ask in development Slack channel
3. Consult team documentation wiki
4. Contact technical lead

## Best Practices

### Code Quality

1. **Run linters before committing**
```bash
# Go
golangci-lint run

# JavaScript
npm run lint

# Swift
swiftlint

# Kotlin
./gradlew ktlintCheck
```

2. **Write tests for new features**
- Aim for 80% code coverage
- Include unit and integration tests
- Test error cases

3. **Follow coding standards**
- Go: Effective Go guidelines
- JavaScript: Airbnb style guide
- Swift: Swift API Design Guidelines
- Kotlin: Kotlin Coding Conventions

### Git Workflow

1. Create feature branch
```bash
git checkout -b feature/JIRA-123-new-feature
```

2. Make atomic commits
```bash
git add -p
git commit -m "feat: add new provider search filter"
```

3. Keep branch updated
```bash
git fetch origin
git rebase origin/main
```

4. Create pull request with description

### Security

- Never commit secrets or API keys
- Use environment variables for configuration
- Validate all user inputs
- Follow OWASP guidelines
- Run security scans regularly

## Resources

### Documentation
- [Architecture Overview](./ARCHITECTURE.md)
- [API Documentation](./API.md)
- [Deployment Guide](./DEPLOYMENT.md)

### External Resources
- [RIBs Documentation](https://github.com/uber/RIBs)
- [Fusion.js Guide](https://fusionjs.com)
- [gRPC Documentation](https://grpc.io/docs/)
- [Kubernetes Documentation](https://kubernetes.io/docs/)

### Tools
- [Postman Collection](./postman/sydney-health.json)
- [K6 Performance Tests](./tests/performance/)
- [Mock Data Generator](./tools/mock-data-generator/)