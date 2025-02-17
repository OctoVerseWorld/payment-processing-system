package bank_handler

import (
	"PaymentProcessingSystem/internal"
	"PaymentProcessingSystem/internal/infra/rest/responses"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// getAllByPlanetID ...
func (b *BankHandler) getAllByPlanetID(w http.ResponseWriter, r *http.Request) {
	planetIDStr := chi.URLParam(r, "planetID")
	planetID, err := strconv.Atoi(planetIDStr)
	if err != nil {
		responses.RenderErrorResponse(w, r,
			internal.WrapErrorf(err, internal.ErrorCodeUnknown, "strconv.Atoi"))
		return
	}

	banks, err := b.svc.GetAllByPlanetID(r.Context(), int32(planetID))
	if err != nil {
		responses.RenderErrorResponse(w, r, err)
		return
	}

	result := make([]BankResponse, len(banks))
	for i, bank := range banks {
		result[i] = BankResponse{
			ID:             bank.ID,
			PlanetID:       bank.PlanetID,
			OrganizationID: bank.OrganizationID,
			Name:           bank.Name,
		}
	}

	responses.RenderResponse(w, r, result, http.StatusOK)
}
