package payments

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Service_GetPayment(t *testing.T) {
	// Arrange
	id := "7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3"
	p := mockNewPayment(id)
	expectedRes := GetPaymentResponse{Payment: p}
	repositoryMock := &MockRepository{}
	repositoryMock.On("GetPayment", mock.Anything).Return(p, nil)
	service, _ := NewPaymentService(repositoryMock)

	//Act
	res, err := service.GetPayment(GetPaymentRequest{PaymentID: id})

	//Assert
	assert.NoError(t, err)
	assert.NotNil(t, res, "result should not be nil")
	assert.Equal(t, expectedRes, *res)
}

func Test_Service_GetListOfPayments(t *testing.T) {
	// Arrange

	id1 := "7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3"
	id2 := "6ef6057f-0ed4-48c9-a128-f85b8f024519"
	pays := []Payment{mockNewPayment(id1), mockNewPayment(id2)}
	expectedRes := GetListOfPaymentsResponse{Data: pays, HateoasLink: HateoasLink{Self: "localhost:8080/v1/payments/"}}
	repositoryMock := &MockRepository{}
	repositoryMock.On("GetListOfPayments", mock.Anything).Return(pays, nil)
	service, _ := NewPaymentService(repositoryMock)

	//Act
	res, err := service.GetListOfPayments(GetListOfPaymentsRequest{})

	//Assert
	assert.NoError(t, err)
	assert.NotNil(t, res, "result should not be nil")
	assert.Equal(t, expectedRes, *res)
}

func Test_Service_PostPayment(t *testing.T) {
	// Arrange
	id := "7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3"
	p := mockNewPayment(id)
	expectedRes := CreatePaymentResponse{PaymentID: id, HateoasLink: HateoasLink{Self: fmt.Sprintf("localhost:8080/v1/payments/%s/", id)}}
	repositoryMock := &MockRepository{}
	repositoryMock.On("CreatePayment", mock.Anything).Return(id, nil)
	service, _ := NewPaymentService(repositoryMock)

	//Act
	res, err := service.PostPayment(CreatePaymentRequest{Payment: p})

	//Assert
	assert.NoError(t, err)
	assert.NotNil(t, res, "result should not be nil")
	assert.Equal(t, expectedRes, *res)
}

func Test_Service_UpdatePayment(t *testing.T) {
	// Arrange
	id := "7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3"
	p := mockNewPayment(id)
	expectedRes := UpdatePaymentResponse{PaymentID: id}
	repositoryMock := &MockRepository{}
	repositoryMock.On("UpdatePayment", mock.Anything, mock.Anything).Return(nil)
	service, _ := NewPaymentService(repositoryMock)

	//Act
	res, err := service.UpdatePayment(UpdatePaymentRequest{Payment: p, PaymentID: id})

	//Assert
	assert.NoError(t, err)
	assert.NotNil(t, res, "result should not be nil")
	assert.Equal(t, expectedRes, *res)
}
func Test_Service_DeletePayment(t *testing.T) {
	// Arrange
	id := "7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3"
	expectedRes := DeletePaymentResponse{PaymentID: id}
	repositoryMock := &MockRepository{}
	repositoryMock.On("DeletePayment", mock.Anything).Return(nil)
	service, _ := NewPaymentService(repositoryMock)

	//Act
	res, err := service.DeletePayment(DeletePaymentRequest{PaymentID: id})

	//Assert
	assert.NoError(t, err)
	assert.NotNil(t, res, "result should not be nil")
	assert.Equal(t, expectedRes, *res)
}
