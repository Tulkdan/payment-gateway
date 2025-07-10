package web

import (
	"net/http"

	"github.com/Tulkdan/payment-gateway/internal/service"
	"github.com/Tulkdan/payment-gateway/internal/web/handler"
)

type Server struct {
	port   string
	router *http.ServeMux
	server *http.Server

	paymentsService *service.PaymentService
}

func NewServer(paymentsService *service.PaymentService, port string) *Server {
	return &Server{
		port:            port,
		paymentsService: paymentsService,
	}
}

func (s *Server) ConfigureRouter() {
	r := &http.ServeMux{}

	paymentsHandler := handler.NewPaymentsHandler(s.paymentsService)

	r.HandleFunc("POST /payments", paymentsHandler.Create)
	// r.HandleFunc("POST /refunds", func(http.ResponseWriter, *http.Request) {})
	// r.HandleFunc("GET /payments/{id}", func(w http.ResponseWriter, r *http.Request) {
	// id := r.PathValue("id")
	// })

	s.router = r
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.server.ListenAndServe()
}
