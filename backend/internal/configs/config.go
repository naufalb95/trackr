package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	// Database Config
	DBHost    string
	DBPort    int
	DBName    string
	DBUser    string
	DBPass    string
	DBSSLMode string

	// Server Port
	ServicePort int
}

func LoadConfig() (*ServerConfig, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Config file not found: %v\n", err)
	}

	config := &ServerConfig{
		DBHost:    viper.GetString("DATABASE_HOST"),
		DBPort:    viper.GetInt("DATABASE_PORT"),
		DBName:    viper.GetString("POSTGRES_DB"),
		DBUser:    viper.GetString("POSTGRES_USERNAME"),
		DBPass:    viper.GetString("POSTGRES_PASSWORD"),
		DBSSLMode: viper.GetString("POSTGRES_SSLMODE"),

		ServicePort: viper.GetInt("SERVICE_PORT"),
	}

	return config, nil
}

func (c *ServerConfig) GetDatabaseURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.DBUser,
		c.DBPass,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBSSLMode,
	)
}

func (c *ServerConfig) GetServicePort() string {
	return fmt.Sprintf(":%d", c.ServicePort)
}
