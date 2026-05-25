package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	ServerPort       string
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBName           string
	RedisHost        string
	RedisPort        string
	RedisPassword    string
	RedisDB          int
	JWTSecret        string
	JWTExpirationHrs int
	AIProvider       string
	AIModel          string
	AIAPIKey         string
	AIBaseURL        string
	LogLevel         string
}

func Load() *Config {
	_ = godotenv.Load()

	redisDB, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))
	jwtHrs, _ := strconv.Atoi(getEnv("JWT_EXPIRATION_HOURS", "72"))

	cfg := &Config{
		ServerPort:       getEnv("SERVER_PORT", "8080"),
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBPort:           getEnv("DB_PORT", "3306"),
		DBUser:           getEnv("DB_USER", "root"),
		DBPassword:       getEnv("DB_PASSWORD", "123456"),
		DBName:           getEnv("DB_NAME", "ai_project_os"),
		RedisHost:        getEnv("REDIS_HOST", "localhost"),
		RedisPort:        getEnv("REDIS_PORT", "6379"),
		RedisPassword:    getEnv("REDIS_PASSWORD", ""),
		RedisDB:          redisDB,
		JWTSecret:        getEnv("JWT_SECRET", "change-me-in-production"),
		JWTExpirationHrs: jwtHrs,
		AIProvider:       getEnv("AI_PROVIDER", "mock"),
		AIModel:          getEnv("AI_MODEL", "mock-v1"),
		AIAPIKey:         getEnv("AI_API_KEY", ""),
		AIBaseURL:        getEnv("AI_BASE_URL", ""),
		LogLevel:         getEnv("LOG_LEVEL", "info"),
	}

	if cfg.JWTSecret == "change-me-in-production" {
		zap.L().Warn("JWT_SECRET is using default value — change it in production!")
	}
	return cfg
}

func (c *Config) DSN() string {
	return c.DBUser + ":" + c.DBPassword +
		"@tcp(" + c.DBHost + ":" + c.DBPort + ")/" +
		c.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func (c *Config) RedisAddr() string {
	return c.RedisHost + ":" + c.RedisPort
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
