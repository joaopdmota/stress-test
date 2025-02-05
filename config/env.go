package config

import (
	"os"
	"strconv"
	"time"
)

var env *Env

type Env struct {
	ApiTimeout  time.Duration
	ApiURL      string
	Requests    int
	Concurrency int
}

func LoadEnv() (*Env, error) {
	if env != nil {
		return env, nil
	}

	ApiTimeout, err := getEnvInt("API_TIMEOUT", 2)
	if err != nil {
		return nil, err
	}

	Requests, err := getEnvInt("REQUESTS", 1)
	if err != nil {
		return nil, err
	}

	Concurrency, err := getEnvInt("CONCURRENCY", 1)
	if err != nil {
		return nil, err
	}

	env = &Env{
		ApiTimeout:  time.Duration(ApiTimeout) * time.Second,
		ApiURL:      os.Getenv("URL"),
		Requests:    Requests,
		Concurrency: Concurrency,
	}

	return env, nil
}

func UnloadEnv() {
	if env != nil {
		env = nil
	}
}

func getEnvBool(key string, defaultValue bool) (bool, error) {
	envValue := os.Getenv(key)
	if envValue == "" {
		return defaultValue, nil
	}

	val, err := strconv.ParseBool(envValue)
	if err != nil {
		return false, err
	}
	return val, nil
}

func getEnvFloat64(key string, defaultValue float64) (float64, error) {
	envValue := os.Getenv(key)
	if envValue == "" {
		return defaultValue, nil
	}

	val, err := strconv.ParseFloat(envValue, 64)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func getEnvInt(key string, defaultValue int) (int, error) {
	envValue := os.Getenv(key)
	if envValue == "" {
		return defaultValue, nil
	}

	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return 0, err
	}

	return val, nil
}
