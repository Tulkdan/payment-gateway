package providers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Tulkdan/payment-gateway/internal/domain"
	"github.com/google/uuid"
)

type StripeProvider struct {
	Url string
}

func NewStripeProvider(url string) *StripeProvider {
	return &StripeProvider{Url: url}
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

func (b StripeProvider) Charge(request *domain.Payment) (*domain.Provider, error) {
	body := b.createChargeBody(request)
	response, err := http.Post(b.Url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return b.responseCharge(response)
}

func (b StripeProvider) createChargeBody(request *domain.Payment) []byte {
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

func (b StripeProvider) responseCharge(response *http.Response) (*domain.Provider, error) {
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
