package environment

import (
	"fmt"
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

func CheckEnvVars() {
	checks := []func()error {
		CheckAuthToken,
		CheckAuthSecret,
		CheckClientPort,
	}

	for i := range len(checks) {
		method := checks[i]
		err := method()
		if err != nil {
			log.Fatalf("Env vars check failed on method: %v", err)
		}
	}
}

func GetClientPort() string {
	return strings.TrimSpace(os.Getenv(BAROSA_CLIENT_PORT_KEY))
}

func CheckClientPort() error {
	clientPort := GetClientPort()
	if len(clientPort) == 0 {
		return fmt.Errorf("Client port env var (%s) is empty! we need this provided", BAROSA_CLIENT_PORT_KEY)
	}
	if _, err := strconv.ParseInt(clientPort, 10, 64); err != nil {
		return fmt.Errorf("Client port '%s' is not a valid port: %v", clientPort, err)
	}
	return nil
}

func GetAuthSecret() string {
	return strings.TrimSpace(os.Getenv(AUTH_SECRET_KEY))
}

func CheckAuthSecret() error {
	authSecret := GetAuthSecret()
	authSecretLength := len(authSecret)
	if authSecretLength == 0 {
		return fmt.Errorf("Barosa bearer secret (%s) not found in env vars!", AUTH_SECRET_KEY)	
	} else if authSecretLength < AUTH_SECRET_KEY_MIN_LENGTH{
		return fmt.Errorf("Barosa bearer secret is %d chars, but should be at least %d", authSecretLength, AUTH_SECRET_KEY_MIN_LENGTH)
	}
	return nil
}

func GetAuthToken() string {
	return strings.TrimSpace(os.Getenv(AUTH_ENV_KEY))
}

func CheckAuthToken() error {
	authToken := GetAuthToken() 
	authTokenLength := len(authToken)
	if authTokenLength == 0 {
		return fmt.Errorf("Barosa bearer auth server token (%s) is not present in env vars!", AUTH_ENV_KEY)
	} else if authTokenLength < AUTH_ENV_KEY_MIN_LENGTH {
		return fmt.Errorf("Barosa bearer auth server token is %d chars, but should be at least %d", authTokenLength, AUTH_ENV_KEY_MIN_LENGTH)
	}
	return nil
}
