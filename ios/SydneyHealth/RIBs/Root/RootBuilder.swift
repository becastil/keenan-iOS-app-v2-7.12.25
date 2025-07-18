import RIBs

protocol RootDependency: Dependency {
    // TODO: Declare the set of dependencies required by this RIB, but cannot be
    // created by this RIB.
}

final class RootComponent: Component<RootDependency> {
    // TODO: Declare 'fileprivate' dependencies that are only used by this RIB.
}

// MARK: - Builder
protocol RootBuildable: Buildable {
    func build() -> LaunchRouting
}

final class RootBuilder: Builder<RootDependency>, RootBuildable {

    override init(dependency: RootDependency) {
        super.init(dependency: dependency)
    }

    func build() -> LaunchRouting {
        let component = RootComponent(dependency: dependency)
        let viewController = RootViewController()
        let interactor = RootInteractor(presenter: viewController)
        
        let loggedOutBuilder = LoggedOutBuilder(dependency: component)
        let loggedInBuilder = LoggedInBuilder(dependency: component)
        
        return RootRouter(
            interactor: interactor,
            viewController: viewController,
            loggedOutBuilder: loggedOutBuilder,
            loggedInBuilder: loggedInBuilder
        )
    }
}

// MARK: - Component Extensions
extension RootComponent: LoggedOutDependency {
    var authService: AuthServiceProtocol {
        return AuthService()
    }
}

extension RootComponent: LoggedInDependency {
    var loggedInViewController: LoggedInViewControllable {
        return RootViewController()
    }
    
    var memberService: MemberServiceProtocol {
        return MemberService()
    }
}