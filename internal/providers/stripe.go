package providers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/Tulkdan/payment-gateway/internal/domain"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type StripeProvider struct {
	Url    string
	logger *zap.Logger
}

func NewStripeProvider(url string, logger *zap.Logger) *StripeProvider {
	return &StripeProvider{Url: url, logger: logger.Named("StripeProvider")}
}

type StripeChargeCard struct {
	Number         string `json:"number"`
	HolderName     string `json:"holder"`
	CVV            string `json:"cvv"`
	ExpirationDate string `json:"expiration"`
	Installments   uint   `json:"installmentNumber"`
}

type StripeCharge struct {
	Amount      uint             `json:"amount"`
	Currency    string           `json:"currency"`
	Description string           `json:"statementDescriptor"`
	PaymentType string           `json:"paymentType"`
	Card        StripeChargeCard `json:"card"`
}

func (s *StripeProvider) Charge(ctx context.Context, request *domain.Payment) (*domain.Provider, error) {
	url := s.Url + "/transactions"

	s.logger.Debug("Making request to charge",
		zap.String("url", url))

	body := s.createChargeBody(request)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("transaction-id", ctx.Value("request-id").(string))

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return s.responseCharge(response)
}

func (s *StripeProvider) createChargeBody(request *domain.Payment) []byte {
	toSend := &StripeCharge{
		Amount:      request.Amount,
		Currency:    request.Currency,
		Description: request.Description,
		PaymentType: request.PaymentType,
		Card: StripeChargeCard{
			Number:         request.Card.Number,
			HolderName:     request.Card.HolderName,
			CVV:            request.Card.CVV,
			ExpirationDate: request.Card.ExpirationDate,
			Installments:   request.Card.Installments,
		},
	}
	jsonValue, _ := json.Marshal(toSend)
	return jsonValue
}

type StripeChargeResponse struct {
	Id             uuid.UUID `json:"id"`
	CreatedAt      string    `json:"date"`
	Status         string    `json:"status"` // paid failed voided
	OriginalAmount uint      `json:"originalAmount"`
	CurrentAmount  uint      `json:"amount"`
	Currency       string    `json:"currency"`
	Description    string    `json:"statementDescriptor"`
	PaymentMethod  string    `json:"paymentMethod"`
	CardId         uuid.UUID `json:"cardId"`
}

func (s *StripeProvider) responseCharge(response *http.Response) (*domain.Provider, error) {
	var data StripeChargeResponse
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}

	var status domain.Status
	switch data.Status {
	case "paid":
		status = domain.StatusApproved
	case "failed":
		status = domain.StatusFailed
	case "voided":
		status = domain.StatusRejected
	}

	providerResponse := &domain.Provider{
		Id:             data.Id,
		CreatedAt:      data.CreatedAt,
		OriginalAmount: data.OriginalAmount,
		CurrentAmount:  data.CurrentAmount,
		Currency:       data.Currency,
		Description:    data.Description,
		PaymentMethod:  data.PaymentMethod,
		CardId:         data.CardId,
		Status:         status,
	}
	return providerResponse, nil
}

func (s *StripeProvider) GetName() string {
	return "Stripe provider"
}
