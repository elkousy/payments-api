package payments

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_MakeEndpoints(t *testing.T) {
	//Arrange
	mockService := &MockService{}
	//Act
	e := MakeEndpoints(mockService)
	//Assert
	assert.NotNil(t, e)
}

func Test_makeDeletePaymentEndpoint(t *testing.T) {
	// Arrange
	res := DeletePaymentResponse{PaymentID: "abcd"}
	testError := errors.New("test error")

	tests := []struct {
		name             string
		Service          func() Service
		isError          bool
		ExpectedError    string
		ExpectedResponse *DeletePaymentResponse
	}{
		{
			name: "makeDeletePaymentEndpoint successfully deleted a payment",
			Service: func() Service {
				mockService := &MockService{}
				mockService.On("DeletePayment", mock.Anything).Return(&res, nil)
				return mockService
			},
			isError:          false,
			ExpectedResponse: &res,
			ExpectedError:    "",
		},
		{
			name: "makeDeletePaymentEndpoint failed to delete a payment",
			Service: func() Service {
				mockService := &MockService{}
				mockService.On("DeletePayment", mock.Anything).Return(nil, testError)
				return mockService
			},
			isError:          true,
			ExpectedResponse: nil,
			ExpectedError:    "test error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := makeDeletePaymentEndpoint(tt.Service())

			//Act
			wl, err := got(nil, DeletePaymentRequest{})

			//Assert
			if tt.isError {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), tt.ExpectedError)
				assert.Nil(t, wl)
				return
			}
			assert.Nil(t, err)
			result, ok := wl.(*DeletePaymentResponse)
			assert.Equal(t, true, ok)
			assert.NotNil(t, result)

		})
	}
}

func Test_makeUpdatePaymentEndpoint(t *testing.T) {
	// Arrange
	res := UpdatePaymentResponse{PaymentID: "abcd"}
	testError := errors.New("test error")

	tests := []struct {
		name             string
		Service          func() Service
		isError          bool
		ExpectedError    string
		ExpectedResponse *UpdatePaymentResponse
	}{
		{
			name: "makeUpdatePaymentEndpoint successfully updated a payment",
			Service: func() Service {
				mockService := &MockService{}
				mockService.On("UpdatePayment", mock.Anything).Return(&res, nil)
				return mockService
			},
			isError:          false,
			ExpectedResponse: &res,
			ExpectedError:    "",
		},
		{
			name: "makeUpdatePaymentEndpoint failed to update a payment",
			Service: func() Service {
				mockService := &MockService{}
				mockService.On("UpdatePayment", mock.Anything).Return(nil, testError)
				return mockService
			},
			isError:          true,
			ExpectedResponse: nil,
			ExpectedError:    "test error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := makeUpdatePaymentEndpoint(tt.Service())

			//Act
			wl, err := got(nil, UpdatePaymentRequest{})

			//Assert
			if tt.isError {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), tt.ExpectedError)
				assert.Nil(t, wl)
				return
			}
			assert.Nil(t, err)
			result, ok := wl.(*UpdatePaymentResponse)
			assert.Equal(t, true, ok)
			assert.NotNil(t, result)

		})
	}
}

func Test_makePostPaymentEndpoint(t *testing.T) {
	// Arrange
	res := CreatePaymentResponse{PaymentID: "abcd"}
	testError := errors.New("test error")

	tests := []struct {
		name             string
		Service          func() Service
		isError          bool
		ExpectedError    string
		ExpectedResponse *CreatePaymentResponse
	}{
		{
			name: "makePostPaymentEndpoint successfully created a new payment",
			Service: func() Service {
				mockService := &MockService{}
				mockService.On("PostPayment", mock.Anything).Return(&res, nil)
				return mockService
			},
			isError:          false,
			ExpectedResponse: &res,
			ExpectedError:    "",
		},
		{
			name: "makePostPaymentEndpoint failed to create a new payment",
			Service: func() Service {
				mockService := &MockService{}
				mockService.On("PostPayment", mock.Anything).Return(nil, testError)
				return mockService
			},
			isError:          true,
			ExpectedResponse: nil,
			ExpectedError:    "test error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := makePostPaymentEndpoint(tt.Service())

			//Act
			wl, err := got(nil, CreatePaymentRequest{})

			//Assert
			if tt.isError {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), tt.ExpectedError)
				assert.Nil(t, wl)
				return
			}
			assert.Nil(t, err)
			result, ok := wl.(*CreatePaymentResponse)
			assert.Equal(t, true, ok)
			assert.NotNil(t, result)

		})
	}
}

func Test_makeGetListOfPaymentsEndpoint(t *testing.T) {
	// Arrange
	res := GetListOfPaymentsResponse{}
	testError := errors.New("test error")

	tests := []struct {
		name             string
		Service          func() Service
		isError          bool
		ExpectedError    string
		ExpectedResponse *GetListOfPaymentsResponse
	}{
		{
			name: "makeGetListOfPaymentsEndpoint successfully retrieves all payments",
			Service: func() Service {
				mockService := &MockService{}
				mockService.On("GetListOfPayments", mock.Anything).Return(&res, nil)
				return mockService
			},
			isError:          false,
			ExpectedResponse: &res,
			ExpectedError:    "",
		},
		{
			name: "makeGetListOfPaymentsEndpoint failed to retrieve all payments",
			Service: func() Service {
				mockService := &MockService{}
				mockService.On("GetListOfPayments", mock.Anything).Return(nil, testError)
				return mockService
			},
			isError:          true,
			ExpectedResponse: nil,
			ExpectedError:    "test error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := makeGetListOfPaymentsEndpoint(tt.Service())

			//Act
			wl, err := got(nil, GetListOfPaymentsRequest{})

			//Assert
			if tt.isError {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), tt.ExpectedError)
				assert.Nil(t, wl)
				return
			}
			assert.Nil(t, err)
			result, ok := wl.(*GetListOfPaymentsResponse)
			assert.Equal(t, true, ok)
			assert.NotNil(t, result)

		})
	}
}

func Test_makeGetPaymentEndpoint(t *testing.T) {
	// Arrange
	res := GetPaymentResponse{}
	testError := errors.New("test error")

	tests := []struct {
		name             string
		Service          func() Service
		isError          bool
		ExpectedError    string
		ExpectedResponse *GetPaymentResponse
	}{
		{
			name: "makeGetPaymentEndpoint successfully retrieves a payment",
			Service: func() Service {
				mockService := &MockService{}
				mockService.On("GetPayment", mock.Anything).Return(&res, nil)
				return mockService
			},
			isError:          false,
			ExpectedResponse: &res,
			ExpectedError:    "",
		},
		{
			name: "makeGetPaymentEndpoint failed to retrieve a payment",
			Service: func() Service {
				mockService := &MockService{}
				mockService.On("GetPayment", mock.Anything).Return(nil, testError)
				return mockService
			},
			isError:          true,
			ExpectedResponse: nil,
			ExpectedError:    "test error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := makeGetPaymentEndpoint(tt.Service())

			//Act
			wl, err := got(nil, GetPaymentRequest{})

			//Assert
			if tt.isError {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), tt.ExpectedError)
				assert.Nil(t, wl)
				return
			}
			assert.Nil(t, err)
			result, ok := wl.(*GetPaymentResponse)
			assert.Equal(t, true, ok)
			assert.NotNil(t, result)

		})
	}
}
