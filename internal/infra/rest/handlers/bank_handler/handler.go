package bank_handler

import (
	"PaymentProcessingSystem/internal/domain"
	"context"
	"fmt"

	"github.com/go-chi/chi/v5"
)

// bankIDRegEx ...
const bankIDRegEx string = `[0-9][0-9][0-9][0-9]`

// BankHandler ...
type BankHandler struct {
	svc BankService
}

// NewBankHandler ...
func NewBankHandler(svc BankService) *BankHandler {
	return &BankHandler{
		svc: svc,
	}
}

// Register connects the handlers to the handlers.
func (b *BankHandler) Register(r *chi.Mux) {
	const prefix = "/banks"
	r.Post(
		prefix+"",
		b.create,
	)
	r.Put(
		fmt.Sprintf(prefix+"/{bankID:%s}", bankIDRegEx),
		b.updateName,
	)
	r.Delete(
		fmt.Sprintf(prefix+"/{bankID:%s}", bankIDRegEx),
		b.delete,
	)
	r.Get(
		fmt.Sprintf(prefix+"/{bankID:%s}", bankIDRegEx),
		b.getOne,
	)
	r.Get(
		prefix+"/planet/{planetID}",
		b.getAllByPlanetID,
	)
}

//-

// BankService ...
type BankService interface {
	Create(ctx context.Context, params domain.BankCreateParams) (domain.BankID, error)
	UpdateName(ctx context.Context, id domain.BankID, name string) error
	Delete(ctx context.Context, id domain.BankID) error
	GetOneByID(ctx context.Context, id domain.BankID) (domain.Bank, error)
	GetAllByPlanetID(ctx context.Context, planetID int32) ([]domain.Bank, error)
}

//-

// BankResponse ...
type BankResponse struct {
	ID             domain.BankID `json:"id"`
	PlanetID       int32         `json:"planet_id"`
	OrganizationID int32         `json:"organization_id"`
	Name           string        `json:"name"`
}
