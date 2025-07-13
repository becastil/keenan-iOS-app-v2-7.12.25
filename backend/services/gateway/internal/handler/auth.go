package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sydney-health-clone/backend/shared/config"
	"github.com/sydney-health-clone/backend/shared/logger"
	
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

type contextKey string

const (
	UserContextKey contextKey = "user"
)

type UserClaims struct {
	MemberID string `json:"member_id"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func AuthMiddleware(authConfig config.AuthConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip auth for health check
			if r.URL.Path == "/health" {
				next.ServeHTTP(w, r)
				return
			}
			
			// Extract token from Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				respondUnauthorized(w, "Missing authorization header")
				return
			}
			
			// Check Bearer prefix
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				respondUnauthorized(w, "Invalid authorization header format")
				return
			}
			
			tokenString := parts[1]
			
			// Parse and validate token
			token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
				// Validate signing method
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return []byte(authConfig.JWTSecret), nil
			})
			
			if err != nil {
				logger.Debug("Token validation failed", zap.Error(err))
				respondUnauthorized(w, "Invalid token")
				return
			}
			
			claims, ok := token.Claims.(*UserClaims)
			if !ok || !token.Valid {
				respondUnauthorized(w, "Invalid token claims")
				return
			}
			
			// Add user claims to context
			ctx := context.WithValue(r.Context(), UserContextKey, claims)
			r = r.WithContext(ctx)
			
			next.ServeHTTP(w, r)
		})
	}
}

func GetUserClaims(ctx context.Context) (*UserClaims, bool) {
	claims, ok := ctx.Value(UserContextKey).(*UserClaims)
	return claims, ok
}

func respondUnauthorized(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}