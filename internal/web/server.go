package web

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/Tulkdan/payment-gateway/internal/service"
	"github.com/Tulkdan/payment-gateway/internal/web/handler"
	"github.com/Tulkdan/payment-gateway/internal/web/middleware"
	"go.uber.org/zap"
)

type Server struct {
	port   string
	router *http.ServeMux
	server *http.Server
	logger *zap.Logger

	paymentsService *service.PaymentService
}

func NewServer(paymentsService *service.PaymentService, port string, logger *zap.Logger) *Server {
	return &Server{
		port:            port,
		paymentsService: paymentsService,
		logger:          logger,
	}
}

func (s *Server) ConfigureRouter() {
	mux := http.NewServeMux()

	paymentsHandler := handler.NewPaymentsHandler(s.paymentsService, s.logger)

	mux.HandleFunc("POST /payments", middleware.WithRequestId(paymentsHandler.Create))
	// r.HandleFunc("POST /refunds", func(http.ResponseWriter, *http.Request) {})
	// r.HandleFunc("GET /payments/{id}", func(w http.ResponseWriter, r *http.Request) {
	// id := r.PathValue("id")
	// })

	s.router = mux
}

func (s *Server) Start(ctx context.Context) error {
	s.server = &http.Server{
		Addr:         ":" + s.port,
		Handler:      s.router,
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s.server.ListenAndServe()
}

func (s *Server) Shutdown() error {
	return s.server.Shutdown(context.Background())
}
