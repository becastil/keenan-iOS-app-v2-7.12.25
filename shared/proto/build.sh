#!/bin/bash

# Build script for generating gRPC code from proto files

set -e

PROTO_DIR=$(dirname "$0")
ROOT_DIR=$(cd "$PROTO_DIR/../.." && pwd)

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${YELLOW}Building protocol buffer files...${NC}"

# Check if protoc is installed
if ! command -v protoc &> /dev/null; then
    echo -e "${RED}protoc is not installed. Please install protocol buffers compiler.${NC}"
    exit 1
fi

# Check if protoc-gen-go is installed
if ! command -v protoc-gen-go &> /dev/null; then
    echo -e "${RED}protoc-gen-go is not installed. Run: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest${NC}"
    exit 1
fi

# Check if protoc-gen-go-grpc is installed
if ! command -v protoc-gen-go-grpc &> /dev/null; then
    echo -e "${RED}protoc-gen-go-grpc is not installed. Run: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest${NC}"
    exit 1
fi

# Create output directories
mkdir -p "$ROOT_DIR/backend/shared/pb"
mkdir -p "$ROOT_DIR/web/src/generated"

# Generate Go code
echo -e "${GREEN}Generating Go code...${NC}"
protoc \
    --proto_path="$PROTO_DIR" \
    --go_out="$ROOT_DIR/backend/shared/pb" \
    --go_opt=paths=source_relative \
    --go-grpc_out="$ROOT_DIR/backend/shared/pb" \
    --go-grpc_opt=paths=source_relative \
    "$PROTO_DIR"/*.proto

# Generate JavaScript/TypeScript code for web
if command -v protoc-gen-ts &> /dev/null; then
    echo -e "${GREEN}Generating TypeScript code...${NC}"
    protoc \
        --proto_path="$PROTO_DIR" \
        --plugin=protoc-gen-ts="$(which protoc-gen-ts)" \
        --ts_out="$ROOT_DIR/web/src/generated" \
        "$PROTO_DIR"/*.proto
else
    echo -e "${YELLOW}protoc-gen-ts not installed. Skipping TypeScript generation.${NC}"
fi

# Generate Swift code for iOS
if command -v protoc-gen-swift &> /dev/null && command -v protoc-gen-grpc-swift &> /dev/null; then
    echo -e "${GREEN}Generating Swift code...${NC}"
    mkdir -p "$ROOT_DIR/ios/Generated"
    protoc \
        --proto_path="$PROTO_DIR" \
        --swift_out="$ROOT_DIR/ios/Generated" \
        --grpc-swift_out="$ROOT_DIR/ios/Generated" \
        "$PROTO_DIR"/*.proto
else
    echo -e "${YELLOW}Swift protoc plugins not installed. Skipping Swift generation.${NC}"
fi

# Generate Java/Kotlin code for Android
if command -v protoc-gen-grpc-java &> /dev/null; then
    echo -e "${GREEN}Generating Java code...${NC}"
    mkdir -p "$ROOT_DIR/android/app/src/main/java/generated"
    protoc \
        --proto_path="$PROTO_DIR" \
        --java_out="$ROOT_DIR/android/app/src/main/java/generated" \
        --grpc-java_out="$ROOT_DIR/android/app/src/main/java/generated" \
        "$PROTO_DIR"/*.proto
else
    echo -e "${YELLOW}Java gRPC plugin not installed. Skipping Java generation.${NC}"
fi

echo -e "${GREEN}Protocol buffer build complete!${NC}"