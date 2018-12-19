package payments

import (
	"log"
	"testing"

	mocket "github.com/Selvatico/go-mocket"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_newConnection(t *testing.T){
	// Arrange
	expectedRes := "host=localhost port=5432 dbname=postgres user=postgres password=postgres sslmode=disable connect_timeout=5 application_name=payments"
	// Act
	cnx := newConnection("localhost", 5432, "postgres", "postgres", "postgres",5,"payments")
	// Assert
	assert.Equal(t, expectedRes, cnx)
}

func SetupDBTests() *gorm.DB {
	mocket.Catcher.Register()

	db, err := gorm.Open(mocket.DriverName, "connection_string")
	if err != nil {
		log.Fatalf("error mocking gorm: %s", err)
	}
	// Log mode shows the query gorm uses, so we can replicate and mock it
	db.LogMode(true)

	return db
}

func Test_CreatePayment(t *testing.T) {
	//Arrange
	idStr := "7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3"
	pid, _ := uuid.FromString(idStr)
	p := Payment{
		ID: pid,
	}

	db := SetupDBTests()
	defer db.Close()

	mocket.Catcher.Reset().NewMock().WithQuery("INSERT INTO \"payments\"")
	r := NewPaymentRepository(db)

	//Act
	id, err := r.CreatePayment(p)

	//Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
}

func Test_UpdatePayment(t *testing.T) {
	//Arrange
	idStr := "7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3"
	pid, _ := uuid.FromString(idStr)
	mockReply := []map[string]interface{}{{}}
	p := Payment{
		ID: pid,
	}

	db := SetupDBTests()
	defer db.Close()

	//mocket.Catcher.Reset().NewMock().WithQuery("SELECT * FROM \"payments\"").WithReply(mockReply)
	//mocket.Catcher.Reset().NewMock().WithQuery("UPDATE \"payments\"")

	mocket.Catcher.Attach([]*mocket.FakeResponse{
		{
			Pattern:  "SELECT * FROM \"payments\"",
			Response: mockReply,
		},
	})
	mocket.Catcher.Attach([]*mocket.FakeResponse{
		{
			Pattern:  "UPDATE \"payments\"",
			Response: mockReply,
		},
	})

	r := NewPaymentRepository(db)

	//Act
	err := r.UpdatePayment(idStr, p)

	//Assert
	assert.NoError(t, err)
}

func Test_DeletePayment(t *testing.T) {
	//Arrange
	idStr := "7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3"
	mockReply := []map[string]interface{}{{}}

	db := SetupDBTests()
	defer db.Close()

	mocket.Catcher.Attach([]*mocket.FakeResponse{
		{
			Pattern:  "SELECT * FROM \"payments\"",
			Response: mockReply,
		},
	})
	mocket.Catcher.Attach([]*mocket.FakeResponse{
		{
			Pattern:  "DELETE \"payments\"",
			Response: mockReply,
		},
	})

	r := NewPaymentRepository(db)

	//Act
	err := r.DeletePayment(idStr)

	//Assert
	assert.NoError(t, err)
}

func Test_GetPayment(t *testing.T) {
	//Arrange
	idStr := "7c95bd23-b67f-4cc9-bfb2-9e4e31f093e3"
	mockReply := []map[string]interface{}{{
		"created_at":      "test",
		"updated_at":      "test",
		"deleted_at":      "test",
		"id":              "test",
		"type":            "test",
		"version":         "test",
		"organisation_id": "test",
		"attributes_id":   "test",
	}}

	db := SetupDBTests()
	defer db.Close()

	mocket.Catcher.Attach([]*mocket.FakeResponse{
		{
			Pattern:  "SELECT * FROM \"payments\"",
			Response: mockReply,
		},
	})

	r := NewPaymentRepository(db)

	//Act
	p, err := r.GetPayment(idStr)

	//Assert
	assert.NotNil(t, p)
	assert.NoError(t, err)
}

func Test_GetListOfPayments(t *testing.T) {
	//Arrange
	mockReply := []map[string]interface{}{
		{
			"created_at":      "test",
			"updated_at":      "test",
			"deleted_at":      "test",
			"id":              "test",
			"type":            "test",
			"version":         "test",
			"organisation_id": "test",
			"attributes_id":   "test",
		},
		{
			"created_at":      "test",
			"updated_at":      "test",
			"deleted_at":      "test",
			"id":              "test",
			"type":            "test",
			"version":         "test",
			"organisation_id": "test",
			"attributes_id":   "test",
		},
	}

	db := SetupDBTests()
	defer db.Close()

	mocket.Catcher.Attach([]*mocket.FakeResponse{
		{
			Pattern:  "SELECT * FROM \"payments\"",
			Response: mockReply,
		},
	})

	r := NewPaymentRepository(db)

	//Act
	p, err := r.GetListOfPayments()
	println(len(p))
	//Assert
	assert.NotNil(t, p)
	assert.NoError(t, err)
}

func mockNewPayment(id string) Payment {
	pID, _ := uuid.FromString(id)
	p := Payment{
		Type:           "Payment",
		Version:        0,
		ID:             pID,
		OrganisationID: uuid.NewV4(),
		Attributes: Attributes{
			Amount: "100.21",
			BeneficiaryParty: BeneficiaryParty{
				AccountType: 0,
				DebtorParty: DebtorParty{
					AccountName:       "aqssbb",
					AccountNumberCode: "IBAN",
					Address:           "34 frfrf ded",
					Name:              "ING Dfh",
					SponsorParty: SponsorParty{
						AccountNumber: "5678923",
						BankID:        "134667",
						BankIDCode:    "GSDFE",
					},
				},
			},
			ChargesInformation: ChargesInformation{
				BearerCode: "SHAR",
				SenderCharges: []Charge{
					Charge{Amount: "5.00", Currency: "GBP"},
					Charge{Amount: "10.00", Currency: "USD"},
				},
				ReceiverChargesAmount:   "1.00",
				ReceiverChargesCurrency: "USD",
			},
			Currency: "GBP",
			DebtorParty: DebtorParty{
				AccountName:       "deded",
				AccountNumberCode: "IBAN",
				Address:           "1 dhhde ded",
				Name:              "alspnfh",
				SponsorParty: SponsorParty{
					AccountNumber: "5678923",
					BankID:        "134667",
					BankIDCode:    "GSDFE",
				},
			},
			EndToEndReference: "Wil def ee",
			Forex: Forex{
				ContractReference: "FX123",
				ExchangeRate:      "2.0000",
				OriginalAmount:    "200.42",
				OriginalCurrency:  "USD",
			},
			NumericReference:     "10223453",
			PayID:                "123344556790",
			PaymentPurpose:       "course",
			PaymentScheme:        "FPS",
			PaymentType:          "Credit",
			ProcessingDate:      "2017-01-18",
			Reference:            "PAYmen",
			SchemePaymentSubType: "InternetBanking",
			SchemePaymentType:    "Immediate Pay",
			SponsorParty: SponsorParty{
				AccountNumber: "5678923",
				BankID:        "134667",
				BankIDCode:    "GSDFE",
			},
		},
	}

	return p
}

func mockNewPaymentMissingFields(id string) Payment {
	pID, _ := uuid.FromString(id)
	p := Payment{
		Type:           "Payment",
		Version:        0,
		ID:             pID,
		OrganisationID: uuid.NewV4(),
		Attributes: Attributes{
			Amount: "100.21",
			ChargesInformation: ChargesInformation{
				BearerCode: "SHAR",
				SenderCharges: []Charge{
					Charge{Amount: "5.00", Currency: "GBP"},
					Charge{Amount: "10.00", Currency: "USD"},
				},
				ReceiverChargesAmount:   "1.00",
				ReceiverChargesCurrency: "USD",
			},
			Currency: "GBP",
			DebtorParty: DebtorParty{
				AccountName:       "deded",
				AccountNumberCode: "IBAN",
				Address:           "1 dhhde ded",
				Name:              "alspnfh",
				SponsorParty: SponsorParty{
					AccountNumber: "5678923",
					BankID:        "134667",
					BankIDCode:    "GSDFE",
				},
			},
			EndToEndReference: "Wil def ee",
			Forex: Forex{
				ContractReference: "FX123",
				ExchangeRate:      "2.0000",
				OriginalAmount:    "200.42",
				OriginalCurrency:  "USD",
			},
			NumericReference:     "10223453",
			PayID:                "123344556790",
			PaymentPurpose:       "course",
			PaymentScheme:        "FPS",
			PaymentType:          "Credit",
			ProcessingDate:       "2017-01-18",
			Reference:            "PAYmen",
			SchemePaymentSubType: "InternetBanking",
			SchemePaymentType:    "Immediate Pay",
			SponsorParty: SponsorParty{
				AccountNumber: "5678923",
				BankID:        "134667",
				BankIDCode:    "GSDFE",
			},
		},
	}

	return p
}
