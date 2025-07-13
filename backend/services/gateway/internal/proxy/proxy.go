package proxy

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sydney-health-clone/backend/shared/config"
	"github.com/sydney-health-clone/backend/shared/logger"
	pb "github.com/sydney-health-clone/backend/shared/pb"
	
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceProxy struct {
	cfg              *config.Config
	memberClient     pb.MemberServiceClient
	benefitsClient   pb.BenefitsServiceClient
	providerClient   pb.ProviderServiceClient
	claimsClient     pb.ClaimsServiceClient
	messagingClient  pb.MessagingServiceClient
}

func NewServiceProxy(cfg *config.Config) (*ServiceProxy, error) {
	proxy := &ServiceProxy{cfg: cfg}
	
	// Connect to Member Service
	memberConn, err := createConnection(cfg.Services.MemberService)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to member service: %w", err)
	}
	proxy.memberClient = pb.NewMemberServiceClient(memberConn)
	
	// Connect to Benefits Service
	benefitsConn, err := createConnection(cfg.Services.BenefitsService)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to benefits service: %w", err)
	}
	proxy.benefitsClient = pb.NewBenefitsServiceClient(benefitsConn)
	
	// Connect to Provider Service
	providerConn, err := createConnection(cfg.Services.ProviderService)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to provider service: %w", err)
	}
	proxy.providerClient = pb.NewProviderServiceClient(providerConn)
	
	// Connect to Claims Service
	claimsConn, err := createConnection(cfg.Services.ClaimsService)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to claims service: %w", err)
	}
	proxy.claimsClient = pb.NewClaimsServiceClient(claimsConn)
	
	// Connect to Messaging Service
	messagingConn, err := createConnection(cfg.Services.MessagingService)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to messaging service: %w", err)
	}
	proxy.messagingClient = pb.NewMessagingServiceClient(messagingConn)
	
	return proxy, nil
}

func createConnection(endpoint config.ServiceEndpoint) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	target := fmt.Sprintf("%s:%d", endpoint.Host, endpoint.Port)
	
	conn, err := grpc.DialContext(ctx, target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	
	if err != nil {
		return nil, err
	}
	
	return conn, nil
}

// Member Service Handlers

func (p *ServiceProxy) GetMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberID := vars["memberId"]
	
	ctx := r.Context()
	resp, err := p.memberClient.GetMember(ctx, &pb.GetMemberRequest{
		MemberId: memberID,
	})
	
	if err != nil {
		handleError(w, err)
		return
	}
	
	respondJSON(w, http.StatusOK, resp.Member)
}

func (p *ServiceProxy) UpdateMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberID := vars["memberId"]
	
	var member pb.Member
	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	
	member.MemberId = memberID
	
	ctx := r.Context()
	resp, err := p.memberClient.UpdateMember(ctx, &pb.UpdateMemberRequest{
		Member: &member,
	})
	
	if err != nil {
		handleError(w, err)
		return
	}
	
	respondJSON(w, http.StatusOK, resp.Member)
}

func (p *ServiceProxy) GetMemberCard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberID := vars["memberId"]
	coverageType := r.URL.Query().Get("coverage_type")
	
	ctx := r.Context()
	resp, err := p.memberClient.GetMemberCard(ctx, &pb.GetMemberCardRequest{
		MemberId:     memberID,
		CoverageType: parseCoverageType(coverageType),
	})
	
	if err != nil {
		handleError(w, err)
		return
	}
	
	respondJSON(w, http.StatusOK, resp.Card)
}

func (p *ServiceProxy) ListDependents(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberID := vars["memberId"]
	
	ctx := r.Context()
	resp, err := p.memberClient.ListDependents(ctx, &pb.ListDependentsRequest{
		MemberId: memberID,
	})
	
	if err != nil {
		handleError(w, err)
		return
	}
	
	respondJSON(w, http.StatusOK, resp.Dependents)
}

// Benefits Service Handlers

