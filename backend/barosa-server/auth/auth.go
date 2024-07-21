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
		return []byte(environment.GetAuthSecret()), nil 
	})
	if err != nil {
		return "", fmt.Errorf("Could not parse encoded jwt: %v", err)
	}

	if !token.Valid {
		return "", fmt.Errorf("Parsed jwt is invalid")
	}

	return token.Raw, nil
}

func AuthCheckBearerTokenServerMiddleware(c *gin.Context) {
	rejectAuth := func(c *gin.Context) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Not authorized",
		})
		c.Abort()
	}

	authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
	if len(authHeader) == 0 {	
		rejectAuth(c)
		return
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 {
		rejectAuth(c)
		return
	}

	if authHeaderParts[0] != "Bearer" {
		rejectAuth(c)
		return
	}

	authToken := strings.TrimSpace(authHeaderParts[1])
	if len(authToken) == 0 {
		rejectAuth(c) 
		return
	}

	authToken, err := AuthDecodeJWT(authToken)
	if err != nil {
		log.Printf("Auth JWT decode failed: %v", err)
		rejectAuth(c)
		return
	}

	serverToken := environment.GetAuthToken()
	if authToken != serverToken {
		rejectAuth(c)
		return
	}

	c.Next()
}
