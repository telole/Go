package config

import (
    "os"
    "github.com/joho/godotenv"
)

func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
        panic("Gagal membaca file .env")
    }
}

func GetEnv(key string, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}
