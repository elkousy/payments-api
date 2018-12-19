package payments

import (
	"net/http"

	apierrors "github.com/elkousy/payments-api/utility/errors"
)

var (
	// ErrInvalidPaymentID is thrown when payment ID is not valid
	ErrInvalidPaymentID = apierrors.APIError{
		ResponseCode: http.StatusBadRequest,
		Message:      "invalid payment ID",
	}

	// ErrInvalidPaymentPayload is thrown when payment ID is not valid
	ErrInvalidPaymentPayload = apierrors.APIError{
		ResponseCode: http.StatusBadRequest,
		Message:      "some payment fields are missing",
	}

	// ErrNotFound is thrown when ressource requested was not found
	ErrNotFound = apierrors.APIError{
		ResponseCode: http.StatusNotFound,
		Message:      "payment not found",
	}

	// ErrInternalServer is thrown when there an unexpected server error
	ErrInternalServer = apierrors.APIError{
		ResponseCode: http.StatusInternalServerError,
		Message:      "an internal server error occurred",
	}

	// ErrInvalidBody is thrown when the json is not a good format
	ErrInvalidBody = apierrors.APIError{
		ResponseCode: http.StatusBadRequest,
		Message:      "invalid body",
	}
)
