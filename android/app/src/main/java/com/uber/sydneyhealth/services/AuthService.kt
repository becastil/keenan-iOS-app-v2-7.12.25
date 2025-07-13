package com.uber.sydneyhealth.services

import android.content.Context
import androidx.security.crypto.EncryptedSharedPreferences
import androidx.security.crypto.MasterKeys
import com.uber.sydneyhealth.models.LoginRequest
import com.uber.sydneyhealth.models.LoginResponse
import io.reactivex.rxjava3.core.Single
import java.util.concurrent.TimeUnit

class AuthService(context: Context) {
    
    companion object {
        private const val PREFS_NAME = "sydney_health_secure_prefs"
        private const val KEY_AUTH_TOKEN = "auth_token"
        private const val KEY_MEMBER_ID = "member_id"
        private const val KEY_MEMBER_NAME = "member_name"
    }
    
    private val masterKeyAlias = MasterKeys.getOrCreate(MasterKeys.AES256_GCM_SPEC)
    
    private val encryptedPrefs = EncryptedSharedPreferences.create(
        PREFS_NAME,
        masterKeyAlias,
        context,
        EncryptedSharedPreferences.PrefKeyEncryptionScheme.AES256_SIV,
        EncryptedSharedPreferences.PrefValueEncryptionScheme.AES256_GCM
    )
    
    fun isAuthenticated(): Boolean {
        return getAuthToken() != null
    }
    
    fun getAuthToken(): String? {
        return encryptedPrefs.getString(KEY_AUTH_TOKEN, null)
    }
    
    fun getCurrentMemberId(): String? {
        return encryptedPrefs.getString(KEY_MEMBER_ID, null)
    }
    
    fun getCurrentMemberName(): String? {
        return encryptedPrefs.getString(KEY_MEMBER_NAME, null)
    }
    
    fun login(request: LoginRequest): Single<LoginResponse> {
        return Single.create { emitter ->
            // Mock authentication - in production, this would call the API
            if (request.memberId == "M123456" && request.password == "demo") {
                val response = LoginResponse(
                    token = "mock-jwt-token",
                    memberId = request.memberId,
                    memberName = "John Doe"
                )
                
                // Save credentials
                encryptedPrefs.edit().apply {
                    putString(KEY_AUTH_TOKEN, response.token)
                    putString(KEY_MEMBER_ID, response.memberId)
                    putString(KEY_MEMBER_NAME, response.memberName)
                    apply()
                }
                
                emitter.onSuccess(response)
            } else {
                emitter.onError(AuthException.InvalidCredentials)
            }
        }.delay(1, TimeUnit.SECONDS) // Simulate network delay
    }
    
    fun logout() {
        encryptedPrefs.edit().apply {
            remove(KEY_AUTH_TOKEN)
            remove(KEY_MEMBER_ID)
            remove(KEY_MEMBER_NAME)
            apply()
        }
    }
    
    sealed class AuthException : Exception() {
        object InvalidCredentials : AuthException()
        object NetworkError : AuthException()
        object UnknownError : AuthException()
    }
}