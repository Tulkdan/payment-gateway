package providers

import (
	"context"
	"errors"
	"time"

	"github.com/Tulkdan/payment-gateway/internal/domain"
)

var thirtySecondTimout = 30 * time.Second

type Provider interface {
	Charge(ctx context.Context, request *domain.Payment) (*domain.Provider, error)
}

type UseProviders struct {
	providers []Provider
	timeout   time.Duration
}

func NewUseProviders(providers []Provider) *UseProviders {
	return ConfigurableUseProvider(providers, thirtySecondTimout)
}

func ConfigurableUseProvider(providers []Provider, timeout time.Duration) *UseProviders {
	return &UseProviders{
		providers: providers,
		timeout:   timeout,
	}
}

func (p *UseProviders) Payment(ctx context.Context, payment *domain.Payment) (*domain.Provider, error) {
	var err error = nil

	for _, provider := range p.providers {
		select {
		case data := <-charge(ctx, payment, provider):
			return data, nil
		case <-time.After(p.timeout):
			err = errors.New("Timeout")
			continue
		}
	}

	return nil, err
}

func charge(ctx context.Context, charge *domain.Payment, provider Provider) chan *domain.Provider {
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
