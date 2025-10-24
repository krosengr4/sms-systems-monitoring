package alerts

import (
	"fmt"
	"time"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// config struct
type Config struct {
	SmtpServer   string `env:"SMTP_SERVER"`
	SmtpPort     string `env:"SMTP_PORT"`
	SmtpUser     string `env:"SMTP_USER"`
	SmtpPassword string `env:"SMTP_PASSWORD"`
}

const (
	MemoryThreshold  = 70               // Memory % Alert Threshold
	MemCheckInterval = 5 * time.Second  // Check Memory every 5s
	BatteryThreshold = 30               // Batter % Alert Threshold
	BatCheckInterval = 60 * time.Second // Check Battery Every Min
)

func SendCPUAlert(cpuUsage float64) {

}

func ConfigSMTP() (*Config, error) {
	// Use godotenv to load env variables
	if err := godotenv.Load(); err != nil {
		log.Warn().Err(err).Msg("No .env file found or failed to load.")
	}

	cfg := &Config{}

	// Use caarlos0/env
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse the environment variables: %w", err)
	}

	// Validate the env variables
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	log.Info().
		Str("SMTP Server:", cfg.SmtpServer).
		Str("SMTP Port:", cfg.SmtpServer).
		Str("SMTP User:", cfg.SmtpServer)

	return cfg, nil

}

// Validate the configuration
func (c *Config) Validate() error {
	if c.SmtpServer == "" {
		return fmt.Errorf("SMTP_SERVER is required")
	}

	if c.SmtpPort == "" {
		return fmt.Errorf("SMTP_PORT is required")
	}

	if c.SmtpUser == "" {
		return fmt.Errorf("SMTP_USER is required")
	}

	if c.SmtpPassword == "" {
		return fmt.Errorf("SMTP_PASSWORD is required")
	}

	return nil
}
