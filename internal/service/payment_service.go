package service

import (
	"context"
	"fmt"

	"github.com/Tulkdan/payment-gateway/internal/domain"
	"github.com/Tulkdan/payment-gateway/internal/dto"
)

type PaymentService struct{}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (p *PaymentService) CreatePayment(ctx context.Context, input dto.PaymentInput) (*dto.PaymentOutput, error) {
	payment, err := domain.NewPayment(input.Amount, input.Currency, input.Description, input.PaymentType, domain.PaymentCard(input.Card))
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v", payment)

	return &dto.PaymentOutput{Message: "Processed successfully"}, nil
}
