package providers

import (
	"errors"
	"time"

	"github.com/Tulkdan/payment-gateway/internal/domain"
)

var thirtySecondTimout = 30 * time.Second

type Provider interface {
	Charge(request *domain.Payment) (*domain.Provider, error)
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

func (p *UseProviders) Payment(payment *domain.Payment) (*domain.Provider, error) {
	var err error = nil

	for _, provider := range p.providers {
		select {
		case data := <-charge(payment, provider):
			return data, nil
		case <-time.After(p.timeout):
			err = errors.New("Timeout")
			continue
		}
	}

	return nil, err
}

func charge(charge *domain.Payment, provider Provider) chan *domain.Provider {
	ch := make(chan *domain.Provider)

	go func() {
		response, err := provider.Charge(charge)
		if err != nil {
			close(ch)
			return
		}
		ch <- response
	}()

	return ch
}
