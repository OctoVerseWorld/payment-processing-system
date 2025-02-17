package bank_handler

import (
	"PaymentProcessingSystem/internal"
	"PaymentProcessingSystem/internal/domain"
	"PaymentProcessingSystem/internal/infra/rest/responses"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// getOne ...
func (b *BankHandler) getOne(w http.ResponseWriter, r *http.Request) {
	bankIDStr := chi.URLParam(r, "id")
	bankID, err := strconv.Atoi(bankIDStr)
	if err != nil {
		responses.RenderErrorResponse(w, r,
			internal.WrapErrorf(err, internal.ErrorCodeUnknown, "strconv.Atoi"))
		return
	}

	bank, err := b.svc.GetOneByID(r.Context(), domain.BankID(bankID))
	if err != nil {
		responses.RenderErrorResponse(w, r, err)
		return
	}

	responses.RenderResponse(w, r, BankResponse{
		ID:             bank.ID,
		PlanetID:       bank.PlanetID,
		OrganizationID: bank.OrganizationID,
		Name:           bank.Name,
	}, http.StatusOK)
}
