package responses

import (
	"PaymentProcessingSystem/internal"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.opentelemetry.io/otel"
)

const otelName = "/internal/rest"

type Response struct {
	ResultResponse interface{}   `json:"result,omitempty"`
	ErrorResponse  ErrorResponse `json:"error,omitempty"`
}

// ErrorResponse represents a response containing an error details.
type ErrorResponse struct {
	ErrorCode        internal.ErrorCode `json:"code,omitempty"`
	Details          string             `json:"details,omitempty"`
	Debug            string             `json:"debug,omitempty"`
	ValidationErrors validation.Errors  `json:"validation_errors,omitempty"`
}

func RenderErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	var status int
	resp := Response{
		ErrorResponse: ErrorResponse{},
	}

	var ierr *internal.Error

	if errors.As(err, &ierr) {
		resp.ErrorResponse.ErrorCode = ierr.GetCode()
		resp.ErrorResponse.Details = ierr.Error()
		resp.ErrorResponse.Debug = ierr.GetOrig()
		switch resp.ErrorResponse.ErrorCode {
		case internal.ErrorCodeNotFound:
			status = http.StatusNotFound
		case internal.ErrorCodeInvalidArgument:
			var verr validation.Errors
			status = http.StatusUnprocessableEntity
			if errors.As(err, &verr) {
				status = http.StatusUnprocessableEntity
				resp.ErrorResponse.ValidationErrors = verr
			}
		case internal.ErrorCodeUnknown:
			fallthrough
		default:
			status = http.StatusInternalServerError
		}
	} else {
		status = http.StatusInternalServerError
		resp.ErrorResponse.Details = "internal error"
	}

	if err != nil {
		_, span := otel.Tracer(otelName).Start(r.Context(), "renderErrorResponse")
		defer span.End()

		span.RecordError(err)
	}

	render.Status(r, status)
	render.JSON(w, r, &resp)
}
