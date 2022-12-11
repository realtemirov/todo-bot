package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Postgres_HOST    string
	Postgres_PORT    string
	Postgres_USER    string
	Postgres_PASS    string
	Postgres_DBNAME  string
	Postgres_SSLMODE string
	Redis_HOST       string
	Redis_PORT       string
	Redis_PASS       string
	Redis_DB         int
	Redis_EXPIRE     int

	TOKEN string
	ADMIN string
}

func Load() (Config, error) {
	conf := Config{}
	if err := godotenv.Load(); err != nil {
		return conf, err
	}
	conf.Postgres_HOST = cast.ToString(getOrDefault("POSTGRES_HOST", "localhost"))
	conf.Postgres_PORT = cast.ToString(getOrDefault("POSTGRES_PORT", "5432"))
	conf.Postgres_USER = cast.ToString(getOrDefault("POSTGRES_USER", "postgres"))
	conf.Postgres_PASS = cast.ToString(getOrDefault("POSTGRES_PASS", "postgres"))
	conf.Postgres_DBNAME = cast.ToString(getOrDefault("POSTGRES_DBNAME", "postgres"))
	conf.Postgres_SSLMODE = cast.ToString(getOrDefault("POSTGRES_SSLMODE", "disable"))

	conf.Redis_HOST = cast.ToString(getOrDefault("Redis_HOST", "localhost"))
	conf.Redis_PORT = cast.ToString(getOrDefault("Redis_PORT", "0000"))
	conf.Redis_PASS = cast.ToString(getOrDefault("Redis_PASS", ""))
	conf.Redis_DB = cast.ToInt(getOrDefault("Redis_DB", 0))
	conf.Redis_EXPIRE = cast.ToInt(getOrDefault("Redis_EXPIRE", 3600))

	conf.ADMIN = cast.ToString(getOrDefault("BOT_ADMIN", "123456789"))
	conf.TOKEN = cast.ToString(getOrDefault("BOT_TOKEN", "1234567ABC:qwerty"))
	return conf, nil
}

func getOrDefault(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
