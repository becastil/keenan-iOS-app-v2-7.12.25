import RIBs

protocol LoggedInDependency: Dependency {
    var loggedInViewController: LoggedInViewControllable { get }
    var memberService: MemberServiceProtocol { get }
}

final class LoggedInComponent: Component<LoggedInDependency> {
    
    fileprivate var loggedInViewController: LoggedInViewControllable {
        return dependency.loggedInViewController
    }
    
    fileprivate var memberService: MemberServiceProtocol {
        return dependency.memberService
    }
    
    var memberId: String
    
    init(dependency: LoggedInDependency, memberId: String) {
        self.memberId = memberId
        super.init(dependency: dependency)
    }
}

// MARK: - Builder
protocol LoggedInBuildable: Buildable {
    func build(withListener listener: LoggedInListener, memberId: String) -> ViewableRouting
}

protocol LoggedInListener: AnyObject {
    func didLogout()
}

protocol LoggedInViewControllable: ViewControllable {
    // Define methods the router invokes to manipulate the view hierarchy.
}

final class LoggedInBuilder: Builder<LoggedInDependency>, LoggedInBuildable {

    override init(dependency: LoggedInDependency) {
        super.init(dependency: dependency)
    }

    func build(withListener listener: LoggedInListener, memberId: String) -> ViewableRouting {
        let component = LoggedInComponent(dependency: dependency, memberId: memberId)
        let viewController = TabBarController()
        let interactor = LoggedInInteractor(presenter: viewController)
        interactor.listener = listener
        
        let dashboardBuilder = DashboardBuilder(dependency: component)
        let benefitsBuilder = BenefitsBuilder(dependency: component)
        let claimsBuilder = ClaimsBuilder(dependency: component)
        let providersBuilder = ProvidersBuilder(dependency: component)
        let memberCardBuilder = MemberCardBuilder(dependency: component)
        
        return LoggedInRouter(
            interactor: interactor,
            viewController: viewController,
            dashboardBuilder: dashboardBuilder,
            benefitsBuilder: benefitsBuilder,
            claimsBuilder: claimsBuilder,
            providersBuilder: providersBuilder,
            memberCardBuilder: memberCardBuilder
        )
    }
}

// MARK: - Component Extensions
extension LoggedInComponent: DashboardDependency {}
extension LoggedInComponent: BenefitsDependency {}
extension LoggedInComponent: ClaimsDependency {}
extension LoggedInComponent: ProvidersDependency {}
extension LoggedInComponent: MemberCardDependency {}