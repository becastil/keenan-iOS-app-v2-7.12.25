package com.uber.sydneyhealth.ribs.root

import com.uber.rib.core.ViewRouter
import com.uber.sydneyhealth.ribs.loggedin.LoggedInBuilder
import com.uber.sydneyhealth.ribs.loggedin.LoggedInRouter
import com.uber.sydneyhealth.ribs.loggedout.LoggedOutBuilder
import com.uber.sydneyhealth.ribs.loggedout.LoggedOutRouter

class RootRouter(
    view: RootView,
    interactor: RootInteractor,
    component: RootBuilder.Component,
    private val loggedOutBuilder: LoggedOutBuilder,
    private val loggedInBuilder: LoggedInBuilder
) : ViewRouter<RootView, RootInteractor>(view, interactor, component) {

    private var loggedOutRouter: LoggedOutRouter? = null
    private var loggedInRouter: LoggedInRouter? = null

    fun attachLoggedOut() {
        loggedOutRouter = loggedOutBuilder.build(view)
        attachChild(loggedOutRouter)
        view.addView(loggedOutRouter?.view)
    }

    fun detachLoggedOut() {
        loggedOutRouter?.let {
            detachChild(it)
            view.removeView(it.view)
            loggedOutRouter = null
        }
    }

    fun attachLoggedIn(memberId: String) {
        loggedInRouter = loggedInBuilder.build(view, memberId)
        attachChild(loggedInRouter)
        view.addView(loggedInRouter?.view)
    }

    fun detachLoggedIn() {
        loggedInRouter?.let {
            detachChild(it)
            view.removeView(it.view)
            loggedInRouter = null
        }
    }

    fun handleBackPress(): Boolean {
        // Delegate back press to child routers
        return loggedInRouter?.handleBackPress() ?: false
    }
}