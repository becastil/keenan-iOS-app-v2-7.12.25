package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/joho/godotenv"
	"github.com/sydney-health/backend/internal/database"
	"github.com/sydney-health/backend/services/member/repository"
	"github.com/sydney-health/backend/services/member/service"
	pb "github.com/sydney-health/backend/shared/pb"
)

func main() {
	// Load environment variables
	if err := godotenv.Load("../../.env.development"); err != nil {
		log.Printf("Warning: .env.development file not found")
	}

	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize repository
	memberRepo := repository.NewMemberRepository(db)

	// Initialize service
	memberService := service.NewMemberService(memberRepo)

	// Get port from environment
	port := os.Getenv("MEMBER_SERVICE_PORT")
	if port == "" {
		port = "50051"
	}

	// Create gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Register member service
	pb.RegisterMemberServiceServer(grpcServer, memberService)

	// Register health service
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	// Register reflection service for debugging
	reflection.Register(grpcServer)

	// Start server in a goroutine
	go func() {
		log.Printf("Member service listening on port %s", port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Setup graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down member service...")
	grpcServer.GracefulStop()
	log.Println("Member service stopped")
}