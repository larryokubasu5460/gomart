package config

import "os"

type Config struct {
	Port              string
	DBUrl             string
	UserServiceURL    string
	ProductServiceURL string
}

func LoadConfig() Config {
	return Config{
		Port:              getEnv("PORT", "8082"),
		DBUrl:             getEnv("DB_URL", ""),
		UserServiceURL:    getEnv("USER_SERVICE_URL", ""),
		ProductServiceURL: getEnv("PRODUCT_SERVICE_URL", ""),
	}
}

func getEnv(key string, fallback string) string {
	val := os.Getenv(key)

	if val == "" {
		return fallback
	}

	return val
}
