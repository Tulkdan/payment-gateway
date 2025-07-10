package providers_test

import (
	"context"
	"testing"
	"time"

	"github.com/Tulkdan/payment-gateway/internal/domain"
	"github.com/Tulkdan/payment-gateway/internal/providers"
	"go.uber.org/zap"
)

type SpyProvider struct {
	Calls    uint
	Timeout  time.Duration
	Response *domain.Provider
}

func (s *SpyProvider) Charge(ctx context.Context, request *domain.Payment) (*domain.Provider, error) {
	time.Sleep(s.Timeout)
	s.Calls++

	return s.Response, nil
}

func (s *SpyProvider) GetName() string {
	return "Mock"
}

func TestProvider(t *testing.T) {
	t.Run("should make request for first provider", func(t *testing.T) {
		logger, _ := zap.NewDevelopment()
		defer logger.Sync()
		sugar := logger.Sugar()

		spyFirst := &SpyProvider{Timeout: 10 * time.Millisecond, Response: &domain.Provider{Description: "First"}}
		spySecond := &SpyProvider{Timeout: 10 * time.Millisecond, Response: &domain.Provider{Description: "Second"}}

		payment, _ := domain.NewPayment(1000, "R$", "", "card", domain.PaymentCard{Number: "", HolderName: "", CVV: "", ExpirationDate: "02/2025", Installments: 1})

		useProvider := providers.ConfigurableUseProvider([]providers.Provider{spyFirst, spySecond}, sugar, 15*time.Millisecond)
		data, err := useProvider.Payment(context.Background(), payment)

		if err != nil {
			t.Fatal("Got error. didn't want one")
		}

		if data != spyFirst.Response {
			t.Fatalf("expected data to be %q, got %q", spyFirst.Response, data)
		}

		assertSpyCalled(t, spyFirst, "spyFirst", 1)
		assertSpyNotCalled(t, spySecond, "spySecond")
	})

	t.Run("should make request for second provider when first provider timeouts", func(t *testing.T) {
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		sugar := logger.Sugar()

		spyFirst := &SpyProvider{Timeout: 20 * time.Millisecond, Response: &domain.Provider{Description: "First"}}
		spySecond := &SpyProvider{Timeout: 10 * time.Millisecond, Response: &domain.Provider{Description: "Second"}}

		payment, _ := domain.NewPayment(1000, "R$", "", "card", domain.PaymentCard{Number: "", HolderName: "", CVV: "", ExpirationDate: "02/2025", Installments: 1})

		useProvider := providers.ConfigurableUseProvider([]providers.Provider{spyFirst, spySecond}, sugar, 15*time.Millisecond)
		data, err := useProvider.Payment(context.Background(), payment)

		if err != nil {
			t.Fatal("Got error. didn't want one")
		}

		if data != spySecond.Response {
			t.Fatalf("expected data to be %q, got %q", spySecond.Response, data)
		}

		assertSpyCalled(t, spyFirst, "spyFirst", 1)
		assertSpyCalled(t, spySecond, "spySecond", 1)
	})

	t.Run("should return error when all providers timeout", func(t *testing.T) {
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		sugar := logger.Sugar()

		spyFirst := &SpyProvider{Timeout: 20 * time.Millisecond, Response: &domain.Provider{Description: "First"}}
		spySecond := &SpyProvider{Timeout: 20 * time.Millisecond, Response: &domain.Provider{Description: "Second"}}

		payment, _ := domain.NewPayment(1000, "R$", "", "card", domain.PaymentCard{Number: "", HolderName: "", CVV: "", ExpirationDate: "02/2025", Installments: 1})

		useProvider := providers.ConfigurableUseProvider([]providers.Provider{spyFirst, spySecond}, sugar, 5*time.Millisecond)
		data, err := useProvider.Payment(context.Background(), payment)

		if data != nil {
			t.Fatalf("Got data but didn't expected one, got %q", data)
		}

		if err.Error() != "Timeout" {
			t.Fatalf("expected error to be %s, got %s", "Timeout", err.Error())
		}

		assertSpyNotCalled(t, spyFirst, "spyFirst")
		assertSpyNotCalled(t, spySecond, "spySecond")
	})
}

func assertSpyCalled(t testing.TB, spy *SpyProvider, name string, wantTimes uint) {
	t.Helper()

	if spy.Calls != wantTimes {
		t.Fatalf("not enough calls to %s, want %d got %d", name, wantTimes, spy.Calls)
	}
}

func assertSpyNotCalled(t testing.TB, spy *SpyProvider, name string) {
	t.Helper()

	if spy.Calls != 0 {
		t.Fatalf("%s has been called, want 0 got %d", name, spy.Calls)
	}
}
