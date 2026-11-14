package config

import "os"

// DB Configuration
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// 設定を読み込んで返す関数
// 戻り値が*Config型のポインターになっているのは、同じ設定をコピーせず渡すため
func Load() *Config {
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "user"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "007_go_practice_db"),
	}
}

// 環境変数があればそれを使い、無ければデフォルトを使う 関数です。
func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}
