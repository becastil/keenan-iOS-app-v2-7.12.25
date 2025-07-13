#!/bin/bash

# Sydney Health Development Environment Setup Script
# This script sets up the complete development environment for all platforms

set -e

echo "🏥 Sydney Health Development Environment Setup"
echo "============================================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Function to print status
print_status() {
    echo -e "${GREEN}✓${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠${NC} $1"
}

# Check prerequisites
echo -e "\n📋 Checking prerequisites..."

# Check for required tools
MISSING_TOOLS=()
if ! command_exists node; then MISSING_TOOLS+=("node"); fi
if ! command_exists npm; then MISSING_TOOLS+=("npm"); fi
if ! command_exists go; then MISSING_TOOLS+=("go"); fi
if ! command_exists docker; then MISSING_TOOLS+=("docker"); fi
if ! command_exists java; then MISSING_TOOLS+=("java"); fi
if ! command_exists pod; then MISSING_TOOLS+=("cocoapods"); fi

if [ ${#MISSING_TOOLS[@]} -ne 0 ]; then
    print_error "Missing required tools: ${MISSING_TOOLS[*]}"
    echo "Please install the missing tools and run this script again."
    exit 1
fi

print_status "All required tools are installed"

# Install Protocol Buffers compiler
echo -e "\n📦 Installing Protocol Buffers compiler..."
if ! command_exists protoc; then
    print_warning "protoc not found. Installing..."
    
    # Detect OS
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        # Linux installation
        PROTOC_VERSION="25.1"
        wget https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip -O /tmp/protoc.zip
        sudo unzip -o /tmp/protoc.zip -d /usr/local bin/protoc
        sudo unzip -o /tmp/protoc.zip -d /usr/local 'include/*'
        rm -f /tmp/protoc.zip
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        # macOS installation
        brew install protobuf
    else
        print_error "Unsupported OS for automatic protoc installation"
        echo "Please install protoc manually: https://grpc.io/docs/protoc-installation/"
        exit 1
    fi
fi
print_status "protoc installed"

# Install protoc plugins
echo -e "\n📦 Installing protoc plugins..."

# Go plugins
print_status "Installing Go protoc plugins..."
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# TypeScript/JavaScript plugins
print_status "Installing TypeScript protoc plugins..."
npm install -g grpc-tools grpc_tools_node_protoc_ts

# Swift plugin (if on macOS)
if [[ "$OSTYPE" == "darwin"* ]]; then
    print_status "Installing Swift protoc plugins..."
    brew install swift-protobuf
fi

# Java plugin is included with protoc
print_status "Java protoc plugin is included with protoc"

# Generate Protocol Buffer code
echo -e "\n🔧 Generating Protocol Buffer code..."
cd shared/proto
chmod +x build.sh
./build.sh
cd ../..

if [ -d "backend/shared/pb" ] && [ -d "web/src/generated" ]; then
    print_status "Protocol Buffer code generated successfully"
else
    print_error "Protocol Buffer code generation failed"
fi

# Setup Android environment
echo -e "\n🤖 Setting up Android environment..."
cd android

# Add gradle wrapper
if [ ! -f "gradlew" ]; then
    print_status "Adding Gradle wrapper..."
    gradle wrapper --gradle-version=8.3 --distribution-type=all
fi

# Make gradlew executable
chmod +x gradlew

# Create local.properties if it doesn't exist
if [ ! -f "local.properties" ]; then
    echo "sdk.dir=$ANDROID_HOME" > local.properties
    print_status "Created local.properties"
fi

print_status "Android environment configured"
cd ..

# Setup iOS environment
echo -e "\n🍎 Setting up iOS environment..."
cd ios

# Create Info.plist
cat > SydneyHealth/Info.plist << 'EOF'
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>CFBundleDevelopmentRegion</key>
    <string>$(DEVELOPMENT_LANGUAGE)</string>
    <key>CFBundleExecutable</key>
    <string>$(EXECUTABLE_NAME)</string>
    <key>CFBundleIdentifier</key>
    <string>$(PRODUCT_BUNDLE_IDENTIFIER)</string>
    <key>CFBundleInfoDictionaryVersion</key>
    <string>6.0</string>
    <key>CFBundleName</key>
    <string>$(PRODUCT_NAME)</string>
    <key>CFBundlePackageType</key>
    <string>$(PRODUCT_BUNDLE_PACKAGE_TYPE)</string>
    <key>CFBundleShortVersionString</key>
    <string>1.0</string>
    <key>CFBundleVersion</key>
    <string>1</string>
    <key>LSRequiresIPhoneOS</key>
    <true/>
    <key>UIApplicationSceneManifest</key>
    <dict>
        <key>UIApplicationSupportsMultipleScenes</key>
        <false/>
    </dict>
    <key>UILaunchStoryboardName</key>
    <string>LaunchScreen</string>
    <key>UIRequiredDeviceCapabilities</key>
    <array>
        <string>armv7</string>
    </array>
    <key>UISupportedInterfaceOrientations</key>
    <array>
        <string>UIInterfaceOrientationPortrait</string>
        <string>UIInterfaceOrientationLandscapeLeft</string>
        <string>UIInterfaceOrientationLandscapeRight</string>
    </array>
    <key>NSFaceIDUsageDescription</key>
    <string>Sydney Health uses Face ID to secure your health information</string>
    <key>NSCameraUsageDescription</key>
    <string>Sydney Health needs camera access to scan insurance cards and documents</string>
    <key>NSPhotoLibraryUsageDescription</key>
    <string>Sydney Health needs photo library access to upload claim receipts</string>
</dict>
</plist>
EOF
print_status "Created Info.plist"

# Install CocoaPods dependencies
if command_exists pod; then
    print_status "Installing CocoaPods dependencies..."
    pod install
else
    print_warning "CocoaPods not installed. Run 'pod install' manually later."
fi

cd ..

# Setup backend environment
echo -e "\n🔧 Setting up backend environment..."
cd backend

# Download Go dependencies
print_status "Downloading Go dependencies..."
go mod download

# Create .env file for local development
cat > .env << 'EOF'
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=sydney_health
DB_PASSWORD=dev_password
DB_NAME=sydney_health

# Kafka
KAFKA_BROKERS=localhost:9092

# JWT
JWT_SECRET=dev_secret_key_change_in_production

# Service Ports
GATEWAY_PORT=8080
MEMBER_SERVICE_PORT=50051
BENEFITS_SERVICE_PORT=50052
CLAIMS_SERVICE_PORT=50053
PROVIDER_SERVICE_PORT=50054
MESSAGING_SERVICE_PORT=50055

# Metrics
METRICS_PORT=9091
EOF
print_status "Created backend .env file"

cd ..

# Setup web environment
echo -e "\n🌐 Setting up web environment..."
cd web

# Install dependencies
print_status "Installing web dependencies..."
npm install

# Create .env file
cat > .env << 'EOF'
# API Configuration
REACT_APP_API_URL=http://localhost:8080
REACT_APP_WS_URL=ws://localhost:8080

# Feature Flags
REACT_APP_ENABLE_MOCK_DATA=true
REACT_APP_ENABLE_ANALYTICS=false
EOF
print_status "Created web .env file"

cd ..

# Setup local development infrastructure
echo -e "\n🐳 Setting up local development infrastructure..."

# Create docker-compose for local development
cat > docker-compose.dev.yml << 'EOF'
version: '3.8'

services:
  postgres:
    image: postgres:15
    environment:
      POSTGRES_USER: sydney_health
      POSTGRES_PASSWORD: dev_password
      POSTGRES_DB: sydney_health
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./backend/migrations:/docker-entrypoint-initdb.d

  zookeeper:
    image: confluentinc/cp-zookeeper:7.5.0
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      ZOOKEEPER_TICK_TIME: 2000

  kafka:
    image: confluentinc/cp-kafka:7.5.0
    depends_on:
      - zookeeper
    ports:
      - "9092:9092"
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://kafka:29092,PLAINTEXT_HOST://localhost:9092
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
      KAFKA_TRANSACTION_STATE_LOG_MIN_ISR: 1
      KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR: 1

  kafka-ui:
    image: provectuslabs/kafka-ui:latest
    depends_on:
      - kafka
    ports:
      - "8090:8080"
    environment:
      KAFKA_CLUSTERS_0_NAME: local
      KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS: kafka:29092

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"

volumes:
  postgres_data:
EOF
print_status "Created docker-compose.dev.yml"

# Create development scripts
echo -e "\n📝 Creating development scripts..."

# Create a main development script
cat > dev.sh << 'EOF'
#!/bin/bash

# Sydney Health Development Helper Script

case "$1" in
    start)
        echo "Starting development infrastructure..."
        docker-compose -f docker-compose.dev.yml up -d
        echo "Infrastructure started. Services available at:"
        echo "  - PostgreSQL: localhost:5432"
        echo "  - Kafka: localhost:9092"
        echo "  - Kafka UI: http://localhost:8090"
        echo "  - Redis: localhost:6379"
        ;;
    stop)
        echo "Stopping development infrastructure..."
        docker-compose -f docker-compose.dev.yml down
        ;;
    reset)
        echo "Resetting development infrastructure..."
        docker-compose -f docker-compose.dev.yml down -v
        docker-compose -f docker-compose.dev.yml up -d
        ;;
    backend)
        echo "Starting backend services..."
        cd backend/cmd/gateway && go run main.go &
        cd backend/cmd/member && go run main.go &
        echo "Backend services started"
        ;;
    web)
        echo "Starting web application..."
        cd web && npm run dev
        ;;
    android)
        echo "Building Android app..."
        cd android && ./gradlew assembleDebug
        ;;
    ios)
        echo "Opening iOS project..."
        cd ios && open SydneyHealth.xcworkspace
        ;;
    proto)
        echo "Regenerating Protocol Buffer code..."
        cd shared/proto && ./build.sh
        ;;
    test)
        echo "Running all tests..."
        cd backend && go test ./...
        cd ../web && npm test
        cd ../android && ./gradlew test
        ;;
    *)
        echo "Usage: ./dev.sh {start|stop|reset|backend|web|android|ios|proto|test}"
        exit 1
        ;;
