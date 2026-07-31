package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv      string
	DBHost      string
	DBPort      int
	DBUser      string
	DBPassword  string
	DBName      string
	DBSSLMode   string
	ServerPort  string
	UploadDir   string
	CORSOrigins []string

	StorageDriver string

	SupabaseURL    string
	SupabaseKey    string
	SupabaseBucket string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	port, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		panic("DB_PORT must be a valid integer")
	}

	return Config{
		AppEnv:         getEnv("APP_ENV", "development"),
		ServerPort:     getEnv("PORT", getEnv("SERVER_PORT", "8081")),
		DBHost:         requireEnv("DB_HOST"),
		DBPort:         port,
		DBUser:         requireEnv("DB_USER"),
		DBPassword:     requireEnv("DB_PASSWORD"),
		DBName:         getEnv("DB_NAME", "postgres"),
		DBSSLMode:      getEnv("DB_SSLMODE", "require"),
		UploadDir:      getEnv("UPLOAD_DIR", "./uploads/files"),
		StorageDriver:  getEnv("STORAGE_DRIVER", "local"),
		SupabaseURL:    getEnv("SUPABASE_URL", ""),
		SupabaseKey:    getEnv("SUPABASE_SERVICE_KEY", ""),
		SupabaseBucket: getEnv("SUPABASE_BUCKET", "assets"),
		CORSOrigins:    splitCSV(getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:5173")),
	}
}

func loadLocalEnv() {
	appEnv := getEnv("APP_ENV", "development")

	if appEnv == "production" {
		return
	}

	_ = godotenv.Load(".env")
}

func requireEnv(key string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		panic(fmt.Sprintf("missing required environment variable: %s", key))
	}

	return value
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func splitCSV(value string) []string {
	parts := strings.Split(value, ",")
	values := make([]string, 0, len(parts))

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			values = append(values, trimmed)
		}
	}

	return values
}
