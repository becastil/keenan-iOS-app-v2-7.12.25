# Sydney Health Clone

A cross-platform healthcare application built with Uber's technology stack, replicating the functionality of Anthem Sydney Health app.

## 🚀 Overview

Sydney Health Clone is a comprehensive healthcare management platform that enables members to manage their health benefits, find care providers, track claims, and communicate securely with support teams. Built using modern microservices architecture and native mobile frameworks, it demonstrates enterprise-grade healthcare application development.

## 📋 Key Features

- **Benefits Management**: View medical, dental, vision, and pharmacy coverage details
- **Provider Search**: Find in-network healthcare providers with real-time availability
- **Digital ID Cards**: Access and share member ID cards digitally
- **Claims Tracking**: Monitor claim status from submission to payment
- **Cost Estimation**: Get procedure cost estimates before treatment
- **Secure Messaging**: Communicate with care coordinators and support teams
- **Real-time Updates**: Receive notifications for claim status changes and messages

## 🏗️ Architecture

### Technology Stack

- **Web**: Fusion.js (React-based framework by Uber)
- **iOS**: Swift with RIBs architecture
- **Android**: Kotlin with RIBs architecture
- **Backend**: Go microservices with gRPC
- **Data Layer**: 
  - MySQL/PostgreSQL (transactional data)
  - Apache Kafka (event streaming)
  - Apache Pinot (analytics)
  - Redis (caching)
- **Infrastructure**:
  - Kubernetes (container orchestration)
  - Spinnaker (CI/CD)
  - M3 + Prometheus (metrics)
  - Jaeger (distributed tracing)

### Project Structure

```
sydney-health-clone/
├── web/                    # Fusion.js web application
├── ios/                    # iOS app (Swift + RIBs)
├── android/                # Android app (Kotlin + RIBs)
├── backend/                # Go microservices
│   ├── services/          # Individual microservices
│   │   ├── gateway/       # API Gateway
│   │   ├── member/        # Member management
│   │   ├── benefits/      # Benefits information
│   │   ├── claims/        # Claims processing
│   │   ├── provider/      # Provider search
│   │   └── messaging/     # Secure messaging
│   ├── shared/            # Shared utilities
│   └── migrations/        # Database schemas
├── shared/                 # Cross-platform shared code
│   ├── proto/             # gRPC definitions
│   └── models/            # Shared data models
├── devops/                 # DevOps configuration
│   ├── spinnaker/         # CI/CD pipelines
│   └── observability/     # Monitoring stack
└── docs/                   # Technical documentation
```

## 🚀 Quick Start

### Prerequisites

- Go 1.21+
- Node.js 16+
- Docker & Docker Compose
- Protocol Buffers compiler

### Development Setup

```bash
# Clone the repository
git clone https://github.com/sydney-health/sydney-health-clone.git
cd sydney-health-clone

# Install dependencies
npm install

# Start infrastructure services
docker-compose -f docker-compose.infrastructure.yml up -d

# Generate protocol buffers
cd shared/proto && ./build.sh && cd ../..

# Run backend services
cd backend
docker-compose up -d

# Run web application
cd ../web
npm run dev

# Access the application
open http://localhost:3000
```

### Running Tests

```bash
# Backend tests
cd backend && go test ./...

# Web tests
cd web && npm test

# iOS tests
cd ios && xcodebuild test -workspace SydneyHealth.xcworkspace -scheme SydneyHealth

# Android tests
cd android && ./gradlew test
```

## 📚 Documentation

### Core Documentation
- **[Architecture Overview](./docs/ARCHITECTURE.md)** - System design and technology decisions
- **[API Documentation](./docs/API.md)** - Complete API reference with examples
- **[Development Guide](./docs/DEVELOPMENT.md)** - Setup, workflow, and best practices
- **[Deployment Guide](./docs/DEPLOYMENT.md)** - Production deployment procedures

### Component Guides
- [Web Development](./web/README.md)
- [iOS Development](./ios/README.md)
- [Android Development](./android/README.md)
- [Backend Development](./backend/README.md)

## 🔒 Security

- HIPAA compliant architecture
- End-to-end encryption for sensitive data
- JWT-based authentication with refresh tokens
- Biometric authentication on mobile devices
- Regular security audits and penetration testing

## 🎯 Performance

- Sub-second API response times
- Offline-first mobile architecture
- CDN-distributed static assets
- Horizontal scaling with Kubernetes
- Real-time updates via WebSockets

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guidelines](./docs/CONTRIBUTING.md) for details.

### Development Workflow

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📊 Project Status

- ✅ Core architecture implemented
- ✅ Basic features functional with mock data
- ✅ CI/CD pipeline configured
- ✅ Monitoring and observability setup
- 🚧 Additional features in development
- 📅 Production deployment planned

## 🙏 Acknowledgments

- Built with Uber's open-source technologies
- Inspired by Anthem Sydney Health app
- Mock data for demonstration purposes only

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Contact & Support

- **Documentation**: [docs/](./docs/)
- **Issues**: [GitHub Issues](https://github.com/sydney-health/sydney-health-clone/issues)
- **Discussions**: [GitHub Discussions](https://github.com/sydney-health/sydney-health-clone/discussions)

---

**Note**: This is a demonstration project using mock data. Not intended for production healthcare use.