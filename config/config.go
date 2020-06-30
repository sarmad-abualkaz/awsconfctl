package config

import "os"

type AWSSession struct {
	AWS_REGION string
}

type Config struct {
	AWS AWSSession
}

// Return a new config struct:
func New() *Config {
	return &Config{
		AWS: AWSSession{
			AWS_REGION: getEnv("AWS_REGION", "us-east-1"),
		},
	}
}

// Key function to return environment variable if exists, otherwise return default
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// Initialize config for environmnent
var ConfigEnv = New()
