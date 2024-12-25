package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`
	ServerHost  string `mapstructure:"SERVER_HOST"`
	ServerPort  int    `mapstructure:"SERVER_PORT"`

	// Database
	Db    string `mapstructure:"DB"`
	DbDsn string `mapstructure:"DB_DSN"`

	// Email
	SmtpSender string `mapstructure:"SMTP_SENDER"`
	SmtpHost   string `mapstructure:"SMTP_HOST"`
	SmtpPort   int    `mapstructure:"SMTP_PORT"`
	SmtpUser   string `mapstructure:"SMTP_USER"`
	SmtpPass   string `mapstructure:"SMTP_PASS"`
	StarTLS    bool   `mapstructure:"STARTLS"`
}

func LoadConfig(path string) (config Config) {
	v := viper.New()

	v.SetConfigFile(".env")

	v.SetConfigType("env")
	v.AddConfigPath(path)
	v.AutomaticEnv()

	_ = v.ReadInConfig()
	_ = v.Unmarshal(&config)

	return
}
