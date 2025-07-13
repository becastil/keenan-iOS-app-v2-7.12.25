package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sydney-health-clone/backend/services/gateway/internal/handler"
	"github.com/sydney-health-clone/backend/services/gateway/internal/proxy"
	"github.com/sydney-health-clone/backend/shared/config"
	"github.com/sydney-health-clone/backend/shared/logger"
	
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

var (
	configPath = flag.String("config", "config/gateway.yaml", "Path to configuration file")
)

func main() {
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	if err := logger.Init(cfg.Server.LogLevel); err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	logger.Info("Starting API Gateway",
		zap.String("environment", cfg.Server.Environment),
		zap.Int("port", cfg.Server.Port),
	)

	// Initialize service proxy
	serviceProxy, err := proxy.NewServiceProxy(cfg)
	if err != nil {
		logger.Fatal("Failed to create service proxy", zap.Error(err))
	}

	// Setup routes
	router := setupRoutes(serviceProxy, cfg)

	// Setup CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		ExposedHeaders:   []string{"X-Total-Count"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// Setup HTTP server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      corsHandler.Handler(router),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start metrics server if enabled
	if cfg.Metrics.Enabled {
		go startMetricsServer(cfg.Metrics)
	}

	// Start server in goroutine
	go func() {
		logger.Info("API Gateway listening", zap.String("addr", srv.Addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	logger.Info("Server exited")
}

func setupRoutes(proxy *proxy.ServiceProxy, cfg *config.Config) *mux.Router {
	r := mux.NewRouter()
	
	// Health check
	r.HandleFunc("/health", handler.HealthCheck).Methods("GET")
	
	// API routes
	api := r.PathPrefix("/api/v1").Subrouter()
	
	// Member routes
	api.HandleFunc("/members/{memberId}", proxy.GetMember).Methods("GET")
	api.HandleFunc("/members/{memberId}", proxy.UpdateMember).Methods("PUT")
	api.HandleFunc("/members/{memberId}/card", proxy.GetMemberCard).Methods("GET")
	api.HandleFunc("/members/{memberId}/dependents", proxy.ListDependents).Methods("GET")
	
	// Benefits routes
	api.HandleFunc("/members/{memberId}/benefits", proxy.GetBenefitsSummary).Methods("GET")
	api.HandleFunc("/members/{memberId}/benefits/{benefitId}", proxy.GetBenefitDetails).Methods("GET")
	api.HandleFunc("/members/{memberId}/deductible", proxy.GetDeductibleStatus).Methods("GET")
	api.HandleFunc("/members/{memberId}/out-of-pocket", proxy.GetOutOfPocketStatus).Methods("GET")
	
	// Provider routes
	api.HandleFunc("/providers/search", proxy.SearchProviders).Methods("GET")
	api.HandleFunc("/providers/{providerId}", proxy.GetProvider).Methods("GET")
	api.HandleFunc("/providers/{providerId}/network-status", proxy.CheckNetworkStatus).Methods("GET")
	
	// Claims routes
	api.HandleFunc("/members/{memberId}/claims", proxy.ListClaims).Methods("GET")
	api.HandleFunc("/claims/{claimId}", proxy.GetClaim).Methods("GET")
	api.HandleFunc("/members/{memberId}/cost-estimate", proxy.GetCostEstimate).Methods("POST")
	api.HandleFunc("/members/{memberId}/claims", proxy.SubmitClaim).Methods("POST")
	
	// Messaging routes
	api.HandleFunc("/members/{memberId}/conversations", proxy.ListConversations).Methods("GET")
	api.HandleFunc("/conversations/{conversationId}", proxy.GetConversation).Methods("GET")
	api.HandleFunc("/conversations/{conversationId}/messages", proxy.SendMessage).Methods("POST")
	api.HandleFunc("/messages/mark-read", proxy.MarkAsRead).Methods("POST")
	
	// Apply auth middleware to all API routes
	api.Use(handler.AuthMiddleware(cfg.Auth))
	
	return r
}

func startMetricsServer(cfg config.MetricsConfig) {
	mux := http.NewServeMux()
	mux.Handle(cfg.Path, promhttp.Handler())
	
	addr := fmt.Sprintf(":%d", cfg.Port)
	logger.Info("Starting metrics server", zap.String("addr", addr))
	
	if err := http.ListenAndServe(addr, mux); err != nil {
		logger.Error("Metrics server failed", zap.Error(err))
	}
}