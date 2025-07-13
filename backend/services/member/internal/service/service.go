package service

import (
	"context"
	"fmt"
	"time"

	pb "github.com/sydney-health-clone/backend/shared/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MemberService struct {
	pb.UnimplementedMemberServiceServer
	members map[string]*pb.Member
}

func NewMemberService() *MemberService {
	svc := &MemberService{
		members: make(map[string]*pb.Member),
	}
	
	// Initialize with mock data
	svc.initMockData()
	
	return svc
}

func (s *MemberService) initMockData() {
	// Mock member data
	s.members["M123456"] = &pb.Member{
		MemberId:      "M123456",
		FirstName:     "John",
		LastName:      "Doe",
		MiddleName:    "Michael",
		DateOfBirth:   timestamppb.New(time.Date(1985, 6, 15, 0, 0, 0, 0, time.UTC)),
		Email:         "john.doe@email.com",
		Phone:         "+1-555-123-4567",
		Address: &pb.Address{
			Street1: "123 Main Street",
			Street2: "Apt 4B",
			City:    "San Francisco",
			State:   "CA",
			ZipCode: "94105",
			Country: "USA",
		},
		GroupNumber:    "GRP001234",
		SubscriberId:   "SUB123456",
		EnrollmentDate: timestamppb.New(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
		ActiveCoverages: []pb.CoverageType{
			pb.CoverageType_COVERAGE_TYPE_MEDICAL,
			pb.CoverageType_COVERAGE_TYPE_DENTAL,
			pb.CoverageType_COVERAGE_TYPE_VISION,
			pb.CoverageType_COVERAGE_TYPE_PHARMACY,
		},
	}
	
	// Add dependents
	s.members["M123457"] = &pb.Member{
		MemberId:      "M123457",
		FirstName:     "Jane",
		LastName:      "Doe",
		DateOfBirth:   timestamppb.New(time.Date(1987, 3, 22, 0, 0, 0, 0, time.UTC)),
		Email:         "jane.doe@email.com",
		Phone:         "+1-555-123-4568",
		Address: &pb.Address{
			Street1: "123 Main Street",
			Street2: "Apt 4B",
			City:    "San Francisco",
			State:   "CA",
			ZipCode: "94105",
			Country: "USA",
		},
		GroupNumber:    "GRP001234",
		SubscriberId:   "SUB123456-01",
		EnrollmentDate: timestamppb.New(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
		ActiveCoverages: []pb.CoverageType{
			pb.CoverageType_COVERAGE_TYPE_MEDICAL,
			pb.CoverageType_COVERAGE_TYPE_DENTAL,
			pb.CoverageType_COVERAGE_TYPE_VISION,
			pb.CoverageType_COVERAGE_TYPE_PHARMACY,
		},
	}
	
	s.members["M123458"] = &pb.Member{
		MemberId:      "M123458",
		FirstName:     "Jimmy",
		LastName:      "Doe",
		DateOfBirth:   timestamppb.New(time.Date(2010, 7, 10, 0, 0, 0, 0, time.UTC)),
		Email:         "",
		Phone:         "",
		Address: &pb.Address{
			Street1: "123 Main Street",
			Street2: "Apt 4B",
			City:    "San Francisco",
			State:   "CA",
			ZipCode: "94105",
			Country: "USA",
		},
		GroupNumber:    "GRP001234",
		SubscriberId:   "SUB123456-02",
		EnrollmentDate: timestamppb.New(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
		ActiveCoverages: []pb.CoverageType{
			pb.CoverageType_COVERAGE_TYPE_MEDICAL,
			pb.CoverageType_COVERAGE_TYPE_DENTAL,
			pb.CoverageType_COVERAGE_TYPE_VISION,
		},
	}
}

func (s *MemberService) GetMember(ctx context.Context, req *pb.GetMemberRequest) (*pb.GetMemberResponse, error) {
	member, exists := s.members[req.MemberId]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "member not found: %s", req.MemberId)
	}
	
	return &pb.GetMemberResponse{
		Member: member,
	}, nil
}

func (s *MemberService) UpdateMember(ctx context.Context, req *pb.UpdateMemberRequest) (*pb.UpdateMemberResponse, error) {
	if req.Member == nil {
		return nil, status.Error(codes.InvalidArgument, "member is required")
	}
	
	_, exists := s.members[req.Member.MemberId]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "member not found: %s", req.Member.MemberId)
	}
	
	// Update member (in real implementation, this would update the database)
	s.members[req.Member.MemberId] = req.Member
	
	return &pb.UpdateMemberResponse{
		Member: req.Member,
	}, nil
}

func (s *MemberService) GetMemberCard(ctx context.Context, req *pb.GetMemberCardRequest) (*pb.GetMemberCardResponse, error) {
	member, exists := s.members[req.MemberId]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "member not found: %s", req.MemberId)
	}
	
	// Generate member card based on coverage type
	card := &pb.MemberCard{
		MemberId:     member.MemberId,
		MemberName:   fmt.Sprintf("%s %s", member.FirstName, member.LastName),
		MemberNumber: member.SubscriberId,
		GroupNumber:  member.GroupNumber,
		PlanName:     "Premium Health Plan",
		IssueDate:    member.EnrollmentDate,
		AdditionalInfo: make(map[string]string),
	}
	
	// Add coverage-specific information
	switch req.CoverageType {
	case pb.CoverageType_COVERAGE_TYPE_MEDICAL:
		card.AdditionalInfo["copay_primary"] = "$20"
		card.AdditionalInfo["copay_specialist"] = "$40"
		card.AdditionalInfo["emergency_room"] = "$150"
		card.AdditionalInfo["provider_network"] = "PPO"
	case pb.CoverageType_COVERAGE_TYPE_PHARMACY:
		card.BinNumber = "610502"
		card.PcnNumber = "9999"
		card.RxGroup = "RX1234"
		card.AdditionalInfo["generic_copay"] = "$10"
		card.AdditionalInfo["brand_copay"] = "$35"
		card.AdditionalInfo["specialty_copay"] = "$75"
	case pb.CoverageType_COVERAGE_TYPE_DENTAL:
		card.AdditionalInfo["preventive"] = "100% covered"
		card.AdditionalInfo["basic"] = "80% covered"
		card.AdditionalInfo["major"] = "50% covered"
		card.AdditionalInfo["orthodontia"] = "50% covered"
	case pb.CoverageType_COVERAGE_TYPE_VISION:
		card.AdditionalInfo["eye_exam"] = "$10 copay"
		card.AdditionalInfo["frames"] = "$150 allowance"
		card.AdditionalInfo["lenses"] = "$20 copay"
		card.AdditionalInfo["contacts"] = "$150 allowance"
	}
	
	return &pb.GetMemberCardResponse{
		Card: card,
	}, nil
}

func (s *MemberService) ListDependents(ctx context.Context, req *pb.ListDependentsRequest) (*pb.ListDependentsResponse, error) {
	// For demo purposes, return hardcoded dependents for primary member
	if req.MemberId != "M123456" {
		return &pb.ListDependentsResponse{
			Dependents: []*pb.Member{},
		}, nil
	}
	
	dependents := []*pb.Member{
		s.members["M123457"], // Spouse
		s.members["M123458"], // Child
	}
	
	return &pb.ListDependentsResponse{
		Dependents: dependents,
	}, nil
}