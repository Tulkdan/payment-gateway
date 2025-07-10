package service

import (
	"context"

	"github.com/Tulkdan/payment-gateway/internal/domain"
	"github.com/Tulkdan/payment-gateway/internal/dto"
	"github.com/Tulkdan/payment-gateway/internal/providers"
)

type PaymentService struct {
	providers *providers.UseProviders
}

func NewPaymentService(providers *providers.UseProviders) *PaymentService {
	return &PaymentService{providers: providers}
}

func (p *PaymentService) CreatePayment(ctx context.Context, input dto.PaymentInput) (*dto.PaymentOutput, error) {
	payment, err := domain.NewPayment(input.Amount, input.Currency, input.Description, input.PaymentType, domain.PaymentCard(input.Card))
	if err != nil {
		return nil, err
	}

	_, err = p.providers.Payment(ctx, payment)
	if err != nil {
		return nil, err
	}

	return &dto.PaymentOutput{Message: "Processed successfully"}, nil
}
