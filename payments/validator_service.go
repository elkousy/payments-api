package payments

import (
	"github.com/satori/go.uuid"

	valid "gopkg.in/go-playground/validator.v9"
)

var tags = map[string]string{
	"required": "is_required",
	"exists":   "should_exist",
}

type validator struct {
	next Service
}

// newValidator returns a new instance of payment service with a model validation layer
func newValidator(svc Service) (Service, error) {
	return validator{next: svc}, nil
}

func (v validator) GetPayment(req GetPaymentRequest) (*GetPaymentResponse, error) {
	if err := validatePaymentID(req.PaymentID); err != nil {
		return nil, ErrInvalidPaymentID //.FromError(err)
	}
	return v.next.GetPayment(req)
}

func (v validator) GetListOfPayments(req GetListOfPaymentsRequest) (*GetListOfPaymentsResponse, error) {
	return v.next.GetListOfPayments(req)
}

func (v validator) PostPayment(req CreatePaymentRequest) (*CreatePaymentResponse, error) {
	err := validatePayload(req.Payment)
	if err != nil {
		return nil, ErrInvalidPaymentPayload //.FromError(err)
	}
	return v.next.PostPayment(req)
}

func (v validator) UpdatePayment(req UpdatePaymentRequest) (*UpdatePaymentResponse, error) {
	if err := validatePaymentID(req.PaymentID); err != nil {
		return nil, ErrInvalidPaymentID//.FromError(err)
	}
	err := validatePayload(req.Payment)
	if err != nil {
		return nil, ErrInvalidPaymentPayload //.FromError(err)
	}
	return v.next.UpdatePayment(req)
}

func (v validator) DeletePayment(req DeletePaymentRequest) (*DeletePaymentResponse, error) {
	if err := validatePaymentID(req.PaymentID); err != nil {
		return nil, ErrInvalidPaymentID //.FromError(err)
	}
	return v.next.DeletePayment(req)
}

func validatePaymentID(id string) error {
	_, err := uuid.FromString(id)
	if err != nil {
		return err
	}
	return nil
}

func validatePayload(p Payment) error {
	err := valid.New().Struct(p)
	if err != nil {
		return err
	}
	return nil
}
