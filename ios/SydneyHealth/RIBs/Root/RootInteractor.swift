import RIBs
import RxSwift

protocol RootRouting: ViewableRouting {
    func routeToLoggedIn(memberId: String)
    func routeToLoggedOut()
}

protocol RootPresentable: Presentable {
    var listener: RootPresentableListener? { get set }
}

protocol RootListener: AnyObject {
    // TODO: Declare methods the interactor can invoke to communicate with other RIBs.
}

final class RootInteractor: PresentableInteractor<RootPresentable>, RootInteractable, RootPresentableListener {

    weak var router: RootRouting?
    weak var listener: RootListener?

    override init(presenter: RootPresentable) {
        super.init(presenter: presenter)
        presenter.listener = self
    }

    override func didBecomeActive() {
        super.didBecomeActive()
        
        // Check if user is authenticated
        if AuthService.shared.isAuthenticated {
            router?.routeToLoggedIn(memberId: AuthService.shared.currentMemberId ?? "")
        } else {
            router?.routeToLoggedOut()
        }
    }

    override func willResignActive() {
        super.willResignActive()
    }
    
    // MARK: - LoggedOutListener
    func didLogin(memberId: String) {
        router?.routeToLoggedIn(memberId: memberId)
    }
    
    // MARK: - LoggedInListener
    func didLogout() {
        router?.routeToLoggedOut()
    }
}