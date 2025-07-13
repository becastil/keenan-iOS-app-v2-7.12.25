import Foundation
import RxSwift

protocol AuthServiceProtocol {
    var isAuthenticated: Bool { get }
    var currentMemberId: String? { get }
    func login(memberId: String, password: String) -> Single<LoginResponse>
    func logout()
}

struct LoginResponse {
    let token: String
    let memberId: String
    let memberName: String
}

class AuthService: AuthServiceProtocol {
    
    static let shared = AuthService()
    
    private let userDefaults = UserDefaults.standard
    private let tokenKey = "authToken"
    private let memberIdKey = "memberId"
    
    var isAuthenticated: Bool {
        return userDefaults.string(forKey: tokenKey) != nil
    }
    
    var currentMemberId: String? {
        return userDefaults.string(forKey: memberIdKey)
    }
    
    func login(memberId: String, password: String) -> Single<LoginResponse> {
        return Single.create { single in
            // Mock authentication
            DispatchQueue.global().asyncAfter(deadline: .now() + 1.0) {
                if memberId == "M123456" && password == "demo" {
                    let response = LoginResponse(
                        token: "mock-jwt-token",
                        memberId: memberId,
                        memberName: "John Doe"
                    )
                    
                    self.userDefaults.set(response.token, forKey: self.tokenKey)
                    self.userDefaults.set(response.memberId, forKey: self.memberIdKey)
                    
                    single(.success(response))
                } else {
                    single(.failure(AuthError.invalidCredentials))
                }
            }
            
            return Disposables.create()
        }
    }
    
    func logout() {
        userDefaults.removeObject(forKey: tokenKey)
        userDefaults.removeObject(forKey: memberIdKey)
    }
}

enum AuthError: Error {
    case invalidCredentials
    case networkError
}