func (p *ServiceProxy) GetBenefitsSummary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberID := vars["memberId"]
	coverageType := r.URL.Query().Get("coverage_type")
	
	ctx := r.Context()
	resp, err := p.benefitsClient.GetBenefitsSummary(ctx, &pb.GetBenefitsSummaryRequest{
		MemberId:     memberID,
		CoverageType: parseCoverageType(coverageType),
	})
	
	if err != nil {
		handleError(w, err)
		return
	}
	
	respondJSON(w, http.StatusOK, resp.Benefits)
}

func (p *ServiceProxy) GetBenefitDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberID := vars["memberId"]
	benefitID := vars["benefitId"]
	
	ctx := r.Context()
	resp, err := p.benefitsClient.GetBenefitDetails(ctx, &pb.GetBenefitDetailsRequest{
		MemberId:  memberID,
		BenefitId: benefitID,
	})
	
	if err != nil {
		handleError(w, err)
		return
	}
	
	respondJSON(w, http.StatusOK, resp.Benefit)
}

func (p *ServiceProxy) GetDeductibleStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberID := vars["memberId"]
	coverageType := r.URL.Query().Get("coverage_type")
	
	ctx := r.Context()
	resp, err := p.benefitsClient.GetDeductibleStatus(ctx, &pb.GetDeductibleStatusRequest{
		MemberId:     memberID,
		CoverageType: parseCoverageType(coverageType),
	})
	
	if err != nil {
		handleError(w, err)
		return
	}
	
	respondJSON(w, http.StatusOK, resp.Status)
}

func (p *ServiceProxy) GetOutOfPocketStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberID := vars["memberId"]
	coverageType := r.URL.Query().Get("coverage_type")
	
	ctx := r.Context()
	resp, err := p.benefitsClient.GetOutOfPocketStatus(ctx, &pb.GetOutOfPocketStatusRequest{
		MemberId:     memberID,
		CoverageType: parseCoverageType(coverageType),
	})
	
	if err != nil {
		handleError(w, err)
		return
	}
	
	respondJSON(w, http.StatusOK, resp.Status)
}

// Provider Service Handlers

func (p *ServiceProxy) SearchProviders(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement provider search with query parameters
	respondJSON(w, http.StatusOK, map[string]interface{}{
		"providers": []interface{}{},
		"page": map[string]interface{}{
			"next_page_token": "",
			"total_count":     0,
		},
	})
}

func (p *ServiceProxy) GetProvider(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	providerID := vars["providerId"]
	
	ctx := r.Context()
	resp, err := p.providerClient.GetProvider(ctx, &pb.GetProviderRequest{
		ProviderId: providerID,
	})
	
	if err != nil {
		handleError(w, err)
		return
	}
	
	respondJSON(w, http.StatusOK, resp.Provider)
}

func (p *ServiceProxy) CheckNetworkStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	providerID := vars["providerId"]
	memberID := r.URL.Query().Get("member_id")
	coverageType := r.URL.Query().Get("coverage_type")
	
	ctx := r.Context()
	resp, err := p.providerClient.CheckNetworkStatus(ctx, &pb.CheckNetworkStatusRequest{
		ProviderId:   providerID,
		MemberId:     memberID,
		CoverageType: parseCoverageType(coverageType),
	})
	
	if err != nil {
		handleError(w, err)
		return
	}
	
	respondJSON(w, http.StatusOK, resp)
}

// Claims Service Handlers

func (p *ServiceProxy) ListClaims(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberID := vars["memberId"]
	
	// TODO: Parse query parameters for filtering
	
	ctx := r.Context()
	resp, err := p.claimsClient.ListClaims(ctx, &pb.ListClaimsRequest{
		MemberId: memberID,
	})
	
	if err != nil {
		handleError(w, err)
		return
	}
	
	respondJSON(w, http.StatusOK, resp)
}

