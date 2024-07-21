package environment

import (
	"log"
	"github.com/joho/godotenv"
	"strings"
	"os"
	"strconv"
)

const (
	AUTH_ENV_KEY = "BAROSA_BEARER_AUTH"
	AUTH_SECRET_KEY = "BAROSA_BEARER_SECRET"
	BAROSA_CLIENT_PORT_KEY = "BAROSA_CLIENT_PORT"
	AUTH_ENV_KEY_MIN_LENGTH = 12
	AUTH_SECRET_KEY_MIN_LENGTH = 48
)

func Init() {
	log.Println("Loading env vars")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load env vars: %v\n", err)
		return
	}
}

func GetClientPort() string {
	return strings.TrimSpace(os.Getenv(BAROSA_CLIENT_PORT_KEY))
}

func CheckClientPort() {
	clientPort := GetClientPort()
	if len(clientPort) == 0 {
		log.Fatalf("Client port env var (%s) is empty! we need this provided", BAROSA_CLIENT_PORT_KEY)
		return
	}

	if _, err := strconv.ParseInt(clientPort, 10, 64); err != nil {
		log.Fatalf("Client port '%s' is not a valid port: %v", clientPort, err)
		return
	}
}

func GetAuthSecret() string {
	return strings.TrimSpace(os.Getenv(AUTH_SECRET_KEY))
}

func CheckAuthSecret() {
	authSecret := GetAuthSecret()
	authSecretLength := len(authSecret)
	if authSecretLength == 0 {
		log.Fatalf("Barosa bearer secret (%s) not found in env vars!", AUTH_SECRET_KEY)	
		return
	} else if authSecretLength < AUTH_SECRET_KEY_MIN_LENGTH{
		log.Fatalf("Barosa bearer secret is %d chars, but should be at least %d", authSecretLength, AUTH_SECRET_KEY_MIN_LENGTH)
		return
	}
}

func GetAuthToken() string {
	return strings.TrimSpace(os.Getenv(AUTH_ENV_KEY))
}

func CheckAuthToken() {
	authToken := GetAuthToken() 
	authTokenLength := len(authToken)
	if authTokenLength == 0 {
		log.Fatalf("Barosa bearer auth server token (%s) is not present in env vars!", AUTH_ENV_KEY)	
		return
	} else if authTokenLength < AUTH_ENV_KEY_MIN_LENGTH {
		log.Fatalf("Barosa bearer auth server token is %d chars, but should be at least %d", authTokenLength, AUTH_ENV_KEY_MIN_LENGTH)
		return
	}
}
