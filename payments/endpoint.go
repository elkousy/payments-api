package payments

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
)

// Endpoints contains all go-kit like endpoints used to manipulate payments
type Endpoints struct {
	GetPayment        endpoint.Endpoint
	GetListOfPayments endpoint.Endpoint
	PostPayment       endpoint.Endpoint
	UpdatePayment     endpoint.Endpoint
	DeletePayment     endpoint.Endpoint
}

//MakeEndpoints ...
func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		GetPayment:        makeGetPaymentEndpoint(svc),
		GetListOfPayments: makeGetListOfPaymentsEndpoint(svc),
		PostPayment:       makePostPaymentEndpoint(svc),
		UpdatePayment:     makeUpdatePaymentEndpoint(svc),
		DeletePayment:     makeDeletePaymentEndpoint(svc),
	}
}

// makeGetPaymentEndpoint creates a go-kit like endpoint
// used to retrieve specific payment by ID
func makeGetPaymentEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var r GetPaymentRequest
		var ok bool

		if r, ok = request.(GetPaymentRequest); !ok {
			return nil, errors.New("failed to cast GetPaymentRequest")
		}
		return svc.GetPayment(r)
	}
}

// makeGetListOfPaymentsEndpoint creates a go-kit like endpoint used to get list of payments
func makeGetListOfPaymentsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r := GetListOfPaymentsRequest{}
		return svc.GetListOfPayments(r)
	}
}

// makePostPaymentEndpoint creates a go-kit like endpoint used to post payments
func makePostPaymentEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var r CreatePaymentRequest
		var ok bool

		if r, ok = request.(CreatePaymentRequest); !ok {
			return nil, errors.New("failed to cast CreatePaymentRequest")
		}

		return svc.PostPayment(r)
	}
}

// makeUpdatePaymentEndpoint creates a go-kit like endpoint used to update a payment by ID
func makeUpdatePaymentEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var r UpdatePaymentRequest
		var ok bool

		if r, ok = request.(UpdatePaymentRequest); !ok {
			return nil, errors.New("failed to cast UpdatePaymentRequest")
		}

		return svc.UpdatePayment(r)
	}
}

// makeDeletePaymentEndpoint creates a go-kit like endpoint used to delete a payment by ID
func makeDeletePaymentEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var r DeletePaymentRequest
		var ok bool

		if r, ok = request.(DeletePaymentRequest); !ok {
			return nil, errors.New("failed to cast DeletePaymentRequest")
		}

		return svc.DeletePayment(r)
	}
}