esac
EOF

chmod +x dev.sh
print_status "Created dev.sh helper script"

# Create README for development
cat > DEVELOPMENT_SETUP.md << 'EOF'
# Sydney Health Development Setup

This document describes the development environment setup for the Sydney Health mobile application.

## Prerequisites

- Node.js (v18+)
- Go (v1.21+)
- Docker & Docker Compose
- Java 11+ (for Android)
- Xcode (for iOS, macOS only)
- Android Studio or Android SDK

## Quick Start

1. Run the setup script:
   ```bash
   ./setup-dev-environment.sh
   ```

2. Start the development infrastructure:
   ```bash
   ./dev.sh start
   ```

3. Start backend services:
   ```bash
   ./dev.sh backend
   ```

4. Start the web application:
   ```bash
   ./dev.sh web
   ```

## Development Commands

- `./dev.sh start` - Start Docker infrastructure (DB, Kafka, Redis)
- `./dev.sh stop` - Stop Docker infrastructure
- `./dev.sh reset` - Reset infrastructure (clean state)
- `./dev.sh backend` - Run backend services
- `./dev.sh web` - Run web application
- `./dev.sh android` - Build Android app
- `./dev.sh ios` - Open iOS project in Xcode
- `./dev.sh proto` - Regenerate Protocol Buffer code
- `./dev.sh test` - Run all tests

