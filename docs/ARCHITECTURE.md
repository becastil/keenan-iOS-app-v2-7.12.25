# Sydney Health Clone - Architecture Documentation

## Overview

Sydney Health Clone is a cross-platform healthcare application built with Uber's technology stack, designed to replicate the functionality of Anthem Sydney Health app. The system follows a microservices architecture with native mobile applications and a modern web frontend.

## System Architecture

### High-Level Architecture

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│   Web Client    │     │   iOS Client    │     │ Android Client  │
│  (Fusion.js)    │     │ (Swift + RIBs)  │     │(Kotlin + RIBs)  │
└────────┬────────┘     └────────┬────────┘     └────────┬────────┘
         │                       │                        │
         └───────────────────────┴────────────────────────┘
                                │
                    ┌───────────▼────────────┐
                    │    API Gateway         │
                    │  (Go + gRPC/HTTP)      │
                    └───────────┬────────────┘
                                │
        ┌───────────────────────┴────────────────────────┐
        │                                                │
┌───────▼────────┐  ┌──────────────┐  ┌─────────────────▼────────┐
│ Microservices  │  │    Kafka     │  │    Data Layer           │
│  - Member      │◄─┤              ├─►│  - MySQL/PostgreSQL     │
│  - Benefits    │  │              │  │  - Redis Cache          │
│  - Claims      │  └──────────────┘  │  - Apache Pinot         │
│  - Providers   │                     └──────────────────────────┘
│  - Messaging   │
└────────────────┘
```

## Technology Stack

### Backend
- **Language**: Go 1.21+
- **API Protocol**: gRPC with HTTP/JSON transcoding
- **Service Mesh**: Uber Gateway / NGINX
- **Message Queue**: Apache Kafka
- **Databases**: 
  - MySQL/PostgreSQL (transactional data)
  - Redis (caching)
  - Apache Pinot (analytics)

### Frontend
- **Web**: Fusion.js (React-based framework)
- **iOS**: Swift with RIBs architecture
- **Android**: Kotlin with RIBs architecture

### Infrastructure
- **Container Orchestration**: Kubernetes
- **CI/CD**: Spinnaker + Jenkins
- **Monitoring**: M3 + Prometheus + Grafana
- **Tracing**: Jaeger
- **Service Discovery**: Consul/Kubernetes DNS

## Microservices

### 1. API Gateway Service
- **Responsibility**: Request routing, authentication, rate limiting
- **Port**: 8080 (HTTP), 9090 (gRPC)
- **Key Features**:
  - JWT authentication
  - Request/response transformation
  - Load balancing
  - Circuit breaking

### 2. Member Service
- **Responsibility**: Member profiles, authentication, dependents
- **Port**: 50051
- **Key Endpoints**:
  - GetMember
  - UpdateMember
  - GetMemberCard
  - ListDependents

### 3. Benefits Service
- **Responsibility**: Coverage information, deductibles, out-of-pocket
- **Port**: 50052
- **Key Endpoints**:
  - GetBenefitsSummary
  - GetBenefitDetails
  - GetDeductibleStatus
  - GetOutOfPocketStatus

### 4. Claims Service
- **Responsibility**: Claims processing, cost estimates
- **Port**: 50054
- **Key Endpoints**:
  - ListClaims
  - GetClaim
  - GetCostEstimate
  - SubmitClaim

### 5. Provider Service
- **Responsibility**: Provider search, network status
- **Port**: 50053
- **Key Endpoints**:
  - SearchProviders
  - GetProvider
  - CheckNetworkStatus

### 6. Messaging Service
- **Responsibility**: Secure member communications
- **Port**: 50055
- **Key Endpoints**:
  - ListConversations
  - GetConversation
  - SendMessage
  - StreamMessages

## Data Architecture

### Primary Database Schema
- **Members**: Core member information
- **Coverage**: Active coverage types
- **Benefits**: Benefit definitions and coverage levels
- **Claims**: Claims history and processing
- **Providers**: Provider network information
- **Messages**: Secure communications

### Event Streaming (Kafka Topics)
- `health.claims`: Claims status updates
- `health.messages`: New message notifications
- `health.audit`: Audit log events
- `health.member.updates`: Member profile changes
- `health.benefit.changes`: Benefit modifications

### Caching Strategy
- **Redis**: Session data, frequently accessed member info
- **CDN**: Static assets, member ID card images
- **Application-level**: gRPC response caching

## Security Architecture

### Authentication & Authorization
- JWT tokens with RS256 signing
- Token expiration: 1 hour (configurable)
- Refresh token rotation
- Biometric authentication on mobile

### Data Protection
- TLS 1.3 for all communications
- AES-256 encryption at rest
- PII data masking in logs
- HIPAA compliance measures

### API Security
- Rate limiting per member
- DDoS protection
- API key management
- Request signing for sensitive operations

## Mobile Architecture (RIBs)

### RIBs Structure
```
Root
├── LoggedOut
│   └── Login
└── LoggedIn
    ├── Dashboard
    ├── Benefits
    ├── Claims
    ├── Providers
    ├── MemberCard
    └── Messages
