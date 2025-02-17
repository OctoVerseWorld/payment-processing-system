package bank_handler

import (
	"PaymentProcessingSystem/internal"
	"PaymentProcessingSystem/internal/domain"
	"PaymentProcessingSystem/internal/infra/rest/responses"

	"encoding/json"
	"net/http"
)

// BankCreateRequest ...
type BankCreateRequest struct {
	ID             domain.BankID `json:"id"`
	PlanetID       int32         `json:"planet_id"`
	OrganizationID int32         `json:"organization_id"`
	Name           string        `json:"name"`
}

// BankCreateResponse ...
type BankCreateResponse struct {
	ID domain.BankID `json:"id"`
}

//-

// create ...
func (b *BankHandler) create(w http.ResponseWriter, r *http.Request) {
	var req BankCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.RenderErrorResponse(w, r,
			internal.WrapErrorf(err, internal.ErrorCodeDecoding, "json decoder"))
		return
	}
	defer r.Body.Close()

	createParams := domain.BankCreateParams{
		ID:             req.ID,
		PlanetID:       req.PlanetID,
		OrganizationID: req.OrganizationID,
		Name:           req.Name,
	}
	err := createParams.Validate()
	if err != nil {
		responses.RenderErrorResponse(w, r,
			internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "validation error"))
		return
	}

	bid, err := b.svc.Create(r.Context(), createParams)
	if err != nil {
		responses.RenderErrorResponse(w, r, err)
		return
	}

	responses.RenderResponse(w, r, BankCreateResponse{ID: bid}, http.StatusCreated)
}
