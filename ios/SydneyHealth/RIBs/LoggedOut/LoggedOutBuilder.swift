import RIBs

protocol LoggedOutDependency: Dependency {
    var authService: AuthServiceProtocol { get }
}

final class LoggedOutComponent: Component<LoggedOutDependency> {
    fileprivate var authService: AuthServiceProtocol {
        return dependency.authService
    }
}

// MARK: - Builder
protocol LoggedOutBuildable: Buildable {
    func build(withListener listener: LoggedOutListener) -> ViewableRouting
}

protocol LoggedOutListener: AnyObject {
    func didLogin(memberId: String)
}

final class LoggedOutBuilder: Builder<LoggedOutDependency>, LoggedOutBuildable {

    override init(dependency: LoggedOutDependency) {
        super.init(dependency: dependency)
    }

    func build(withListener listener: LoggedOutListener) -> ViewableRouting {
        let component = LoggedOutComponent(dependency: dependency)
        let viewController = LoggedOutViewController()
        let interactor = LoggedOutInteractor(
            presenter: viewController,
            authService: component.authService
        )
        interactor.listener = listener
        
        return LoggedOutRouter(interactor: interactor, viewController: viewController)
    }
}