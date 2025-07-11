package providers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/Tulkdan/payment-gateway/internal/domain"
	"github.com/Tulkdan/payment-gateway/internal/providers"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func TestStripe(t *testing.T) {
	t.Run("should make request to url", func(t *testing.T) {
		logger, _ := zap.NewDevelopment()
		defer logger.Sync()

		id, _ := uuid.Parse("2ee70bcb-5cb9-4412-a35f-c2a15fb88ef1")
		cardId, _ := uuid.Parse("ed6ecd4c-81d5-4e63-bb12-99439ae559e7")
		ctx := context.WithValue(t.Context(), "request-id", uuid.New().String())

		serverResponse := &providers.StripeChargeResponse{
			Id:             id,
			CreatedAt:      time.Now().Format("YYYY-MM-DD"),
			Status:         "paid",
			OriginalAmount: 1000,
			CurrentAmount:  1000,
			Currency:       "BRL",
			Description:    "",
			PaymentMethod:  "card",
			CardId:         cardId,
		}

		server := createServerStripe(serverResponse)
		defer server.Close()

		charge := &domain.Payment{
			Amount:      1000,
			Currency:    "BRL",
			Description: "",
			PaymentType: "card",
			Card:        domain.PaymentCard{Number: "", HolderName: "", CVV: "", ExpirationDate: "02/2025", Installments: 1},
			Status:      "pending",
			CreatedAt:   time.Now(),
		}
		want := &domain.Provider{
			Id:             id,
			CreatedAt:      time.Now().Format("YYYY-MM-DD"),
			Status:         domain.StatusApproved,
			OriginalAmount: 1000,
			CurrentAmount:  1000,
			Currency:       "BRL",
			Description:    "",
			PaymentMethod:  "card",
			CardId:         cardId,
		}

		provider := providers.NewStripeProvider(server.URL, logger)
		response, err := provider.Charge(ctx, charge)

		if err != nil {
			t.Fatalf("got an error but didn't want one %q", err)
		}

		if !reflect.DeepEqual(response, want) {
			t.Errorf("got %q want %q", response, want)
		}
	})
}

func createServerStripe(serverResponse *providers.StripeChargeResponse) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(serverResponse)
	}))
}
