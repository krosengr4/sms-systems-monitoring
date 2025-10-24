package alerts

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// config struct
type Config struct {
	EmailFrom string `env:"EMAIL_FROM"`
	EmailTo   string `env:"EMAIL_TO"`

	SmtpServer   string `env:"SMTP_SERVER"`
	SmtpPort     string `env:"SMTP_PORT"`
	SmtpUser     string `env:"SMTP_USER"`
	SmtpPassword string `env:"SMTP_PASSWORD"`
}

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
	if c.EmailFrom == "" {
		return fmt.Errorf("EMAIL_FROM is required")
	}

	if c.EmailTo == "" {
		return fmt.Errorf("EMAIL_TO is required")
	}

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
