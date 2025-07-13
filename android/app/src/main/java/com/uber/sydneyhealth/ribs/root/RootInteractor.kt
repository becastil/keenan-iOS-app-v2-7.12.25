package com.uber.sydneyhealth.ribs.root

import com.uber.rib.core.Bundle
import com.uber.rib.core.Interactor
import com.uber.rib.core.RibInteractor
import com.uber.sydneyhealth.services.AuthService
import io.reactivex.rxjava3.core.Observable
import javax.inject.Inject

@RibInteractor
class RootInteractor : Interactor<RootInteractor.RootPresenter, RootRouter>() {

    @Inject
    lateinit var authService: AuthService

    override fun didBecomeActive(savedInstanceState: Bundle?) {
        super.didBecomeActive(savedInstanceState)
        
        // Check authentication status and route accordingly
        if (authService.isAuthenticated()) {
            router.attachLoggedIn(authService.getCurrentMemberId() ?: "")
        } else {
            router.attachLoggedOut()
        }
    }

    override fun willResignActive() {
        super.willResignActive()
    }

    fun login(memberId: String) {
        router.detachLoggedOut()
        router.attachLoggedIn(memberId)
    }

    fun logout() {
        authService.logout()
        router.detachLoggedIn()
        router.attachLoggedOut()
    }

    interface RootPresenter

    interface LoggedOutListener {
        fun onLogin(memberId: String)
    }

    interface LoggedInListener {
        fun onLogout()
    }

    inner class LoggedOutInteractorListener : LoggedOutListener {
        override fun onLogin(memberId: String) {
            login(memberId)
        }
    }

    inner class LoggedInInteractorListener : LoggedInListener {
        override fun onLogout() {
            logout()
        }
    }
}