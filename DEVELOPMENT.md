# Sydney Health Development Guide

This guide provides comprehensive instructions for setting up and working with the Sydney Health mobile application development environment.

## 📋 Prerequisites

Before you begin, ensure you have the following installed:

### Required Tools
- **Node.js** (v18 or higher) - [Download](https://nodejs.org/)
- **Go** (v1.21 or higher) - [Download](https://go.dev/dl/)
- **Docker & Docker Compose** - [Download](https://www.docker.com/products/docker-desktop/)
- **Git** - [Download](https://git-scm.com/downloads)

### Platform-Specific Requirements

#### For Android Development
- **Java 11+** (OpenJDK recommended)
- **Android Studio** or Android SDK Tools
- **ANDROID_HOME** environment variable set

#### For iOS Development (macOS only)
- **Xcode** (latest version)
- **CocoaPods** - Install with: `sudo gem install cocoapods`
- **Swift** (comes with Xcode)

#### For Protocol Buffers
- **protoc** compiler (will be installed by setup script)
- **protoc plugins** for each language (installed by setup)

## 🚀 Quick Start

### 1. Initial Setup

Run the automated setup script to configure your development environment:

```bash
./setup-dev-environment.sh
```

This script will:
- Install Protocol Buffers compiler and plugins
- Generate Protocol Buffer code for all platforms
- Configure Android environment (gradle wrapper)
- Set up iOS environment (Info.plist, pods)
- Create backend configuration files
- Set up web app environment

### 2. Start Infrastructure

Start the Docker-based development infrastructure:

```bash
./dev.sh start
```

This starts:
- PostgreSQL database (with migrations and seed data)
- Kafka message broker
- Redis cache
- Jaeger tracing
- Prometheus metrics
- Grafana dashboards

### 3. Run Services

Start the backend services:

```bash
./dev.sh backend
```

In a new terminal, start the web application:

```bash
./dev.sh web
```

### 4. Access the Application

- **Web App**: http://localhost:3000
- **API Gateway**: http://localhost:8080
- **Demo Login**: Member ID: `M123456`, Password: `demo`

## 🏗️ Project Structure

```
sydney-health-clone/
├── android/          # Android app (Kotlin, RIBs)
├── ios/             # iOS app (Swift, RIBs)
├── web/             # Web app (React, Fusion.js)
├── backend/         # Go microservices
├── shared/          # Shared Protocol Buffers
├── devops/          # Infrastructure & monitoring
├── docs/            # Project documentation
└── dev.sh          # Development helper script
```

## 🛠️ Development Commands

### Infrastructure Management

```bash
# Start all infrastructure services
./dev.sh start

# Stop infrastructure
./dev.sh stop

# Reset infrastructure (deletes all data)
./dev.sh reset

# Check service status
./dev.sh status

# View logs
./dev.sh logs              # All services
./dev.sh logs postgres     # Specific service
```

### Running Applications

```bash
# Backend services
./dev.sh backend

# Web application
./dev.sh web

# Build Android app
./dev.sh android

# Open iOS project
./dev.sh ios
```

### Code Generation & Testing

```bash
# Regenerate Protocol Buffer code
./dev.sh proto

# Run all tests
./dev.sh test

# Run linters
./dev.sh lint
```

## 📱 Platform-Specific Development

### Android Development

1. Build the app:
   ```bash
   ./dev.sh android
   ```

2. The APK will be generated at:
   ```
   android/app/build/outputs/apk/debug/app-debug.apk
   ```

3. Install on device/emulator:
   ```bash
   cd android
   ./gradlew installDebug
   ```

### iOS Development

1. Open the project:
   ```bash
   ./dev.sh ios
   ```

2. In Xcode:
   - Select your development team
   - Choose a simulator or device
   - Press ⌘R to build and run

### Web Development

The web app uses Fusion.js (Uber's React framework) with hot reloading:

```bash
./dev.sh web
```

Changes to source files will automatically reload in the browser.

## 🗄️ Database Management

### Accessing the Database

```bash
# Connect to PostgreSQL
docker exec -it sydney-health-postgres psql -U sydney_health

# Common queries
\dt                    # List tables
\d members            # Describe table
SELECT * FROM members WHERE member_id = 'M123456';
```

### Running Migrations

Migrations run automatically on startup. To run manually:

```bash
docker exec -it sydney-health-postgres psql -U sydney_health -f /docker-entrypoint-initdb.d/001_initial_schema.sql
```

## 📨 Message Queue (Kafka)

### Kafka UI

Access Kafka UI at http://localhost:8090 to:
- View topics and messages
- Monitor consumer groups
- Send test messages

### Kafka Topics

Pre-configured topics:
- `claims` - Claim processing events
- `messages` - In-app messaging
- `audit` - Audit log events
- `member-updates` - Member profile changes
- `benefit-changes` - Benefit updates

## 🔍 Monitoring & Debugging

### Service URLs

- **Jaeger UI** (Tracing): http://localhost:16686
- **Prometheus** (Metrics): http://localhost:9090
- **Grafana** (Dashboards): http://localhost:3001
  - Default login: admin/admin
- **Kafka UI**: http://localhost:8090

### Health Checks

All services expose health endpoints:
- Gateway: http://localhost:8080/health
- Metrics: http://localhost:9091/metrics

## 🧪 Testing

### Running Tests

```bash
# All tests
./dev.sh test

# Platform-specific
cd backend && go test ./...
cd web && npm test
cd android && ./gradlew test
```

### Writing Tests

- **Backend**: Use Go's built-in testing package
- **Web**: Jest for unit tests, React Testing Library
- **Android**: JUnit and Mockito
- **iOS**: XCTest framework

## 🐛 Troubleshooting

### Common Issues

#### Port Conflicts

If you get port conflict errors:

```bash
# Find what's using a port
lsof -i :8080

# Kill the process
kill -9 <PID>
```

#### Docker Issues

```bash
# Clean up Docker resources
docker system prune -a

# Reset everything
./dev.sh reset
```

#### Database Connection Issues

1. Check if PostgreSQL is running:
   ```bash
   ./dev.sh status
   ```

2. Verify connection:
   ```bash
   docker exec -it sydney-health-postgres pg_isready
   ```

#### Build Issues

- **Android**: Ensure `ANDROID_HOME` is set and gradle wrapper exists
- **iOS**: Run `pod install` in the ios directory
- **Backend**: Run `go mod download` to fetch dependencies

### Getting Help

1. Check service logs: `./dev.sh logs <service>`
2. Review error messages in the console
3. Check the GitHub issues for known problems
4. Ask in the development Slack channel

## 🔐 Security Notes

### Development Credentials

**Never use these in production:**
- Database: `sydney_health` / `dev_password`
- Demo user: `M123456` / `demo`
- JWT secret: `dev_secret_key_change_in_production_please`

### Best Practices

1. Use environment variables for sensitive data
2. Never commit `.env` files with real credentials
3. Use proper SSL/TLS in production
4. Implement proper authentication before deploying

## 📚 Additional Resources

- [Project Architecture](docs/ARCHITECTURE.md)
- [API Documentation](docs/API.md)
- [Deployment Guide](docs/DEPLOYMENT.md)
- [Protocol Buffers Guide](shared/proto/README.md)

## 🤝 Contributing

1. Create a feature branch
2. Make your changes
3. Run tests and linters
4. Submit a pull request

See [CONTRIBUTING.md](../CONTRIBUTING.md) for detailed guidelines.

---

Happy coding! 🎉 If you encounter any issues, please don't hesitate to ask for help.