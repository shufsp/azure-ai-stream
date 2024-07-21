package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"barosa.fun/azure-ai-stream-backend/environment"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthDecodeJWT(encodedJwt string) (string, error) {
	token, err := jwt.Parse(encodedJwt, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(environment.GetAuthSecret()), nil
	})

	if err != nil {
		return "", fmt.Errorf("Could not parse encoded jwt: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if bearerToken, ok := claims["bearerToken"].(string); ok {
			return bearerToken, nil
		} else {
			return "", fmt.Errorf("bearerToken claim not found or is not a string")
		}
	} else {
		return "", fmt.Errorf("Token is invalid or claims could not be extracted")
	}
}

func AuthCheckBearerTokenServerMiddleware(c *gin.Context) {
	rejectAuth := func(c *gin.Context, details string) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not authorized",
			"details": details, 
		})
		c.Abort()
	}

	authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
	if len(authHeader) == 0 {	
		rejectAuth(c, "No auth provided")
		return
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 {
		rejectAuth(c, "Expected bearer with token")
		return
	}

	if authHeaderParts[0] != "Bearer" {
		rejectAuth(c, "Expected bearer format")
		return
	}

	authToken := strings.TrimSpace(authHeaderParts[1])
	if len(authToken) == 0 {
		rejectAuth(c, "Empty auth token!") 
		return
	}

	authToken, err := AuthDecodeJWT(authToken)
	if err != nil {
		log.Printf("Auth JWT decode failed: %v", err)
		rejectAuth(c, "Malformed token")
		return
	}

	serverToken := environment.GetAuthToken()
	if authToken != serverToken {
		rejectAuth(c, "Token isn't valid")
		return
	}

	c.Next()
}

func CorsServerMiddleware(c *gin.Context) {
	requesterPort := strings.Split(c.Request.Header.Get("Origin"), ":")[2]
	if requesterPort != environment.GetClientPort() {
		// the request must come from our front end express app
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusOK)
            return
        }

        c.Next()
}
