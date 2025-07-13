package com.uber.sydneyhealth.models

import android.os.Parcelable
import com.google.gson.annotations.SerializedName
import kotlinx.parcelize.Parcelize
import java.math.BigDecimal
import java.util.Date

// Common Models
@Parcelize
data class Address(
    val street1: String,
    val street2: String? = null,
    val city: String,
    val state: String,
    val zipCode: String
) : Parcelable

enum class CoverageType {
    @SerializedName("MEDICAL")
    MEDICAL,
    @SerializedName("DENTAL")
    DENTAL,
    @SerializedName("VISION")
    VISION,
    @SerializedName("PHARMACY")
    PHARMACY
}

// Member Models
@Parcelize
data class Member(
    val memberId: String,
    val firstName: String,
    val lastName: String,
    val email: String,
    val phone: String,
    val dateOfBirth: Date,
    val address: Address,
    val groupNumber: String,
    val subscriberId: String,
    val activeCoverages: List<CoverageType>
) : Parcelable

@Parcelize
data class MemberCard(
    val memberId: String,
    val memberName: String,
    val memberNumber: String,
    val groupNumber: String,
    val planName: String,
    val coverageType: CoverageType,
    val copayPrimary: String? = null,
    val copaySpecialist: String? = null,
    val copayER: String? = null,
    val deductible: String? = null,
    val outOfPocketMax: String? = null
) : Parcelable

// Benefits Models
@Parcelize
data class Benefit(
    val benefitId: String,
    val name: String,
    val description: String,
    val coverageType: CoverageType,
    val isCovered: Boolean,
    val inNetwork: CoverageLevel? = null,
    val outOfNetwork: CoverageLevel? = null
) : Parcelable

@Parcelize
data class CoverageLevel(
    val copay: String? = null,
    val coinsurance: String? = null,
    val deductible: String? = null,
    val annualLimit: String? = null
) : Parcelable

// Claims Models
@Parcelize
data class Claim(
    val claimId: String,
    val memberId: String,
    val providerName: String,
    val serviceDate: Date,
    val processedDate: Date? = null,
    val status: ClaimStatus,
    val coverageType: CoverageType,
    val totalCharged: BigDecimal,
    val allowedAmount: BigDecimal,
    val memberResponsibility: BigDecimal,
    val planPaid: BigDecimal
) : Parcelable

enum class ClaimStatus {
    @SerializedName("PENDING")
    PENDING,
    @SerializedName("APPROVED")
    APPROVED,
    @SerializedName("DENIED")
    DENIED,
    @SerializedName("PROCESSING")
    PROCESSING,
    @SerializedName("PAID")
    PAID
}

// Provider Models
@Parcelize
data class Provider(
    val providerId: String,
    val name: String,
    val practiceName: String? = null,
    val specialties: List<String>,
    val address: Address,
    val phone: String,
    val acceptingNewPatients: Boolean,
    val inNetwork: Boolean,
    val rating: Double? = null,
    val distance: Double? = null
) : Parcelable

// Message Models
@Parcelize
data class Conversation(
    val conversationId: String,
    val subject: String,
    val lastMessage: Message? = null,
    val unreadCount: Int,
    val participants: List<Participant>,
    val createdAt: Date,
    val updatedAt: Date
) : Parcelable

@Parcelize
data class Message(
    val messageId: String,
    val conversationId: String,
    val senderId: String,
    val senderName: String,
    val content: String,
    val isRead: Boolean,
    val sentAt: Date
) : Parcelable

@Parcelize
data class Participant(
    val participantId: String,
    val name: String,
    val type: ParticipantType
) : Parcelable

enum class ParticipantType {
    @SerializedName("MEMBER")
    MEMBER,
    @SerializedName("SUPPORT_AGENT")
    SUPPORT_AGENT,
    @SerializedName("CARE_COORDINATOR")
    CARE_COORDINATOR
}

// Authentication Models
data class LoginRequest(
    val memberId: String,
    val password: String
)

data class LoginResponse(
    val token: String,
    val memberId: String,
    val memberName: String
)