package service

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sydney-health/backend/services/member/repository"
	pb "github.com/sydney-health/backend/shared/pb"
)

// MemberService implements the gRPC MemberService
type MemberService struct {
	pb.UnimplementedMemberServiceServer
	repo *repository.MemberRepository
}

// NewMemberService creates a new member service
func NewMemberService(repo *repository.MemberRepository) *MemberService {
	return &MemberService{
		repo: repo,
	}
}

// GetMember retrieves member information
func (s *MemberService) GetMember(ctx context.Context, req *pb.GetMemberRequest) (*pb.GetMemberResponse, error) {
	log.Printf("GetMember called for member ID: %s", req.MemberId)

	if req.MemberId == "" {
		return nil, status.Error(codes.InvalidArgument, "member_id is required")
	}

	member, err := s.repo.GetMember(ctx, req.MemberId)
	if err != nil {
		log.Printf("Error retrieving member: %v", err)
		return nil, status.Error(codes.NotFound, "member not found")
	}

	return &pb.GetMemberResponse{
		Member: member,
	}, nil
}

// UpdateMember updates member information
func (s *MemberService) UpdateMember(ctx context.Context, req *pb.UpdateMemberRequest) (*pb.UpdateMemberResponse, error) {
	log.Printf("UpdateMember called for member ID: %s", req.Member.MemberId)

	if req.Member == nil || req.Member.MemberId == "" {
		return nil, status.Error(codes.InvalidArgument, "member and member_id are required")
	}

	// Validate update fields
	if req.UpdateMask != nil && len(req.UpdateMask.Paths) == 0 {
		return nil, status.Error(codes.InvalidArgument, "update_mask cannot be empty")
	}

	err := s.repo.UpdateMember(ctx, req.Member)
	if err != nil {
		log.Printf("Error updating member: %v", err)
		return nil, status.Error(codes.Internal, "failed to update member")
	}

	// Retrieve updated member
	updatedMember, err := s.repo.GetMember(ctx, req.Member.MemberId)
	if err != nil {
		log.Printf("Error retrieving updated member: %v", err)
		return nil, status.Error(codes.Internal, "failed to retrieve updated member")
	}

	return &pb.UpdateMemberResponse{
		Member: updatedMember,
	}, nil
}

// GetMemberCard retrieves member insurance card
func (s *MemberService) GetMemberCard(ctx context.Context, req *pb.GetMemberCardRequest) (*pb.GetMemberCardResponse, error) {
	log.Printf("GetMemberCard called for member ID: %s, coverage type: %s", req.MemberId, req.CoverageType)

	if req.MemberId == "" {
		return nil, status.Error(codes.InvalidArgument, "member_id is required")
	}

	// Default to medical if not specified
	coverageType := req.CoverageType
	if coverageType == pb.CoverageType_COVERAGE_TYPE_UNSPECIFIED {
		coverageType = pb.CoverageType_MEDICAL
	}

	card, err := s.repo.GetMemberCard(ctx, req.MemberId, coverageType)
	if err != nil {
		log.Printf("Error retrieving member card: %v", err)
		return nil, status.Error(codes.NotFound, "member card not found")
	}

	return &pb.GetMemberCardResponse{
		Card: card,
	}, nil
}

// ListDependents lists all dependents for a member
func (s *MemberService) ListDependents(ctx context.Context, req *pb.ListDependentsRequest) (*pb.ListDependentsResponse, error) {
	log.Printf("ListDependents called for member ID: %s", req.MemberId)

	if req.MemberId == "" {
		return nil, status.Error(codes.InvalidArgument, "member_id is required")
	}

	dependents, err := s.repo.ListDependents(ctx, req.MemberId)
	if err != nil {
		log.Printf("Error listing dependents: %v", err)
		return nil, status.Error(codes.Internal, "failed to list dependents")
	}

	return &pb.ListDependentsResponse{
		Dependents: dependents,
	}, nil
}