package dto

import "github.com/google/uuid"

type PaymentCardInput struct {
	Number         string `json:"number"`
	HolderName     string `json:"holderName"`
	CVV            string `json:"cvv"`
	ExpirationDate string `json:"expirationDate"`
	Installments   uint   `json:"installments"`
}

type PaymentInput struct {
	Amount      uint             `json:"amount"`
	Currency    string           `json:"currency"`
	Description string           `json:"description"`
	PaymentType string           `json:"paymentType"`
	Card        PaymentCardInput `json:"card"`
}

type PaymentOutput struct {
	Id            uuid.UUID `json:"id"`
	CardId        uuid.UUID `json:"cardId"`
	CurrentAmount uint      `json:"currentAmount"`
}

func NewPaymentOutput(id, cardId uuid.UUID, currentAmount uint) *PaymentOutput {
	return &PaymentOutput{
		Id:            id,
		CardId:        cardId,
		CurrentAmount: currentAmount,
	}
}
