package domain_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/Tulkdan/payment-gateway/internal/domain"
)

func TestPayment(t *testing.T) {
	t.Run("should validate currency in type ISO 4217", func(t *testing.T) {
		currency := "R$"

		_, got := domain.NewPayment(1000, currency, "", "card", domain.PaymentCard{Number: "", HolderName: "", CVV: "", ExpirationDate: "02/2025", Installments: 1})
		want := "Invalid Currency"

		if got.Error() != want {
			t.Errorf("got %q want %q given %q", got, want, currency)
		}
	})

	t.Run("should validate paymentType received as 'card'", func(t *testing.T) {
		paymentType := "2025/02"

		_, got := domain.NewPayment(1000, "R$", "", paymentType, domain.PaymentCard{Number: "", HolderName: "", CVV: "", ExpirationDate: "02/2025", Installments: 1})
		want := "We only accept payments with type 'card'"

		if got.Error() != want {
			t.Errorf("got %q want %q given %q", got, want, paymentType)
		}
	})

	t.Run("should validate expirationDate from card to be in format DD/YYYY", func(t *testing.T) {
		expirationDate := "2025/02"

		_, got := domain.NewPayment(1000, "BRL", "", "card", domain.PaymentCard{Number: "", HolderName: "", CVV: "", ExpirationDate: expirationDate, Installments: 1})
		want := "Invalid ExpirationDate format"

		if got.Error() != want {
			t.Errorf("got %q want %q given %q", got, want, expirationDate)
		}
	})

	t.Run("should receive a Currency passing data", func(t *testing.T) {
		got, err := domain.NewPayment(1000, "BRL", "", "card", domain.PaymentCard{Number: "", HolderName: "", CVV: "", ExpirationDate: "02/2025", Installments: 1})
		want := &domain.Payment{
			Amount:      1000,
			Currency:    "BRL",
			Description: "",
			PaymentType: "card",
			Card:        domain.PaymentCard{Number: "", HolderName: "", CVV: "", ExpirationDate: "02/2025", Installments: 1},
			Status:      "pending",
			CreatedAt:   time.Now(),
		}

		if err != nil {
			t.Fatal("got an error but didn't want one")
		}

		if reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
