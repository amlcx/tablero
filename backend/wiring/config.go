package wiring

import (
	"fmt"
	"os"
	"strconv"
)

// Backend validates incoming requests by parsing a JWT token from the
// request Authorization header, and then validates it against a JWKS,
// which is provided via an endpoint by the auth provider.
type JWKS struct {
	URL string
}

type Server struct {
	Hostname string
	Port     int
}

type Config struct {
	JWKS   JWKS
	Server Server
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

func getEnvInt(key string, defaultValue int) (int, error) {
	if value := os.Getenv(key); value != "" {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return 0, fmt.Errorf("invalid integer value for %s: %w", key, err)
		}
		return intValue, nil
	}
	return defaultValue, nil
}

func LoadConfig() (*Config, error) {
	serverPort, err := getEnvInt("TABLERO_SERVER_PORT", 1323)
	if err != nil {
		return nil, err
	}

	return &Config{
		JWKS: JWKS{
			URL: getEnv("TABLERO_JWKS_URL", "http://frontend:3001/api/auth/jwks"),
		},

		Server: Server{
			Hostname: getEnv("TABLERO_SERVER_HOSTNAME", "localhost"),
			Port:     serverPort,
		},
	}, nil
}

func ValidateConfig(c *Config) error {
	if c.JWKS.URL == "" {
		return fmt.Errorf("invalid config: empty JWKS url")
	}

	if c.Server.Port <= 0 || c.Server.Port > 65535 {
		return fmt.Errorf("invalid config: port %d exceeds allowed range", c.Server.Port)
	}

	return nil
}
