package bank_handler

import (
	"PaymentProcessingSystem/internal"
	"PaymentProcessingSystem/internal/domain"
	"PaymentProcessingSystem/internal/infra/rest/responses"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// delete ...
func (b *BankHandler) delete(w http.ResponseWriter, r *http.Request) {
	bankIDStr := chi.URLParam(r, "id")
	bankID, err := strconv.Atoi(bankIDStr)
	if err != nil {
		responses.RenderErrorResponse(w, r,
			internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "validation error"))
		return
	}

	err = b.svc.Delete(r.Context(), domain.BankID(bankID))
	if err != nil {
		responses.RenderErrorResponse(w, r, err)
		return
	}

	responses.RenderResponse(w, r, struct{}{}, http.StatusNoContent)
}
