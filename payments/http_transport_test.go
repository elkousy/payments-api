package payments

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_MakeHTTPHandler(t *testing.T) {
	router := mux.NewRouter()
	h := MakeHTTPHandler(Endpoints{}, router)
	assert.NotNil(t, h)
}

func Test_decodePostPaymentRequest(t *testing.T) {
	//Arrange
	expected := CreatePaymentRequest{}
	r := httptest.NewRequest("POST", "/v1/payments/", bytes.NewBufferString("{}"))
	//Act
	req, err := decodePostPaymentRequest(context.Background(), r)
	//Assert
	tt := assert.New(t)
	tt.Nil(err)
	tt.Equal(expected, req.(CreatePaymentRequest))
}

func Test_decodeGetListOfPaymentsRequest(t *testing.T) {
	//Arrange
	expected := GetListOfPaymentsRequest{}
	r := httptest.NewRequest("GET", "/v1/payments/", nil)
	//Act
	req, err := decodeGetListOfPaymentsRequest(context.Background(), r)
	//Assert
	tt := assert.New(t)
	tt.Nil(err)
	tt.Equal(expected, req.(GetListOfPaymentsRequest))
}

func Test_decodeGetPaymentRequest(t *testing.T) {
	//Arrange
	expectedResult := GetPaymentRequest{
		PaymentID: "abcd",
	}
	httpRequest, _ := http.NewRequest("GET", "/v1/payments/abcd/", nil)
	httpRequest = mux.SetURLVars(httpRequest, map[string]string{"id": "abcd"})
	//Act
	req, err := decodeGetPaymentRequest(context.Background(), httpRequest)
	//Assert
	require.NoError(t, err)
	require.Equal(t, expectedResult, req)
}

func Test_decodeUpdatePaymentRequest(t *testing.T) {
	//Arrange
	expectedResult := UpdatePaymentRequest{
		PaymentID: "abcd",
	}
	httpRequest, err := http.NewRequest("PUT", "/v1/payments/abcd/", bytes.NewBufferString("{}"))
	httpRequest = mux.SetURLVars(httpRequest, map[string]string{"id": "abcd"})
	//Act
	req, err := decodeUpdatePaymentRequest(context.Background(), httpRequest)
	//Assert
	require.NoError(t, err)
	require.Equal(t, expectedResult, req)
}
func Test_decodeDeletePaymentRequest(t *testing.T) {
	//Arrange
	expectedResult := DeletePaymentRequest{
		PaymentID: "abcd",
	}
	httpRequest, err := http.NewRequest("DELETE", "/v1/payments/abcd/", nil)
	httpRequest = mux.SetURLVars(httpRequest, map[string]string{"id": "abcd"})
	//Act
	req, err := decodeDeletePaymentRequest(context.Background(), httpRequest)
	//Assert
	require.NoError(t, err)
	require.Equal(t, expectedResult, req)
}
func Test_encodeOKResponse(t *testing.T) {
	// Arrange
	rr := httptest.NewRecorder()
	// Act
	err := encodeOKResponse(context.Background(), rr, CreatePaymentResponse{})
	//Assert
	assert.Equal(t, rr.Header().Get("Content-Type"), "application/json; charset=utf-8")
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Nil(t, err)
}
func Test_encodeAcceptedResponse(t *testing.T) {
	// Arrange
	rr := httptest.NewRecorder()
	// Act
	err := encodeAcceptedResponse(context.Background(), rr, DeletePaymentResponse{})
	//Assert
	assert.Equal(t, rr.Code, http.StatusAccepted)
	assert.Nil(t, err)
}
func Test_encodeCreatedResponse(t *testing.T) {
	// Arrange
	rr := httptest.NewRecorder()
	// Act
	err := encodeCreatedResponse(context.Background(), rr, CreatePaymentResponse{})
	//Assert
	assert.Equal(t, rr.Code, http.StatusCreated)
	assert.Nil(t, err)
}
