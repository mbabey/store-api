package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"storeback/internal/keys"
	"storeback/internal/server"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v82"
)

func main() {
	if os.Getenv("ENV") != "debug" {
		if err := godotenv.Load("./.env"); err != nil {
			log.Fatal("Could not load env: %w", err)
		}
	}

	apiKey := os.Getenv("STRIPE_API_KEY")
	stripe.Key = apiKey

	var port int16
	port64, err := strconv.ParseInt(os.Getenv("PORT"), 10, 16)
	if err != nil {
		port = int16(8000)
	} else {
		port = int16(port64)
	}

	ctx := context.WithValue(context.Background(), keys.StringKey("PORT"), port)
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	s := server.CreateServer(ctx)
	server.RunServer(ctx, s)
}
