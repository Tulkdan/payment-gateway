package main

import (
	"log"
	"os"

	"github.com/Tulkdan/payment-gateway/internal/service"
	"github.com/Tulkdan/payment-gateway/internal/web"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	paymentsService := service.NewPaymentService()

	server := web.NewServer(paymentsService, "8000")
	server.ConfigureRouter()

	if err := server.Start(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
