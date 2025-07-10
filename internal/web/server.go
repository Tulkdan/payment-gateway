package web

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/Tulkdan/payment-gateway/internal/service"
	"github.com/Tulkdan/payment-gateway/internal/web/handler"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type Server struct {
	port   string
	router http.Handler
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
	mux := http.NewServeMux()

	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		mux.Handle(pattern, handler)
	}

	paymentsHandler := handler.NewPaymentsHandler(s.paymentsService)

	handleFunc("POST /payments", paymentsHandler.Create)
	// r.HandleFunc("POST /refunds", func(http.ResponseWriter, *http.Request) {})
	// r.HandleFunc("GET /payments/{id}", func(w http.ResponseWriter, r *http.Request) {
	// id := r.PathValue("id")
	// })

	s.router = otelhttp.NewHandler(mux, "/")
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
