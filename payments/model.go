package payments

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// ModelBase model definition, including fields `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in all models
type ModelBase struct {
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

type Model struct {
	ID uint `json:"-" gorm:"primary_key"`
	ModelBase
}

// Payment reprensents a payment resource
type Payment struct {
	ModelBase
	ID             uuid.UUID  `json:"id" gorm:"type:uuid; primary_key"`
	Type           string     `json:"type" validate:"required"`
	Version        uint       `json:"version" binding:"exists"`
	OrganisationID uuid.UUID  `json:"organisation_id" validate:"required"`
	Attributes     Attributes `json:"attributes" gorm:"auto_preload" validate:"required"`
	AttributesID   uint       `json:"-" sql:"index"`
}

// Attributes ...
type Attributes struct {
	Model
	Amount               string             `json:"amount" validate:"required"`
	BeneficiaryParty     BeneficiaryParty   `json:"beneficiary_party" gorm:"auto_preload" validate:"required"`
	BeneficiaryPartyID   uint               `json:"-" sql:"index"`
	ChargesInformation   ChargesInformation `json:"charges_information" gorm:"auto_preload" validate:"required"`
	ChargesInformationID uint               `json:"-" sql:"index"`
	Currency             string             `json:"currency" validate:"required"`
	DebtorParty          DebtorParty        `json:"debtor_party" gorm:"auto_preload" validate:"required"`
	DebtorPartyID        uint               `json:"-" sql:"index"`
	EndToEndReference    string             `json:"end_to_end_reference" validate:"required"`
	Forex                Forex              `json:"fx" gorm:"auto_preload" validate:"required"`
	ForexID              uint               `json:"-" sql:"index"`
	NumericReference     string             `json:"numeric_reference" validate:"required"`
	PayID                string             `json:"payment_id" validate:"required"`
	PaymentPurpose       string             `json:"payment_purpose" validate:"required"`
	PaymentScheme        string             `json:"payment_scheme" validate:"required"`
	PaymentType          string             `json:"payment_type" validate:"required"`
	ProcessingDate       string             `json:"processing_date" validate:"required"`
	Reference            string             `json:"reference" validate:"required"`
	SchemePaymentSubType string             `json:"scheme_payment_sub_type" validate:"required"`
	SchemePaymentType    string             `json:"scheme_payment_type" validate:"required"`
	SponsorParty         SponsorParty       `json:"sponsor_party" gorm:"auto_preload" validate:"required"`
	SponsorPartyID       uint               `json:"-" sql:"index"`
}

// BeneficiaryParty ...
type BeneficiaryParty struct {
	DebtorParty
	AccountType int `json:"account_type" binding:"exists"`
}

// DebtorParty ...
type DebtorParty struct {
	SponsorParty
	AccountName       string `json:"account_name" validate:"required"`
	AccountNumberCode string `json:"account_number_code" validate:"required"`
	Address           string `json:"address" validate:"required"`
	Name              string `json:"name" validate:"required"`
}

// SponsorParty ...
type SponsorParty struct {
	Model
	AccountNumber string `json:"account_number" validate:"required"`
	BankID        string `json:"bank_id" validate:"required"`
	BankIDCode    string `json:"bank_id_code" validate:"required"`
}

// ChargesInformation ...
type ChargesInformation struct {
	Model
	BearerCode              string   `json:"bearer_code" validate:"required"`
	SenderCharges           []Charge `json:"sender_charges" gorm:"auto_preload" validate:"required"`
	ReceiverChargesAmount   string   `json:"receiver_charges_amount" validate:"required"`
	ReceiverChargesCurrency string   `json:"receiver_charges_currency" validate:"required"`
}

// Charge ...
type Charge struct {
	Model
	ChargesInformationID uint   `json:"-" sql:"index"`
	Amount               string `json:"amount" validate:"required"`
	Currency             string `json:"currency" validate:"required"`
}

// Forex ...
type Forex struct {
	Model
	ContractReference string `json:"contract_reference" validate:"required"`
	ExchangeRate      string `json:"exchange_rate" validate:"required"`
	OriginalAmount    string `json:"original_amount" validate:"required"`
	OriginalCurrency  string `json:"original_currency" validate:"required"`
}

/************************/

// GetPaymentRequest is the request parameter used to retrieve a specific payment
type GetPaymentRequest struct {
	PaymentID string
}

// GetPaymentResponse is the response object returned by the get payment endpoint.
type GetPaymentResponse struct {
	Payment
}

// GetListOfPaymentsRequest is the request parameter used to retrieve a specific payment
type GetListOfPaymentsRequest struct{}

// GetListOfPaymentsResponse is the response object returned by the get payment endpoint.
// Data enveloped, a top level object is secure and succinctif you do not envelope JSON arrays.
type GetListOfPaymentsResponse struct {
	Data        []Payment `json:"data"`
	HateoasLink `json:"links"`
}

// CreatePaymentRequest represents the request parameters used for inserting a new payment
type CreatePaymentRequest struct {
	Payment
}

// CreatePaymentResponse represents the response returned after inserting a new payment
type CreatePaymentResponse struct {
	PaymentID   string `json:"id"`
	HateoasLink `json:"links"`
}

// UpdatePaymentRequest is the request object passed to the update payment endpoint.
type UpdatePaymentRequest struct {
	PaymentID string
	Payment
}

// UpdatePaymentResponse is the response object returned by the update payment endpoint.
type UpdatePaymentResponse struct {
	PaymentID string `json:"id"`
}

// DeletePaymentRequest represents the request parameter needed to delete a payment
type DeletePaymentRequest struct {
	PaymentID string
}

// DeletePaymentResponse represents the response sent when a payment is deleted
type DeletePaymentResponse struct {
	PaymentID string `json:"id"`
}

// HateoasLink represents the HATEOS links along with the response
type HateoasLink struct {
	Self string `json:"self"`
}
