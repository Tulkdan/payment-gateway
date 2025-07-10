package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Tulkdan/payment-gateway/internal/dto"
	"github.com/Tulkdan/payment-gateway/internal/service"
)

type PaymentsHandler struct {
	paymentService *service.PaymentService
}

func NewPaymentsHandler(paymentsService *service.PaymentService) *PaymentsHandler {
	return &PaymentsHandler{
		paymentService: paymentsService,
	}
}

func (p *PaymentsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body dto.PaymentInput
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := p.paymentService.CreatePayment(r.Context(), body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
