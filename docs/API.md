# Sydney Health Clone - API Documentation

## Overview

The Sydney Health Clone API is built using gRPC with HTTP/JSON transcoding for web clients. All API endpoints are accessed through the API Gateway service.

## Base URLs

- **Development**: `http://localhost:8080/api/v1`
- **Staging**: `https://api-staging.sydneyhealth.com/api/v1`
- **Production**: `https://api.sydneyhealth.com/api/v1`

## Authentication

All API requests require JWT authentication except for the login endpoint.

### Headers
```
Authorization: Bearer <jwt_token>
Content-Type: application/json
```

### Login
```http
POST /auth/login
```

Request:
```json
{
  "member_id": "M123456",
  "password": "your_password"
}
```

Response:
```json
{
  "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9...",
  "member_id": "M123456",
  "member_name": "John Doe",
  "expires_at": "2024-01-15T10:30:00Z"
}
```

## Member Service API

### Get Member Profile
```http
GET /members/{memberId}
```

Response:
```json
{
  "member_id": "M123456",
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@email.com",
  "phone": "+1-555-123-4567",
  "date_of_birth": "1985-06-15",
  "address": {
    "street1": "123 Main Street",
    "street2": "Apt 4B",
    "city": "San Francisco",
    "state": "CA",
    "zip_code": "94105"
  },
  "group_number": "GRP001234",
  "subscriber_id": "SUB123456",
  "active_coverages": ["MEDICAL", "DENTAL", "VISION", "PHARMACY"]
}
```

### Update Member Profile
```http
PUT /members/{memberId}
```

Request:
```json
{
  "email": "new.email@email.com",
  "phone": "+1-555-987-6543",
  "address": {
    "street1": "456 New Street",
    "city": "San Francisco",
    "state": "CA",
    "zip_code": "94105"
  }
}
```

### Get Member ID Card
```http
GET /members/{memberId}/card?coverage_type=MEDICAL
```

Response:
```json
{
  "member_id": "M123456",
  "member_name": "John Doe",
  "member_number": "SUB123456",
  "group_number": "GRP001234",
  "plan_name": "Premium Health Plan",
  "coverage_type": "MEDICAL",
  "copay_primary": "$20",
  "copay_specialist": "$40",
  "copay_er": "$150",
  "deductible": "$1,500",
  "out_of_pocket_max": "$6,000",
  "card_image_url": "https://api.sydneyhealth.com/cards/M123456_MEDICAL.png"
}
```

### List Dependents
```http
GET /members/{memberId}/dependents
```

Response:
```json
{
  "dependents": [
    {
      "member_id": "M123457",
      "first_name": "Jane",
      "last_name": "Doe",
      "relationship": "SPOUSE",
      "date_of_birth": "1987-03-22"
    },
    {
      "member_id": "M123458",
      "first_name": "Jimmy",
      "last_name": "Doe",
      "relationship": "CHILD",
      "date_of_birth": "2010-07-10"
    }
  ]
}
```

## Benefits Service API

### Get Benefits Summary
```http
GET /members/{memberId}/benefits?coverage_type=MEDICAL
```

Response:
```json
{
  "benefits": [
    {
      "benefit_id": "BEN001",
      "name": "Primary Care Visit",
      "description": "Visits to your primary care physician",
      "coverage_type": "MEDICAL",
      "is_covered": true,
      "in_network": {
        "copay": "$20",
        "coinsurance": "0%"
      },
      "out_of_network": {
        "deductible": "$500",
        "coinsurance": "40%"
      }
    }
  ]
}
```

### Get Deductible Status
```http
GET /members/{memberId}/deductible?coverage_type=MEDICAL
```

Response:
```json
{
  "coverage_type": "MEDICAL",
  "individual_deductible": "$1,500",
  "individual_met": "$750",
  "family_deductible": "$3,000",
  "family_met": "$1,200",
  "period_start": "2024-01-01",
  "period_end": "2024-12-31"
}
```

