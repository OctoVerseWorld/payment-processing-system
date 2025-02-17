package bank_handler

import (
	"PaymentProcessingSystem/internal"
	"PaymentProcessingSystem/internal/domain"
	"PaymentProcessingSystem/internal/infra/rest/responses"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// UpdateBankNameRequest ...
type UpdateBankNameRequest struct {
	Name string `json:"name"`
}

// updateName ...
func (b *BankHandler) updateName(w http.ResponseWriter, r *http.Request) {
	var req UpdateBankNameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.RenderErrorResponse(w, r,
			internal.WrapErrorf(err, internal.ErrorCodeDecoding, "json decoder"))
		return
	}
	defer r.Body.Close()

	bankIDStr := chi.URLParam(r, "id")
	bankID, err := strconv.Atoi(bankIDStr)
	if err != nil {
		responses.RenderErrorResponse(w, r,
			internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "validation error"))
		return
	}

	err = b.svc.UpdateName(r.Context(), domain.BankID(bankID), req.Name)
	if err != nil {
		responses.RenderErrorResponse(w, r, err)
		return
	}

	responses.RenderResponse(w, r, struct{}{}, http.StatusNoContent)
}
