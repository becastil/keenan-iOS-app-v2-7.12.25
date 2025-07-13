import Foundation

// MARK: - Common Models
struct Address: Codable {
    let street1: String
    let street2: String?
    let city: String
    let state: String
    let zipCode: String
}

enum CoverageType: String, Codable {
    case medical = "MEDICAL"
    case dental = "DENTAL"
    case vision = "VISION"
    case pharmacy = "PHARMACY"
}

// MARK: - Member Models
struct Member: Codable {
    let memberId: String
    let firstName: String
    let lastName: String
    let email: String
    let phone: String
    let dateOfBirth: Date
    let address: Address
    let groupNumber: String
    let subscriberId: String
    let activeCoverages: [CoverageType]
}

struct MemberCard: Codable {
    let memberId: String
    let memberName: String
    let memberNumber: String
    let groupNumber: String
    let planName: String
    let coverageType: CoverageType
    let copayPrimary: String?
    let copaySpecialist: String?
    let copayER: String?
    let deductible: String?
    let outOfPocketMax: String?
}

// MARK: - Benefits Models
struct Benefit: Codable {
    let benefitId: String
    let name: String
    let description: String
    let coverageType: CoverageType
    let isCovered: Bool
    let inNetwork: CoverageLevel?
    let outOfNetwork: CoverageLevel?
}

struct CoverageLevel: Codable {
    let copay: String?
    let coinsurance: String?
    let deductible: String?
    let annualLimit: String?
}

// MARK: - Claims Models
struct Claim: Codable {
    let claimId: String
    let memberId: String
    let providerName: String
    let serviceDate: Date
    let processedDate: Date?
    let status: ClaimStatus
    let coverageType: CoverageType
    let totalCharged: Decimal
    let allowedAmount: Decimal
    let memberResponsibility: Decimal
    let planPaid: Decimal
}

enum ClaimStatus: String, Codable {
    case pending = "PENDING"
    case approved = "APPROVED"
    case denied = "DENIED"
    case processing = "PROCESSING"
    case paid = "PAID"
}

// MARK: - Provider Models
struct Provider: Codable {
    let providerId: String
    let name: String
    let practiceName: String?
    let specialties: [String]
    let address: Address
    let phone: String
    let acceptingNewPatients: Bool
    let inNetwork: Bool
    let rating: Double?
    let distance: Double?
}

// MARK: - Message Models
struct Conversation: Codable {
    let conversationId: String
    let subject: String
    let lastMessage: Message?
    let unreadCount: Int
    let participants: [Participant]
    let createdAt: Date
    let updatedAt: Date
}

struct Message: Codable {
    let messageId: String
    let conversationId: String
    let senderId: String
    let senderName: String
    let content: String
    let isRead: Bool
    let sentAt: Date
}

struct Participant: Codable {
    let participantId: String
    let name: String
    let type: ParticipantType
}

enum ParticipantType: String, Codable {
    case member = "MEMBER"
    case supportAgent = "SUPPORT_AGENT"
    case careCoordinator = "CARE_COORDINATOR"
}