### Get Out-of-Pocket Status
```http
GET /members/{memberId}/out-of-pocket?coverage_type=MEDICAL
```

Response:
```json
{
  "coverage_type": "MEDICAL",
  "individual_limit": "$6,000",
  "individual_spent": "$2,500",
  "family_limit": "$12,000",
  "family_spent": "$4,800",
  "period_start": "2024-01-01",
  "period_end": "2024-12-31"
}
```

## Claims Service API

### List Claims
```http
GET /members/{memberId}/claims?status=PAID&start_date=2024-01-01&page_size=20
```

Parameters:
- `status`: Filter by claim status (PENDING, APPROVED, DENIED, PROCESSING, PAID)
- `start_date`: Filter claims after this date
- `end_date`: Filter claims before this date
- `coverage_type`: Filter by coverage type
- `page_size`: Number of results per page (default: 20)
- `page_token`: Token for pagination

Response:
```json
{
  "claims": [
    {
      "claim_id": "CLM001234",
      "member_id": "M123456",
      "provider_name": "Bay Area Medical Center",
      "service_date": "2024-01-15",
      "processed_date": "2024-01-20",
      "status": "PAID",
      "coverage_type": "MEDICAL",
      "total_charged": "$250.00",
      "allowed_amount": "$200.00",
      "member_responsibility": "$25.00",
      "plan_paid": "$175.00",
      "service_description": "Office Visit - Primary Care"
    }
  ],
  "next_page_token": "eyJvZmZzZXQiOjIwfQ==",
  "total_count": 45
}
```

### Get Claim Details
```http
GET /claims/{claimId}
```

Response includes full claim details with line items.

### Get Cost Estimate
```http
POST /members/{memberId}/cost-estimate
```

Request:
```json
{
  "procedure_code": "99213",
  "provider_id": "PRV001234",
  "zip_code": "94105"
}
```

Response:
```json
{
  "procedure_code": "99213",
  "procedure_name": "Office Visit - Established Patient",
  "in_network_estimate": "$150-$200",
  "out_of_network_estimate": "$250-$350",
  "your_cost_estimate": "$20",
  "cost_breakdown": [
    {
      "description": "Copay",
      "amount": "$20"
    }
  ],
  "disclaimer": "This is an estimate only. Actual costs may vary."
}
```

### Submit Claim
```http
POST /members/{memberId}/claims
Content-Type: multipart/form-data
```

Form fields:
- `provider_name`: Provider or facility name
- `service_date`: Date of service (YYYY-MM-DD)
- `amount`: Total charged amount
- `description`: Service description
- `receipt`: Receipt image file (JPEG, PNG, PDF)

## Provider Service API

### Search Providers
```http
GET /providers/search?specialty=cardiology&location=94105&radius=10&in_network=true
```

Parameters:
- `specialty`: Medical specialty
- `location`: ZIP code or city
- `radius`: Search radius in miles
- `provider_name`: Search by name
- `in_network`: Filter by network status
- `accepting_new_patients`: Filter by availability

Response:
```json
{
  "providers": [
    {
      "provider_id": "PRV001234",
      "name": "Dr. Sarah Johnson",
      "practice_name": "Bay Area Cardiology",
      "specialties": ["Cardiology", "Internal Medicine"],
      "address": {
        "street1": "456 Market Street",
        "city": "San Francisco",
        "state": "CA",
        "zip_code": "94105"
      },
      "phone": "+1-555-234-5678",
      "accepting_new_patients": true,
      "in_network": true,
      "rating": 4.8,
      "distance": 2.5
    }
  ],
  "total_count": 25
}
```

### Get Provider Details
```http
GET /providers/{providerId}
```

### Check Network Status
```http
GET /providers/{providerId}/network-status?member_id={memberId}&coverage_type=MEDICAL
```

Response:
```json
{
  "in_network": true,
  "network_tier": "TIER_1",
  "covered_services": [
    "Office Visits",
    "Preventive Care",
    "Diagnostic Tests"
  ]
}
```

