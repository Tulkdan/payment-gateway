package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/Tulkdan/payment-gateway/internal/providers"
	"github.com/Tulkdan/payment-gateway/internal/service"
	"github.com/Tulkdan/payment-gateway/internal/web"
	"go.uber.org/zap"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	providers := providers.NewUseProviders([]providers.Provider{
		providers.NewBraintreeProvider("http://" + getEnv("BRAINTREE_URL", "localhost:8001")),
		providers.NewStripeProvider("http://" + getEnv("STRIPE_URL", "localhost:8002")),
	}, logger)
	paymentsService := service.NewPaymentService(providers)

	port := getEnv("PORT", "8000")
	server := web.NewServer(paymentsService, port, logger)
	server.ConfigureRouter()

	srvErr := make(chan error, 1)
	go func() {
		logger.Info("Starting server", zap.String("port", port))
		srvErr <- server.Start(ctx)
	}()

	select {
	case <-srvErr:
		return
	case <-ctx.Done():
		stop()
	}

	server.Shutdown()
}