func (p *ServiceProxy) GetClaim(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	claimID := vars["claimId"]
	
	ctx := r.Context()
	resp, err := p.claimsClient.GetClaim(ctx, &pb.GetClaimRequest{
		ClaimId: claimID,
	})
	
	if err != nil {
		handleError(w, err)
		return
	}
	
	respondJSON(w, http.StatusOK, resp.Claim)
}

func (p *ServiceProxy) GetCostEstimate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberID := vars["memberId"]
	
	var req pb.GetCostEstimateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	
	req.MemberId = memberID
	
	ctx := r.Context()
	resp, err := p.claimsClient.GetCostEstimate(ctx, &req)
	
	if err != nil {
		handleError(w, err)
		return
	}
	
	respondJSON(w, http.StatusOK, resp.Estimate)
}

func (p *ServiceProxy) SubmitClaim(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberID := vars["memberId"]
	
	// TODO: Handle multipart form data for receipt image
	
	ctx := r.Context()
	resp, err := p.claimsClient.SubmitClaim(ctx, &pb.SubmitClaimRequest{
		MemberId: memberID,
		// Fill in other fields from request
	})
	
	if err != nil {
		handleError(w, err)
		return
	}
	
	respondJSON(w, http.StatusCreated, resp)
}

// Messaging Service Handlers

func (p *ServiceProxy) ListConversations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberID := vars["memberId"]
	
	ctx := r.Context()
	resp, err := p.messagingClient.ListConversations(ctx, &pb.ListConversationsRequest{
		MemberId: memberID,
	})
	
	if err != nil {
		handleError(w, err)
		return
	}
	
	respondJSON(w, http.StatusOK, resp)
}

func (p *ServiceProxy) GetConversation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	conversationID := vars["conversationId"]
	
	ctx := r.Context()
	resp, err := p.messagingClient.GetConversation(ctx, &pb.GetConversationRequest{
		ConversationId: conversationID,
	})
	
	if err != nil {
		handleError(w, err)
		return
	}
	
	respondJSON(w, http.StatusOK, resp)
}

func (p *ServiceProxy) SendMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	conversationID := vars["conversationId"]
	
	var req struct {
		MemberID string `json:"member_id"`
		Content  string `json:"content"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	
	ctx := r.Context()
	resp, err := p.messagingClient.SendMessage(ctx, &pb.SendMessageRequest{
		ConversationId: conversationID,
		MemberId:       req.MemberID,
		Content:        req.Content,
	})
	
	if err != nil {
		handleError(w, err)
		return
	}
	
	respondJSON(w, http.StatusCreated, resp.Message)
}

func (p *ServiceProxy) MarkAsRead(w http.ResponseWriter, r *http.Request) {
	var req struct {
		MemberID   string   `json:"member_id"`
		MessageIDs []string `json:"message_ids"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	
	ctx := r.Context()
	resp, err := p.messagingClient.MarkAsRead(ctx, &pb.MarkAsReadRequest{
		MemberId:    req.MemberID,
		MessageIds:  req.MessageIDs,
	})
	
	if err != nil {
		handleError(w, err)
		return
	}
	
	respondJSON(w, http.StatusOK, resp)
}

// Helper functions

func handleError(w http.ResponseWriter, err error) {
	logger.Error("Request failed", zap.Error(err))
	// TODO: Map gRPC errors to HTTP status codes
	respondError(w, http.StatusInternalServerError, "Internal server error")
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Error("Failed to encode response", zap.Error(err))
	}
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{
		"error": message,
	})
}

func parseCoverageType(s string) pb.CoverageType {
	switch s {
	case "medical":
		return pb.CoverageType_COVERAGE_TYPE_MEDICAL
	case "dental":
		return pb.CoverageType_COVERAGE_TYPE_DENTAL
	case "vision":
		return pb.CoverageType_COVERAGE_TYPE_VISION
	case "pharmacy":
		return pb.CoverageType_COVERAGE_TYPE_PHARMACY
	default:
		return pb.CoverageType_COVERAGE_TYPE_UNSPECIFIED
	}
}