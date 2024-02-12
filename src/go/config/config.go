package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

const (
	EnvKeyEmailSender   = "ENV_EMAIL_SENDER"
	EnvKeyEmailReceiver = "ENV_EMAIL_RECEIVER"
	EnvKeyEmailPassword = "ENV_EMAIL_PASSWORD"
	EnvKeyEmailHost     = "ENV_EMAIL_HOST"
)

func init() {
	AssertEnvVariableNonEmpty(EnvKeyEmailSender)
	AssertEnvVariableNonEmpty(EnvKeyEmailReceiver)
	AssertEnvVariableNonEmpty(EnvKeyEmailPassword)
	AssertEnvVariableNonEmpty(EnvKeyEmailHost)
}

func AssertEnvVariableNonEmpty(key string) {
	if os.Getenv(key) == "" {
		log.Error(fmt.Sprintf("Environment variable %s must be set", key))
		os.Exit(2)
	}
}

func GetEnvironmentVariable(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Error(fmt.Sprintf("Environment variable %s is empty", key))
	}
	return value
}
