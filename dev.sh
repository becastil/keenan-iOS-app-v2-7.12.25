#!/bin/bash

# Sydney Health Development Helper Script
# This script provides convenient commands for development tasks

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_info() {
    echo -e "${BLUE}ℹ${NC} $1"
}

print_success() {
    echo -e "${GREEN}✓${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠${NC} $1"
}

# Main command handler
case "$1" in
    start)
        print_info "Starting development infrastructure..."
        docker-compose -f docker-compose.dev.yml up -d
        
        # Wait for services to be healthy
        print_info "Waiting for services to be ready..."
        sleep 10
        
        print_success "Infrastructure started successfully!"
        echo ""
        echo "Services available at:"
        echo "  📊 PostgreSQL: localhost:5432 (user: sydney_health, pass: dev_password)"
        echo "  📨 Kafka: localhost:9092"
        echo "  🎯 Kafka UI: http://localhost:8090"
        echo "  💾 Redis: localhost:6379"
        echo "  🔍 Jaeger UI: http://localhost:16686"
        echo "  📈 Prometheus: http://localhost:9090"
        echo "  📊 Grafana: http://localhost:3001 (admin/admin)"
        ;;
        
    stop)
        print_info "Stopping development infrastructure..."
        docker-compose -f docker-compose.dev.yml down
        print_success "Infrastructure stopped"
        ;;
        
    reset)
        print_warning "Resetting development infrastructure (this will delete all data)..."
        docker-compose -f docker-compose.dev.yml down -v
        docker-compose -f docker-compose.dev.yml up -d
        print_success "Infrastructure reset complete"
        ;;
        
    status)
        print_info "Checking infrastructure status..."
        docker-compose -f docker-compose.dev.yml ps
        ;;
        
    logs)
        service=${2:-}
        if [ -z "$service" ]; then
            docker-compose -f docker-compose.dev.yml logs -f
        else
            docker-compose -f docker-compose.dev.yml logs -f "$service"
        fi
        ;;
        
    backend)
        print_info "Starting backend services..."
        
        # Load environment variables
        export $(cat backend/.env.development | grep -v '^#' | xargs)
        
        # Start services in background
        print_info "Starting API Gateway..."
        (cd backend/cmd/gateway && go run main.go) &
        GATEWAY_PID=$!
        
        print_info "Starting Member Service..."
        (cd backend/cmd/member && go run main.go) &
        MEMBER_PID=$!
        
        print_success "Backend services started"
        echo ""
        echo "Services running:"
        echo "  🌐 API Gateway: http://localhost:8080"
        echo "  👤 Member Service: localhost:50051 (gRPC)"
        echo ""
        echo "Press Ctrl+C to stop all services"
        
        # Wait for interrupt
        trap "kill $GATEWAY_PID $MEMBER_PID 2>/dev/null" INT
        wait
        ;;
        
    web)
        print_info "Starting web application..."
        cd web
        npm run dev
        ;;
        
    android)
        print_info "Building Android app..."
        cd android
        
        # Download gradle wrapper if needed
        if [ ! -f "gradle/wrapper/gradle-wrapper.jar" ]; then
            print_warning "Gradle wrapper not found. Downloading..."
            ./download-gradle-wrapper.sh
        fi
        
        ./gradlew assembleDebug
        
        APK_PATH="app/build/outputs/apk/debug/app-debug.apk"
        if [ -f "$APK_PATH" ]; then
            print_success "Android build successful!"
            echo "APK location: android/$APK_PATH"
        else
            print_error "Android build failed"
            exit 1
        fi
        ;;
        
    ios)
        print_info "Opening iOS project..."
        cd ios
        
        if [ ! -d "Pods" ]; then
            print_warning "CocoaPods not installed. Installing..."
            pod install
        fi
        
        open SydneyHealth.xcworkspace
        print_success "iOS project opened in Xcode"
        ;;
        
    proto)
        print_info "Regenerating Protocol Buffer code..."
        cd shared/proto
        ./build.sh
        cd ../..
        print_success "Protocol Buffer code regenerated"
        ;;
        
    test)
        print_info "Running all tests..."
        
        # Backend tests
        print_info "Running backend tests..."
        cd backend
        go test -v ./...
        cd ..
        
        # Web tests
        print_info "Running web tests..."
        cd web
        npm test -- --watchAll=false
        cd ..
        
        # Android tests
        print_info "Running Android tests..."
        cd android
        ./gradlew test
        cd ..
        
        print_success "All tests completed"
        ;;
        
    lint)
        print_info "Running linters..."
        
        # Backend
        print_info "Linting Go code..."
        cd backend
        golangci-lint run || print_warning "Go linting issues found"
        cd ..
        
        # Web
        print_info "Linting JavaScript/TypeScript..."
        cd web
        npm run lint || print_warning "JS/TS linting issues found"
        cd ..
        
        # Android
        print_info "Linting Kotlin code..."
        cd android
        ./gradlew lint || print_warning "Android linting issues found"
        cd ..
        
        print_success "Linting completed"
        ;;
        
    setup)
        print_info "Running initial setup..."
        ./setup-dev-environment.sh
        ;;
        
    help|*)
        echo "🏥 Sydney Health Development Helper"
        echo ""
        echo "Usage: ./dev.sh <command> [options]"
        echo ""
        echo "Infrastructure Commands:"
        echo "  start      Start Docker infrastructure (DB, Kafka, Redis, etc.)"
        echo "  stop       Stop Docker infrastructure"
        echo "  reset      Reset infrastructure (deletes all data)"
        echo "  status     Show infrastructure status"
        echo "  logs       Show logs (optionally specify service)"
        echo ""
        echo "Development Commands:"
        echo "  backend    Run backend services"
        echo "  web        Run web application"
        echo "  android    Build Android app"
        echo "  ios        Open iOS project in Xcode"
        echo ""
        echo "Build Commands:"
        echo "  proto      Regenerate Protocol Buffer code"
        echo "  test       Run all tests"
        echo "  lint       Run all linters"
        echo ""
        echo "Setup Commands:"
        echo "  setup      Run initial environment setup"
        echo "  help       Show this help message"
        echo ""
        echo "Examples:"
        echo "  ./dev.sh start           # Start infrastructure"
        echo "  ./dev.sh backend         # Run backend services"
        echo "  ./dev.sh logs kafka      # Show Kafka logs"
        echo "  ./dev.sh test            # Run all tests"
        ;;
esac