package com.uber.sydneyhealth.services

import android.content.Context
import com.uber.sydneyhealth.services.api.ApiService
import com.uber.sydneyhealth.services.api.createApiService

object ServiceLocator {
    
    private lateinit var applicationContext: Context
    
    val authService: AuthService by lazy {
        AuthService(applicationContext)
    }
    
    val apiService: ApiService by lazy {
        createApiService(authService)
    }
    
    val memberService: MemberService by lazy {
        MemberService(apiService)
    }
    
    val benefitsService: BenefitsService by lazy {
        BenefitsService(apiService)
    }
    
    val claimsService: ClaimsService by lazy {
        ClaimsService(apiService)
    }
    
    val providerService: ProviderService by lazy {
        ProviderService(apiService)
    }
    
    val messagingService: MessagingService by lazy {
        MessagingService(apiService)
    }
    
    fun initialize(context: Context) {
        applicationContext = context.applicationContext
    }
}