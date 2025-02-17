package handlers

import (
	"PaymentProcessingSystem/internal/infra/db/repository"
	"PaymentProcessingSystem/internal/infra/rest/handlers/bank_handler"
	"PaymentProcessingSystem/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Router ...

type BaseHandler struct {
	DB *pgxpool.Pool
}

// NewBaseHandler ...
func NewBaseHandler(DB *pgxpool.Pool) *BaseHandler {
	return &BaseHandler{
		DB: DB,
	}
}

// Register connects the handlers to the handlers.
func (b *BaseHandler) Register(r *chi.Mux) {
	// Banks
	bankRepo := repository.NewBankRepository(b.DB)
	bankSvc := service.NewBankService(bankRepo)
	bankHandler := bank_handler.NewBankHandler(bankSvc)
	bankHandler.Register(r)
}
