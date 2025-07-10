package main

import (
	"log"
	"os"

	"github.com/Tulkdan/payment-gateway/internal/providers"
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
	providers := providers.NewUseProviders([]providers.Provider{
		providers.NewBraintreeProvider(getEnv("BRAINTREE_URL", "localhost:8001")),
		providers.NewStripeProvider(getEnv("STRIPE_URL", "localhost:8002")),
	})
	paymentsService := service.NewPaymentService(providers)

	port := getEnv("PORT", "8000")
	server := web.NewServer(paymentsService, port)
	server.ConfigureRouter()

	if err := server.Start(); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
