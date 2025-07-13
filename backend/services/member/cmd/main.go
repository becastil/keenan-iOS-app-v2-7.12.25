package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/sydney-health-clone/backend/services/member/internal/service"
	"github.com/sydney-health-clone/backend/shared/config"
	"github.com/sydney-health-clone/backend/shared/logger"
	pb "github.com/sydney-health-clone/backend/shared/pb"
	
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	configPath = flag.String("config", "config/member.yaml", "Path to configuration file")
	port       = flag.Int("port", 50051, "gRPC server port")
)

func main() {
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		// Use default config for demo
		cfg = &config.Config{
			Server: config.ServerConfig{
				LogLevel: "debug",
			},
		}
	}

	if err := logger.Init(cfg.Server.LogLevel); err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	logger.Info("Starting Member Service", zap.Int("port", *port))

	// Create gRPC server
	grpcServer := grpc.NewServer()
	
	// Initialize service with mock data
	memberService := service.NewMemberService()
	pb.RegisterMemberServiceServer(grpcServer, memberService)

	// Start listening
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logger.Fatal("Failed to listen", zap.Error(err))
	}

	// Start server in goroutine
	go func() {
		logger.Info("Member Service listening", zap.String("addr", lis.Addr().String()))
		if err := grpcServer.Serve(lis); err != nil {
			logger.Fatal("Failed to serve", zap.Error(err))
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	
	logger.Info("Shutting down Member Service...")
	grpcServer.GracefulStop()
	logger.Info("Member Service exited")
}