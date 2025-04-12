package config

import (
	"fmt"
	"os"
	"strconv"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewDBConfig() DBConfig {
	port, err := strconv.Atoi(getEnvOrDefault("DB_PORT", "5433"))
	if err != nil {
		port = 5433
	}

	return DBConfig{
		Host:     getEnvOrDefault("DB_HOST", "localhost"),
		Port:     port,
		User:     getEnvOrDefault("DB_USER", "postgres"),
		Password: getEnvOrDefault("DB_PASSWORD", "postgres"),
		DBName:   getEnvOrDefault("DB_NAME", "ideadb"),
		SSLMode:  getEnvOrDefault("DB_SSLMODE", "disable"),
	}
}

func (c DBConfig) GetDSNPG() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMode)
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
