package providers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Tulkdan/payment-gateway/internal/domain"
	"github.com/google/uuid"
)

type BraintreeProvider struct {
	Url string
}

func NewBraintreeProvider(url string) *BraintreeProvider {
	return &BraintreeProvider{Url: url}
}

type BraintreeChargeCard struct {
	Number         string `json:"number"`
	HolderName     string `json:"holderName"`
	CVV            string `json:"cvv"`
	ExpirationDate string `json:"expirationDate"`
	Installments   uint   `json:"installments"`
}

type BraintreeChargePaymentMethod struct {
	Type string              `json:"type"`
	Card BraintreeChargeCard `json:"card"`
}

type BraintreeCharge struct {
	Amount        uint                         `json:"amount"`
	Currency      string                       `json:"currency"`
	Description   string                       `json:"description"`
	PaymentMethod BraintreeChargePaymentMethod `json:"paymentMethod"`
}

func (b BraintreeProvider) Charge(request *domain.Payment) (*domain.Provider, error) {
	body := b.createChargeBody(request)
	response, err := http.Post(b.Url+"/charges", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return b.responseCharge(response)
}

func (b BraintreeProvider) createChargeBody(request *domain.Payment) []byte {
	toSend := &BraintreeCharge{
		Amount:      request.Amount,
		Currency:    request.Currency,
		Description: request.Description,
		PaymentMethod: BraintreeChargePaymentMethod{
			Type: request.PaymentType,
			Card: BraintreeChargeCard{
				Number:         request.Card.Number,
				HolderName:     request.Card.HolderName,
				CVV:            request.Card.CVV,
				ExpirationDate: request.Card.ExpirationDate,
				Installments:   request.Card.Installments,
			},
		},
	}
	jsonValue, _ := json.Marshal(toSend)
	return jsonValue
}

type BraintreeChargeResponse struct {
	Id             uuid.UUID `json:"id"`
	CreatedAt      string    `json:"createdAt"`
	Status         string    `json:"status"` // authorized failed refunded
	OriginalAmount uint      `json:"originalAmount"`
	CurrentAmount  uint      `json:"currentAmount"`
	Currency       string    `json:"currency"`
	Description    string    `json:"description"`
	PaymentMethod  string    `json:"paymentMethod"`
	CardId         uuid.UUID `json:"cardId"`
}

func (b BraintreeProvider) responseCharge(response *http.Response) (*domain.Provider, error) {
	var data BraintreeChargeResponse
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}

	var status domain.Status
	switch data.Status {
	case "authorized":
		status = domain.StatusApproved
	case "failed":
		status = domain.StatusFailed
	case "refunded":
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