## Messaging Service API

### List Conversations
```http
GET /members/{memberId}/conversations?unread_only=true
```

Response:
```json
{
  "conversations": [
    {
      "conversation_id": "CONV001",
      "subject": "Question about benefits",
      "type": "BENEFITS",
      "last_message": {
        "content": "Thank you for the clarification",
        "sent_at": "2024-01-15T10:30:00Z"
      },
      "unread_count": 0,
      "participants": [
        {
          "name": "John Doe",
          "type": "MEMBER"
        },
        {
          "name": "Support Agent",
          "type": "SUPPORT_AGENT"
        }
      ],
      "updated_at": "2024-01-15T10:30:00Z"
    }
  ]
}
```

### Get Conversation Messages
```http
GET /conversations/{conversationId}?page_size=50
```

### Send Message
```http
POST /conversations/{conversationId}/messages
```

Request:
```json
{
  "content": "I have a question about my recent claim",
  "attachments": []
}
```

### Mark Messages as Read
```http
POST /messages/mark-read
```

Request:
```json
{
  "message_ids": ["MSG001", "MSG002", "MSG003"]
}
```

## Error Handling

All errors follow a consistent format:

```json
{
  "error": {
    "code": "INVALID_REQUEST",
    "message": "The request is missing required fields",
    "details": {
      "missing_fields": ["member_id", "coverage_type"]
    }
  }
}
```

### Error Codes

| Code | HTTP Status | Description |
|------|-------------|-------------|
| UNAUTHORIZED | 401 | Invalid or missing authentication |
| FORBIDDEN | 403 | Access denied to resource |
| NOT_FOUND | 404 | Resource not found |
| INVALID_REQUEST | 400 | Invalid request parameters |
| RATE_LIMITED | 429 | Too many requests |
| INTERNAL_ERROR | 500 | Internal server error |

## Rate Limiting

- **Authenticated requests**: 1000 requests per hour per member
- **Search endpoints**: 100 requests per hour
- **File uploads**: 50 requests per hour

Rate limit headers:
```
X-RateLimit-Limit: 1000
X-RateLimit-Remaining: 999
X-RateLimit-Reset: 1705320000
```

## Webhooks

Configure webhooks for real-time updates:

### Webhook Events
- `claim.status_changed`: Claim status updates
- `message.received`: New message in conversation
- `benefit.changed`: Benefit modifications
- `member.updated`: Member profile changes

### Webhook Payload
```json
{
  "event_id": "evt_1234567890",
  "event_type": "claim.status_changed",
  "timestamp": "2024-01-15T10:30:00Z",
  "data": {
    "claim_id": "CLM001234",
    "old_status": "PENDING",
    "new_status": "APPROVED"
  }
}
```

## SDK Examples

### JavaScript/TypeScript
```javascript
import { SydneyHealthClient } from '@sydneyhealth/sdk';

const client = new SydneyHealthClient({
  apiKey: 'your_api_key',
  environment: 'production'
});

// Get member profile
const member = await client.members.get('M123456');

// Search providers
const providers = await client.providers.search({
  specialty: 'cardiology',
  location: '94105',
  radius: 10
});
```

### Swift (iOS)
```swift
let client = SydneyHealthClient(apiKey: "your_api_key")

// Get member profile
client.members.get(memberId: "M123456") { result in
    switch result {
    case .success(let member):
        print("Member: \(member.firstName) \(member.lastName)")
    case .failure(let error):
        print("Error: \(error)")
    }
}
```

### Kotlin (Android)
```kotlin
val client = SydneyHealthClient(apiKey = "your_api_key")

// Get member profile
client.members.get("M123456")
    .subscribeOn(Schedulers.io())
    .observeOn(AndroidSchedulers.mainThread())
    .subscribe(
        { member -> println("Member: ${member.firstName} ${member.lastName}") },
        { error -> println("Error: $error") }
    )
```