package payments

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_validatorService_GetPayment(t *testing.T) {
	type args struct {
		req GetPaymentRequest
	}
	type serviceResult struct {
		res *GetPaymentResponse
		err error
	}
	tests := []struct {
		name              string
		args              args
		mockServiceResult *serviceResult
		want              *GetPaymentResponse
		wantErr           error
	}{
		{
			name: "Should return error invalid payment id when it is not a valid uuid",
			args: args{
				req: GetPaymentRequest{
					PaymentID: "1",
				},
			},
			wantErr: ErrInvalidPaymentID,
		},
		{
			name: "Should return a successful get response",
			args: args{
				req: GetPaymentRequest{
					PaymentID: "7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3",
				},
			},
			mockServiceResult: &serviceResult{
				res: &GetPaymentResponse{},
				err: nil,
			},
			want: &GetPaymentResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockService := &MockService{}
			if tt.mockServiceResult != nil {
				mockService.On("GetPayment", tt.args.req).Return(tt.mockServiceResult.res, tt.mockServiceResult.err)
			}
			s, _ := newValidator(mockService)
			// Act & Assert
			got, err := s.GetPayment(tt.args.req)
			if err != tt.wantErr {
				t.Errorf("validatorService.GetPayment() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validatorervice.GetPayment() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_validatorService_UpdatePayment(t *testing.T) {
	type args struct {
		req UpdatePaymentRequest
	}
	type serviceResult struct {
		res *UpdatePaymentResponse
		err error
	}
	tests := []struct {
		name              string
		args              args
		mockServiceResult *serviceResult
		want              *UpdatePaymentResponse
		wantErr           error
	}{
		{
			name: "Should return error invalid payment id when it is not a valid uuid",
			args: args{
				req: UpdatePaymentRequest{
					PaymentID: "1",
					Payment:   mockNewPayment("7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3"),
				},
			},
			wantErr: ErrInvalidPaymentID,
		},
		{
			name: "Should return invalid payload when required fields are missing",
			args: args{
				req: UpdatePaymentRequest{
					PaymentID: "7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3",
					Payment:   mockNewPaymentMissingFields("7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3"),
				},
			},
			wantErr: ErrInvalidPaymentPayload,
		},
		{
			name: "Should return a successful update response",
			args: args{
				req: UpdatePaymentRequest{
					PaymentID: "7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3",
					Payment:   mockNewPayment("7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3"),
				},
			},
			mockServiceResult: &serviceResult{
				res: &UpdatePaymentResponse{PaymentID: "7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3"},
				err: nil,
			},
			want: &UpdatePaymentResponse{PaymentID: "7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockService := &MockService{}
			if tt.mockServiceResult != nil {
				mockService.On("UpdatePayment", tt.args.req).Return(tt.mockServiceResult.res, tt.mockServiceResult.err)
			}
			s, _ := newValidator(mockService)
			// Act & Assert
			got, err := s.UpdatePayment(tt.args.req)
			if err != tt.wantErr {
				t.Errorf("validatorService.UpdatePayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validatorervice.UpdatePayment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validatorService_DeletePayment(t *testing.T) {
	type args struct {
		req DeletePaymentRequest
	}
	type serviceResult struct {
		deletePaymentResponse *DeletePaymentResponse
		err                   error
	}
	tests := []struct {
		name              string
		args              args
		mockServiceResult *serviceResult
		want              *DeletePaymentResponse
		wantErr           error
	}{
		{
			name: "Should return error invalid payment id when it is not a valid uuid",
			args: args{
				req: DeletePaymentRequest{
					PaymentID: "1",
				},
			},
			wantErr: ErrInvalidPaymentID,
		},
		{
			name: "Should return delete payment response",
			args: args{
				req: DeletePaymentRequest{
					PaymentID: "7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3",
				},
			},
			mockServiceResult: &serviceResult{
				deletePaymentResponse: &DeletePaymentResponse{PaymentID: "7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3"},
				err:                   nil,
			},
			want: &DeletePaymentResponse{PaymentID: "7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockService := &MockService{}
			if tt.mockServiceResult != nil {
				mockService.On("DeletePayment", tt.args.req).Return(tt.mockServiceResult.deletePaymentResponse, tt.mockServiceResult.err)
			}
			s, _ := newValidator(mockService)
			// Act & Assert
			got, err := s.DeletePayment(tt.args.req)
			if err != tt.wantErr {
				t.Errorf("validatorService.DeletePayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validatorervice.DeletePayment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validatePaymentID_OK(t *testing.T) {
	//Arrange
	id := "7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3"
	//Act
	err := validatePaymentID(id)
	//Assert
	require.NoError(t, err)
}
func Test_validatePaymentID_Fail(t *testing.T) {
	//Arrange
	id := "7c95bd23b67f4cc9bfb29e4e31f093e"
	//Act
	err := validatePaymentID(id)
	//Assert
	require.Error(t, err)
}