```

### Key Principles
- Business logic in Interactors
- Routers manage navigation
- Views are passive
- Dependency injection via Builders
- Unidirectional data flow

## Deployment Architecture

### Kubernetes Configuration
- **Namespaces**: dev, staging, production
- **Resource Limits**: CPU and memory constraints
- **Autoscaling**: HPA based on CPU/memory
- **Health Checks**: Liveness and readiness probes

### Service Mesh Features
- Mutual TLS between services
- Traffic management and routing
- Circuit breaking and retries
- Distributed tracing

## Monitoring & Observability

### Metrics (Prometheus + M3)
- Request rate, error rate, duration (RED)
- Business metrics (claims processed, messages sent)
- Infrastructure metrics (CPU, memory, disk)

### Logging
- Structured logging with correlation IDs
- Log aggregation with ELK stack
- Audit trail for compliance

### Tracing (Jaeger)
- End-to-end request tracing
- Performance bottleneck identification
- Service dependency mapping

### Alerting
- Critical: PagerDuty integration
- Warning: Slack notifications
- Custom alert rules per service

## Development Workflow

### Local Development
```bash
# Start all services
docker-compose up

# Run specific service
cd backend/services/member
go run cmd/main.go

# Run web app
cd web
npm run dev
```

### Testing Strategy
- Unit tests: 80% coverage target
- Integration tests: API contract testing
- E2E tests: Critical user journeys
- Performance tests: Load testing with k6

### CI/CD Pipeline
1. Code commit triggers build
2. Run tests and security scans
3. Build Docker images
4. Deploy to staging
5. Run integration tests
6. Manual approval for production
7. Blue-green deployment
8. Smoke tests
9. Monitoring and rollback if needed

## Performance Considerations

### Backend Optimization
- Connection pooling for databases
- gRPC streaming for real-time updates
- Batch processing for claims
- Async message processing

### Frontend Optimization
- Code splitting and lazy loading
- Service worker for offline support
- Image optimization and WebP format
- CDN for static assets

### Mobile Optimization
- Offline-first architecture
- Background sync
- Push notifications
- Biometric caching

## Disaster Recovery

### Backup Strategy
- Database: Daily backups, 30-day retention
- File storage: Continuous replication
- Configuration: Git-based versioning

### RTO/RPO Targets
- RTO: 1 hour
- RPO: 15 minutes
- Automated failover for critical services

## Compliance & Security

### HIPAA Compliance
- Encryption in transit and at rest
- Access controls and audit logging
- Regular security assessments
- Employee training

### Data Privacy
- GDPR compliance for EU members
- Data minimization principles
- Right to deletion implementation
- Privacy by design

## Future Enhancements

### Planned Features
- Telemedicine integration
- Wearable device connectivity
- AI-powered health insights
- Voice assistant integration

### Technical Improvements
- GraphQL federation
- Service mesh migration
- Kubernetes operator for services
- Multi-region deployment