## Service URLs

- Web App: http://localhost:3000
- API Gateway: http://localhost:8080
- Kafka UI: http://localhost:8090
- Metrics: http://localhost:9091/metrics

## Default Credentials

- Demo User: M123456 / demo
- Database: sydney_health / dev_password

## Architecture Overview

```
┌─────────────┐ ┌─────────────┐ ┌─────────────┐
│   iOS App   │ │ Android App │ │   Web App   │
└──────┬──────┘ └──────┬──────┘ └──────┬──────┘
       │               │               │
       └───────────────┴───────────────┘
                       │
                ┌──────▼──────┐
                │ API Gateway │
                └──────┬──────┘
                       │
       ┌───────────────┴───────────────┐
       │                               │
┌──────▼──────┐  ┌────────────┐  ┌────▼────┐
│   Member    │  │  Benefits  │  │ Claims  │
│   Service   │  │  Service   │  │ Service │
└─────────────┘  └────────────┘  └─────────┘
       │               │               │
       └───────────────┴───────────────┘
                       │
                ┌──────▼──────┐
                │  PostgreSQL │
                └─────────────┘
```

## Troubleshooting

### Port Conflicts
If you get port conflicts, check what's using the ports:
```bash
lsof -i :8080  # API Gateway
lsof -i :5432  # PostgreSQL
lsof -i :9092  # Kafka
```

### Database Issues
Reset the database:
```bash
./dev.sh reset
```

### Protocol Buffer Issues
Regenerate protobuf code:
```bash
./dev.sh proto
```

## Next Steps

1. Complete the missing service implementations
2. Add comprehensive test coverage
3. Implement proper authentication
4. Connect services to the database
5. Build out the mobile app features

For more information, see the main README.md
EOF
print_status "Created DEVELOPMENT_SETUP.md"

echo -e "\n✅ Development environment setup complete!"
echo -e "\n📖 Next steps:"
echo "1. Start the infrastructure: ./dev.sh start"
echo "2. Run backend services: ./dev.sh backend"
echo "3. Start web app: ./dev.sh web"
echo "4. See DEVELOPMENT_SETUP.md for more details"