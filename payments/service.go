package payments

import (
	"errors"
	"fmt"
)

// Service defines the payment service
type Service interface {
	GetPayment(req GetPaymentRequest) (*GetPaymentResponse, error)
	GetListOfPayments(req GetListOfPaymentsRequest) (*GetListOfPaymentsResponse, error)
	PostPayment(req CreatePaymentRequest) (*CreatePaymentResponse, error)
	UpdatePayment(req UpdatePaymentRequest) (*UpdatePaymentResponse, error)
	DeletePayment(req DeletePaymentRequest) (*DeletePaymentResponse, error)
}

type service struct {
	repository Repository
}

// NewPaymentService returns a new instance of the payment service
func NewPaymentService(repository Repository) (Service, error) {
	svc, err := newService(repository)
	if err != nil {
		return nil, err
	}

	// add model validator service
	svc, err = newValidator(svc)
	if err != nil {
		return nil, err
	}

	return svc, nil
}

func newService(repository Repository) (Service, error) {
	if repository == nil {
		return nil, errors.New("cannot create new payments service, repository cannot be nil")
	}

	return service{repository: repository}, nil
}

// GetPayment retrieves a specific payment by ID
func (s service) GetPayment(req GetPaymentRequest) (*GetPaymentResponse, error) {
	// get a payment
	p, err := s.repository.GetPayment(req.PaymentID)
	if err != nil {
		return nil, err
	}
	return &GetPaymentResponse{Payment: p}, nil
}

// GetListOfPayments returns a list of payments
func (s service) GetListOfPayments(req GetListOfPaymentsRequest) (*GetListOfPaymentsResponse, error) {
	// get list of payments
	payments, err := s.repository.GetListOfPayments()
	if err != nil {
		return nil, err
	}
	return &GetListOfPaymentsResponse{Data: payments, HateoasLink: HateoasLink{Self: "localhost:8080/v1/payments/"}}, nil
}

// PostPayment inserts a new payment in DB
func (s service) PostPayment(req CreatePaymentRequest) (*CreatePaymentResponse, error) {
	// create payment
	id, err := s.repository.CreatePayment(req.Payment)
	if err != nil {
		return nil, err
	}
	return &CreatePaymentResponse{PaymentID: id, HateoasLink: HateoasLink{Self: fmt.Sprintf("localhost:8080/v1/payments/%s/", id)}}, nil
}

// UpdatePayment update a payment ressource
func (s service) UpdatePayment(req UpdatePaymentRequest) (*UpdatePaymentResponse, error) {
	// udpate payment
	err := s.repository.UpdatePayment(req.PaymentID, req.Payment)
	if err != nil {
		return nil, err
	}
	return &UpdatePaymentResponse{PaymentID: req.PaymentID}, nil
}

// DeletePayment deletes a given payment by ID
func (s service) DeletePayment(req DeletePaymentRequest) (*DeletePaymentResponse, error) {
	// delete a payment
	err := s.repository.DeletePayment(req.PaymentID)
	if err != nil {
		return nil, err
	}

	return &DeletePaymentResponse{PaymentID: req.PaymentID}, err
}
