package config

import (
	"fmt"
	"strconv"
	utils "test_project/test/pkg"
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
	port, err := strconv.Atoi(utils.GetEnvOrDefault("DB_PORT", "5432"))
	if err != nil {
		port = 5433
	}

	return DBConfig{
		Host:     utils.GetEnvOrDefault("DB_HOST", "localhost"),
		Port:     port,
		User:     utils.GetEnvOrDefault("DB_USER", "postgres"),
		Password: utils.GetEnvOrDefault("DB_PASSWORD", "postgres"),
		DBName:   utils.GetEnvOrDefault("DB_NAME", "ideadb"),
		SSLMode:  utils.GetEnvOrDefault("DB_SSLMODE", "disable"),
	}
}

func (c DBConfig) GetDSNPG() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMode)
}
