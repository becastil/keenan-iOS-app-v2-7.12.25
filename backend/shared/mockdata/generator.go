package mockdata

import (
	"fmt"
	"math/rand"
	"time"

	pb "github.com/sydney-health-clone/backend/shared/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MockDataGenerator struct {
	rand *rand.Rand
}

func NewMockDataGenerator() *MockDataGenerator {
	return &MockDataGenerator{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Member Data Generators

func (g *MockDataGenerator) GenerateMember(memberID string) *pb.Member {
	firstNames := []string{"John", "Jane", "Michael", "Sarah", "David", "Emily", "Robert", "Lisa"}
	lastNames := []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis"}
	cities := []string{"San Francisco", "Los Angeles", "New York", "Chicago", "Houston", "Phoenix", "Philadelphia"}
	states := []string{"CA", "CA", "NY", "IL", "TX", "AZ", "PA"}
	
	firstName := firstNames[g.rand.Intn(len(firstNames))]
	lastName := lastNames[g.rand.Intn(len(lastNames))]
	cityIdx := g.rand.Intn(len(cities))
	
	return &pb.Member{
		MemberId:    memberID,
		FirstName:   firstName,
		LastName:    lastName,
		MiddleName:  "M",
		DateOfBirth: g.randomDate(1950, 2010),
		Email:       fmt.Sprintf("%s.%s@email.com", firstName, lastName),
		Phone:       g.randomPhone(),
		Address: &pb.Address{
			Street1: fmt.Sprintf("%d Main Street", g.rand.Intn(9999)+1),
			Street2: "",
			City:    cities[cityIdx],
			State:   states[cityIdx],
			ZipCode: fmt.Sprintf("%05d", g.rand.Intn(99999)),
			Country: "USA",
		},
		GroupNumber:    fmt.Sprintf("GRP%06d", g.rand.Intn(999999)),
		SubscriberId:   fmt.Sprintf("SUB%06d", g.rand.Intn(999999)),
		EnrollmentDate: g.randomDate(2015, 2023),
		ActiveCoverages: []pb.CoverageType{
			pb.CoverageType_COVERAGE_TYPE_MEDICAL,
			pb.CoverageType_COVERAGE_TYPE_DENTAL,
			pb.CoverageType_COVERAGE_TYPE_VISION,
			pb.CoverageType_COVERAGE_TYPE_PHARMACY,
		},
	}
}

// Benefit Data Generators

func (g *MockDataGenerator) GenerateBenefits() []*pb.Benefit {
	benefits := []*pb.Benefit{
		{
			BenefitId:    "BEN001",
			Name:         "Primary Care Visit",
			Description:  "Visits to your primary care physician for routine care",
			CoverageType: pb.CoverageType_COVERAGE_TYPE_MEDICAL,
			IsCovered:    true,
			InNetwork: &pb.CoverageLevel{
				Copay: &pb.Money{Cents: 2000, Currency: "USD"},
				CoinsurancePercentage: 0,
			},
			OutOfNetwork: &pb.CoverageLevel{
				Deductible: &pb.Money{Cents: 50000, Currency: "USD"},
				CoinsurancePercentage: 40,
			},
		},
		{
			BenefitId:    "BEN002",
			Name:         "Specialist Visit",
			Description:  "Visits to medical specialists",
			CoverageType: pb.CoverageType_COVERAGE_TYPE_MEDICAL,
			IsCovered:    true,
			InNetwork: &pb.CoverageLevel{
				Copay: &pb.Money{Cents: 4000, Currency: "USD"},
				CoinsurancePercentage: 0,
			},
			OutOfNetwork: &pb.CoverageLevel{
				Deductible: &pb.Money{Cents: 50000, Currency: "USD"},
				CoinsurancePercentage: 40,
			},
		},
		{
			BenefitId:    "BEN003",
			Name:         "Emergency Room",
			Description:  "Emergency room visits",
			CoverageType: pb.CoverageType_COVERAGE_TYPE_MEDICAL,
			IsCovered:    true,
			InNetwork: &pb.CoverageLevel{
				Copay: &pb.Money{Cents: 15000, Currency: "USD"},
				CoinsurancePercentage: 20,
			},
		},
		{
			BenefitId:    "BEN004",
			Name:         "Preventive Care",
			Description:  "Annual checkups, immunizations, and screenings",
			CoverageType: pb.CoverageType_COVERAGE_TYPE_MEDICAL,
			IsCovered:    true,
			InNetwork: &pb.CoverageLevel{
				Copay: &pb.Money{Cents: 0, Currency: "USD"},
				CoinsurancePercentage: 0,
			},
		},
		{
			BenefitId:    "BEN005",
			Name:         "Routine Dental Exam",
			Description:  "Dental cleanings and exams (twice per year)",
			CoverageType: pb.CoverageType_COVERAGE_TYPE_DENTAL,
			IsCovered:    true,
			InNetwork: &pb.CoverageLevel{
				Copay: &pb.Money{Cents: 0, Currency: "USD"},
				CoinsurancePercentage: 0,
			},
		},
		{
			BenefitId:    "BEN006",
			Name:         "Basic Dental Services",
			Description:  "Fillings, extractions, and basic procedures",
			CoverageType: pb.CoverageType_COVERAGE_TYPE_DENTAL,
			IsCovered:    true,
			InNetwork: &pb.CoverageLevel{
				CoinsurancePercentage: 20,
			},
		},
		{
			BenefitId:    "BEN007",
			Name:         "Eye Exam",
			Description:  "Annual comprehensive eye examination",
			CoverageType: pb.CoverageType_COVERAGE_TYPE_VISION,
			IsCovered:    true,
			InNetwork: &pb.CoverageLevel{
				Copay: &pb.Money{Cents: 1000, Currency: "USD"},
			},
		},
		{
			BenefitId:    "BEN008",
			Name:         "Prescription Eyewear",
			Description:  "Eyeglasses or contact lenses",
			CoverageType: pb.CoverageType_COVERAGE_TYPE_VISION,
			IsCovered:    true,
			InNetwork: &pb.CoverageLevel{
				AnnualLimit: &pb.Money{Cents: 15000, Currency: "USD"},
			},
		},
	}
	
	return benefits
}

// Provider Data Generators

func (g *MockDataGenerator) GenerateProviders(count int) []*pb.Provider {
	providerNames := []string{
		"Dr. Sarah Johnson", "Dr. Michael Chen", "Dr. Emily Davis", "Dr. Robert Wilson",
		"Dr. Jennifer Martinez", "Dr. David Anderson", "Dr. Lisa Thompson", "Dr. James Lee",
	}
	
	specialties := []string{
		"Primary Care", "Cardiology", "Dermatology", "Orthopedics",
		"Pediatrics", "OB/GYN", "Psychiatry", "Neurology",
	}
	
	practiceNames := []string{
		"Bay Area Medical Group", "City Health Partners", "Premier Medical Associates",
		"Community Health Center", "Regional Medical Clinic", "Family Care Practice",
	}
	
	providers := make([]*pb.Provider, count)
	
	for i := 0; i < count; i++ {
		name := providerNames[g.rand.Intn(len(providerNames))]
		specialty := specialties[g.rand.Intn(len(specialties))]
		
		provider := &pb.Provider{
			ProviderId:   fmt.Sprintf("PRV%06d", i+1),
			Npi:          fmt.Sprintf("%010d", g.rand.Intn(9999999999)),
			FirstName:    name,
			LastName:     "",
			PracticeName: practiceNames[g.rand.Intn(len(practiceNames))],
			Specialties:  []string{specialty},
			Locations: []*pb.ProviderLocation{
				g.generateProviderLocation(),
			},
			AcceptedPlans:        []string{"Premium Health Plan", "Basic Health Plan"},
			Languages:            []string{"English", "Spanish"},
			Gender:               []string{"M", "F"}[g.rand.Intn(2)],
			AcceptingNewPatients: g.rand.Float32() > 0.3,
			Rating:               float64(g.rand.Intn(20)+30) / 10.0, // 3.0 to 5.0
			ReviewCount:          int32(g.rand.Intn(200)),
		}
		
		providers[i] = provider
	}
	
	return providers
}

func (g *MockDataGenerator) generateProviderLocation() *pb.ProviderLocation {
	streets := []string{"Market St", "Mission St", "Broadway", "Main St", "First Ave"}
	cities := []string{"San Francisco", "Oakland", "San Jose", "Berkeley", "Palo Alto"}
	
	return &pb.ProviderLocation{
		LocationId: fmt.Sprintf("LOC%06d", g.rand.Intn(999999)),
		Address: &pb.Address{
			Street1: fmt.Sprintf("%d %s", g.rand.Intn(9999)+1, streets[g.rand.Intn(len(streets))]),
			City:    cities[g.rand.Intn(len(cities))],
			State:   "CA",
			ZipCode: fmt.Sprintf("94%03d", g.rand.Intn(999)),
			Country: "USA",
		},
		Phone:         g.randomPhone(),
		Fax:           g.randomPhone(),
		OfficeHours:   []string{"Mon-Fri: 9:00 AM - 5:00 PM", "Sat: 9:00 AM - 1:00 PM"},
		DistanceMiles: float64(g.rand.Intn(50)) / 10.0,
	}
}

// Claims Data Generators

func (g *MockDataGenerator) GenerateClaims(memberID string, count int) []*pb.Claim {
	providerNames := []string{
		"Bay Area Medical Center", "St. Mary's Hospital", "City Health Clinic",
		"Regional Medical Group", "Community Hospital", "University Medical Center",
	}
	
	serviceDescriptions := []string{
		"Office Visit - Primary Care", "Lab Work - Blood Test", "X-Ray - Chest",
		"Physical Therapy Session", "Specialist Consultation", "Preventive Care Exam",
		"Emergency Room Visit", "MRI Scan", "Prescription Medication",
	}
	
	claims := make([]*pb.Claim, count)
	
	for i := 0; i < count; i++ {
		serviceDate := g.randomDateRecent(180) // Last 6 months
		processedDate := serviceDate.Add(time.Duration(g.rand.Intn(30)) * 24 * time.Hour)
		
		totalCharged := int64(g.rand.Intn(50000) + 5000) // $50 to $500
		allowedAmount := int64(float64(totalCharged) * (0.5 + g.rand.Float64()*0.4))
		
		deductibleApplied := int64(0)
		if g.rand.Float32() > 0.7 {
			deductibleApplied = int64(g.rand.Intn(10000))
		}
		
		copay := int64(0)
		if deductibleApplied == 0 && g.rand.Float32() > 0.5 {
			copay = int64(g.rand.Intn(5000) + 2000) // $20 to $70
		}
		
		coinsurance := int64(float64(allowedAmount-deductibleApplied-copay) * 0.2)
		memberResponsibility := deductibleApplied + copay + coinsurance
		planPaid := allowedAmount - memberResponsibility
		
		status := pb.ClaimStatus_CLAIM_STATUS_PAID
		if g.rand.Float32() > 0.8 {
			status = pb.ClaimStatus_CLAIM_STATUS_PENDING
			planPaid = 0
		}
		
		claim := &pb.Claim{
			ClaimId:              fmt.Sprintf("CLM%08d", i+1000000),
			MemberId:             memberID,
			ProviderName:         providerNames[g.rand.Intn(len(providerNames))],
			ServiceDate:          timestamppb.New(serviceDate),
			ProcessedDate:        timestamppb.New(processedDate),
			Status:               status,
			CoverageType:         pb.CoverageType_COVERAGE_TYPE_MEDICAL,
			TotalCharged:         &pb.Money{Cents: totalCharged, Currency: "USD"},
			AllowedAmount:        &pb.Money{Cents: allowedAmount, Currency: "USD"},
			DeductibleApplied:    &pb.Money{Cents: deductibleApplied, Currency: "USD"},
			Copay:                &pb.Money{Cents: copay, Currency: "USD"},
			Coinsurance:          &pb.Money{Cents: coinsurance, Currency: "USD"},
			MemberResponsibility: &pb.Money{Cents: memberResponsibility, Currency: "USD"},
			PlanPaid:             &pb.Money{Cents: planPaid, Currency: "USD"},
			LineItems: []*pb.ClaimLine{
				{
					ServiceCode:        fmt.Sprintf("%05d", g.rand.Intn(99999)),
					ServiceDescription: serviceDescriptions[g.rand.Intn(len(serviceDescriptions))],
					Quantity:           1,
					ChargedAmount:      &pb.Money{Cents: totalCharged, Currency: "USD"},
					AllowedAmount:      &pb.Money{Cents: allowedAmount, Currency: "USD"},
					PaidAmount:         &pb.Money{Cents: planPaid, Currency: "USD"},
				},
			},
		}
		
		claims[i] = claim
	}
	
	return claims
}

// Message Data Generators

func (g *MockDataGenerator) GenerateConversations(memberID string, count int) []*pb.Conversation {
	subjects := []string{
		"Question about my benefits",
		"Need help finding a provider",
		"Claim status inquiry",
		"Prescription coverage question",
		"Update my contact information",
		"Prior authorization request",
	}
	
	conversations := make([]*pb.Conversation, count)
	
	for i := 0; i < count; i++ {
		conversationType := pb.ConversationType_CONVERSATION_TYPE_GENERAL
		if g.rand.Float32() > 0.5 {
			conversationType = []pb.ConversationType{
				pb.ConversationType_CONVERSATION_TYPE_CLAIMS,
				pb.ConversationType_CONVERSATION_TYPE_BENEFITS,
				pb.ConversationType_CONVERSATION_TYPE_PROVIDER,
			}[g.rand.Intn(3)]
		}
		
		conversation := &pb.Conversation{
			ConversationId: fmt.Sprintf("CONV%06d", i+1),
			MemberId:       memberID,
			Subject:        subjects[g.rand.Intn(len(subjects))],
			Type:           conversationType,
			Participants: []*pb.Participant{
				{
					ParticipantId: memberID,
					Name:          "Member",
					Type:          pb.ParticipantType_PARTICIPANT_TYPE_MEMBER,
				},
				{
					ParticipantId: fmt.Sprintf("AGENT%03d", g.rand.Intn(100)),
					Name:          "Support Agent",
					Type:          pb.ParticipantType_PARTICIPANT_TYPE_SUPPORT_AGENT,
				},
			},
			UnreadCount: int32(g.rand.Intn(3)),
			CreatedAt:   g.randomDateRecent(30),
			UpdatedAt:   g.randomDateRecent(7),
		}
		
		conversations[i] = conversation
	}
	
	return conversations
}

// Helper functions

func (g *MockDataGenerator) randomDate(startYear, endYear int) *timestamppb.Timestamp {
	min := time.Date(startYear, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(endYear, 12, 31, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min
	
	sec := g.rand.Int63n(delta) + min
	return timestamppb.New(time.Unix(sec, 0))
}

func (g *MockDataGenerator) randomDateRecent(days int) *timestamppb.Timestamp {
	now := time.Now()
	daysAgo := g.rand.Intn(days)
	return timestamppb.New(now.AddDate(0, 0, -daysAgo))
}

func (g *MockDataGenerator) randomPhone() string {
	return fmt.Sprintf("+1-%03d-%03d-%04d",
		g.rand.Intn(899)+100,
		g.rand.Intn(899)+100,
		g.rand.Intn(9999))
}