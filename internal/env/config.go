package env

import (
	"fmt"
	"os"
)

type config struct {
	ApiServerAddr string
	DBName        string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPass        string
}

func (cfg *config) DBConn() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
}

var Config = initConfig()

func initConfig() config {
	return config{
		ApiServerAddr: getEnv("API_SERVER_ADDR", ":8080"),
		DBHost:        getEnv("DB_HOST", "127.0.0.1"),
		DBPort:        getEnv("DB_PORT", "5432"),
		DBName:        getEnv("DB_DB", "rss-scraper"),
		DBUser:        getEnv("DB_USER", "postgres"),
		DBPass:        getEnv("DB_PASSWORD", "postgres"),
	}
}

// getEnv returns the value of an environment variable, or a fallback value if it's not set.
func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
