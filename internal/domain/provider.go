package domain

import "github.com/google/uuid"

type Provider struct {
	Id             uuid.UUID
	CreatedAt      string
	Status         Status
	OriginalAmount uint
	CurrentAmount  uint
	Currency       string
	Description    string
	PaymentMethod  string
	CardId         uuid.UUID
}
