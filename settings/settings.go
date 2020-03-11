package settings

import (
	"os"
)

func getEnv(envName, envDefault string) string {
	if envValue := os.Getenv(envName); envValue != "" {
		return envValue
	}

	return envDefault
}

func GetLogLevel() string {
	return getEnv("LOG_LEVEL", "info")
}

func GetServerHost() string {
	return getEnv("SERVER_HOST", "localhost")
}

func GetServerPort() string {
	return getEnv("SERVER_PORT", "9988")
}

func GetServerPublicKey() string {
	return getEnv("SERVER_PUBLIC_KEY", "./certs/cert.pem")
}

func GetServerPrivateKey() string {
	return getEnv("SERVER_PRIVATE_KEY", "./certs/key.pem")
}

func GetTokenSecret() string {
	return getEnv("TOKEN_SECRET", "notSoSecret")
}
