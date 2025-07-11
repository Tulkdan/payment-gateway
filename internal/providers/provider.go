package providers

import (
	"context"
	"errors"
	"time"

	"github.com/Tulkdan/payment-gateway/internal/domain"
	"go.uber.org/zap"
)

var thirtySecondTimout = 5 * time.Second

type Provider interface {
	GetName() string
	Charge(ctx context.Context, request *domain.Payment) (*domain.Provider, error)
}

type UseProviders struct {
	providers []Provider
	timeout   time.Duration
	logger    *zap.Logger
}

func NewUseProviders(providers []Provider, logger *zap.Logger) *UseProviders {
	return ConfigurableUseProvider(providers, logger, thirtySecondTimout)
}

func ConfigurableUseProvider(providers []Provider, logger *zap.Logger, timeout time.Duration) *UseProviders {
	return &UseProviders{
		providers: providers,
		logger:    logger,
		timeout:   timeout,
	}
}

func (p *UseProviders) Payment(ctx context.Context, payment *domain.Payment) (*domain.Provider, error) {
	var err error = nil
	attempts := 0

	for _, provider := range p.providers {
		select {
		case data := <-p.charge(ctx, payment, provider):
			p.logger.Debug("[Payment] Received request successfully",
				zap.String("provider", provider.GetName()),
				zap.Int("attempt", attempts))

			return data, nil
		case <-time.After(p.timeout):
			p.logger.Error("[Payment] Timeout for provider to respond",
				zap.String("provider", provider.GetName()),
				zap.Int("attempt", attempts))

			err = errors.New("Timeout")
			continue
		}
	}

	return nil, err
}

func (p *UseProviders) charge(ctx context.Context, charge *domain.Payment, provider Provider) chan *domain.Provider {
	ch := make(chan *domain.Provider)

	go func() {
		response, err := provider.Charge(ctx, charge)
		if err != nil {
			close(ch)
			return
		}
		ch <- response
	}()

	return ch
}
