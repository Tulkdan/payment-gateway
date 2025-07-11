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
		requestCtx, cancel := context.WithTimeout(ctx, p.timeout)
		defer cancel()

		dataCh, errCh := p.charge(requestCtx, payment, provider)
		select {
		case data := <-dataCh:
			p.logger.Debug("[Payment] Received request successfully",
				zap.String("provider", provider.GetName()),
				zap.Int("attempt", attempts))

			return data, nil
		case error := <-errCh:
			p.logger.Error("[Payment] Received request with error",
				zap.String("provider", provider.GetName()),
				zap.Int("attempt", attempts),
				zap.String("error", error.Error()))

			err = error
			continue
		case <-time.After(p.timeout):
			p.logger.Error("[Payment] Timeout for provider to respond",
				zap.String("provider", provider.GetName()),
				zap.Int("attempt", attempts))

			cancel()

			err = errors.New("Timeout")
			continue
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	return nil, err
}

func (p *UseProviders) charge(ctx context.Context, charge *domain.Payment, provider Provider) (chan *domain.Provider, chan error) {
	ch := make(chan *domain.Provider)
	chError := make(chan error)

	go func() {
		response, err := provider.Charge(ctx, charge)
		if err != nil {
			chError <- err
			return
		}
		ch <- response
	}()

	return ch, chError
}
