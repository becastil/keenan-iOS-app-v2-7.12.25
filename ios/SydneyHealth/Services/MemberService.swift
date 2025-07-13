import Foundation
import RxSwift

protocol MemberServiceProtocol {
    func getMember(memberId: String) -> Single<Member>
    func getMemberCard(memberId: String, coverageType: CoverageType) -> Single<MemberCard>
}

class MemberService: MemberServiceProtocol {
    
    func getMember(memberId: String) -> Single<Member> {
        return Single.create { single in
            // Mock member data
            DispatchQueue.global().asyncAfter(deadline: .now() + 0.5) {
                let member = Member(
                    memberId: memberId,
                    firstName: "John",
                    lastName: "Doe",
                    email: "john.doe@email.com",
                    phone: "+1-555-123-4567",
                    dateOfBirth: Date(timeIntervalSince1970: 519091200), // 1985-06-15
                    address: Address(
                        street1: "123 Main Street",
                        street2: "Apt 4B",
                        city: "San Francisco",
                        state: "CA",
                        zipCode: "94105"
                    ),
                    groupNumber: "GRP001234",
                    subscriberId: "SUB123456",
                    activeCoverages: [.medical, .dental, .vision, .pharmacy]
                )
                single(.success(member))
            }
            
            return Disposables.create()
        }
    }
    
    func getMemberCard(memberId: String, coverageType: CoverageType) -> Single<MemberCard> {
        return Single.create { single in
            // Mock member card data
            DispatchQueue.global().asyncAfter(deadline: .now() + 0.5) {
                let card = MemberCard(
                    memberId: memberId,
                    memberName: "John Doe",
                    memberNumber: "SUB123456",
                    groupNumber: "GRP001234",
                    planName: "Premium Health Plan",
                    coverageType: coverageType,
                    copayPrimary: "$20",
                    copaySpecialist: "$40",
                    copayER: "$150",
                    deductible: "$1,500",
                    outOfPocketMax: "$6,000"
                )
                single(.success(card))
            }
            
            return Disposables.create()
        }
    }
}