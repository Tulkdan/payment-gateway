package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Tulkdan/payment-gateway/internal/dto"
	"github.com/Tulkdan/payment-gateway/internal/service"
	"go.uber.org/zap"
)

type PaymentsHandler struct {
	paymentService *service.PaymentService

	logger *zap.Logger
}

func NewPaymentsHandler(paymentsService *service.PaymentService, logger *zap.Logger) *PaymentsHandler {
	return &PaymentsHandler{
		paymentService: paymentsService,
		logger:         logger.Named("PaymentHandler"),
	}
}

func (p *PaymentsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body dto.PaymentInput
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		p.logger.Error("Body incomplete",
			zap.String("error", err.Error()),
			zap.String("requestId", r.Context().Value("request-id").(string)))

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := p.paymentService.CreatePayment(r.Context(), body)
	if err != nil {
		p.logger.Error("Failed to create payment",
			zap.String("error", err.Error()),
			zap.String("requestId", r.Context().Value("request-id").(string)))

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.logger.Debug("Processed request",
		zap.Any("response", response),
		zap.String("requestId", r.Context().Value("request-id").(string)))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
