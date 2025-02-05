package main

import (
	"log"
	"stress-test/config"
	"stress-test/handlers"
	"stress-test/internal/app"
)

func main() {
	env, err := config.LoadEnv()

	if env.ApiURL == "" {
		log.Fatal("URL must be set")
	}

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := app.InitDependencies(env)

	handlers.Run(client, env)
}
