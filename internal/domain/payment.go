package domain

import (
	"errors"
	"regexp"
	"sync"
	"time"

	"github.com/Tulkdan/payment-gateway/internal/constants"
)

type Status string

const (
	StatusPending  Status = "pending"  // when it needs to send payment to provider
	StatusApproved Status = "approved" // when provider successfully charged
	StatusRejected Status = "rejected" // when provider rejected payment
	StatusFailed   Status = "failed"   // when provider fails to charge
)

type PaymentCard struct {
	Number         string
	HolderName     string
	CVV            string
	ExpirationDate string
	Installments   uint
}

type Payment struct {
	Amount      uint
	Currency    string
	Description string
	PaymentType string
	Card        PaymentCard
	Status      Status
	CreatedAt   time.Time

	mu sync.Mutex
}

func NewPayment(amount uint, currency string, description string, paymentType string, card PaymentCard) (*Payment, error) {
	if paymentType != "card" {
		return nil, errors.New("We only accept payments with type 'card'")
	}

	isoFormatRgx := regexp.MustCompile(`^[A-Z]{3}$`)
	if !isoFormatRgx.Match([]byte(currency)) || !constants.Lookup(currency) {
		return nil, errors.New("Invalid Currency")
	}

	expirationDateRgx := regexp.MustCompile(`\d{2}\/\d{4}`)
	if !expirationDateRgx.Match([]byte(card.ExpirationDate)) {
		return nil, errors.New("Invalid ExpirationDate format")
	}

	return &Payment{
		Amount:      amount,
		Currency:    currency,
		Description: description,
		PaymentType: paymentType,
		Card:        card,
		Status:      StatusPending,
		CreatedAt:   time.Now(),
	}, nil
}

func (p *Payment) UpdateStatus(status Status) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.Status = status
}
