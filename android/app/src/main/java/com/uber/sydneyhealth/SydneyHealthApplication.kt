package com.uber.sydneyhealth

import android.app.Application
import com.google.firebase.FirebaseApp
import com.google.firebase.crashlytics.FirebaseCrashlytics
import com.uber.sydneyhealth.services.ServiceLocator
import timber.log.Timber

class SydneyHealthApplication : Application() {
    
    companion object {
        lateinit var instance: SydneyHealthApplication
            private set
    }
    
    override fun onCreate() {
        super.onCreate()
        instance = this
        
        initializeLogging()
        initializeFirebase()
        initializeServices()
    }
    
    private fun initializeLogging() {
        if (BuildConfig.DEBUG) {
            Timber.plant(Timber.DebugTree())
        } else {
            Timber.plant(CrashReportingTree())
        }
    }
    
    private fun initializeFirebase() {
        FirebaseApp.initializeApp(this)
        
        // Enable Crashlytics in release builds
        FirebaseCrashlytics.getInstance().setCrashlyticsCollectionEnabled(!BuildConfig.DEBUG)
    }
    
    private fun initializeServices() {
        ServiceLocator.initialize(this)
    }
    
    private class CrashReportingTree : Timber.Tree() {
        override fun log(priority: Int, tag: String?, message: String, t: Throwable?) {
            if (priority == android.util.Log.ERROR) {
                FirebaseCrashlytics.getInstance().log(message)
                t?.let { FirebaseCrashlytics.getInstance().recordException(it) }
            }
        }
    }
}