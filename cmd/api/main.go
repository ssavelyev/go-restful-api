package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ssavelyev/go-restful-api/internal/env"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	addr := os.Getenv("ADDR")

	cfg := config{
		addr: env.GetString(addr, ":8080"),
	}
	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
