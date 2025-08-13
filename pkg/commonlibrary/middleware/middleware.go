package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/context"
)

// AuthMiddleware returns an HTTP middleware that authenticates a request using a JWT.
// It expects the Authorization header in the form "Bearer <token>".
func AuthMiddleware(jwtSecret []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 1. Retrieve the Authorization header.
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
				return
			}

			// 2. Expect header to be in the form "Bearer <token>"
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
				http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
				return
			}

			tokenString := parts[1]

			// 3. Parse and validate the JWT token.
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Ensure token signing method is HMAC (adjust if you use another method)
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				return jwtSecret, nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// 4. Extract user_id from token claims.
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			// Assume that the token includes a "user_id" claim, stored as a number.
			uidFloat, ok := claims["user_id"].(float64)
			if !ok {
				http.Error(w, "user_id not found in token", http.StatusUnauthorized)
				return
			}

			userID := int(uidFloat)

			// 5. Set the userID in the context.
			ctx := context.SetUserIDInt(r.Context(), userID)

			// 6. Pass the updated request to the next handler.
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// AuthMiddlewareString returns an HTTP middleware that authenticates a request using a JWT
// with a string-based user_id (e.g., UUID stored as string).
func AuthMiddlewareString(jwtSecret []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 1. Get Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
				return
			}

			// 2. Expect header in the format: "Bearer <token>"
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
				http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
				return
			}

			tokenString := parts[1]

			// 3. Parse and validate the JWT token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				return jwtSecret, nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// 4. Extract user_id from token claims as string
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			uid, ok := claims["user_id"].(string)
			if !ok || uid == "" {
				http.Error(w, "user_id not found or not a string in token", http.StatusUnauthorized)
				return
			}

			// 5. Set string userID in context
			ctx := context.SetUserID(r.Context(), uid)

			// 6. Proceed to next handler
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
