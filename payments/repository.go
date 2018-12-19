package payments

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/elkousy/payments-api/utility/config"
	_ "github.com/lib/pq" //pq imports the postgres driver
	uuid "github.com/satori/go.uuid"
)

// Repository describes a payments repository used to manipulate payments data
type Repository interface {
	GetPayment(id string) (Payment, error)
	GetListOfPayments() ([]Payment, error)
	CreatePayment(p Payment) (string, error)
	UpdatePayment(id string, p Payment) error
	DeletePayment(id string) error
}

const connectionString = "host=%s port=%d dbname=%s user=%s password=%s sslmode=disable connect_timeout=%d application_name=%s"

type paymentRepository struct {
	db *gorm.DB
}

// NewPaymentRepository ...
func NewPaymentRepository(db *gorm.DB) Repository {
	return &paymentRepository{
		db: db,
	}
}

// DbConnect connects to the db
func DbConnect() (*gorm.DB, error) {
	cnx := newConnection(config.DBHost, config.DBPort, config.DBName, config.DBUser, config.DBPassword, config.DBTimeout, config.AppName)
	db, err := gorm.Open("postgres", cnx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func newConnection(host string, port int, name string, user string, password string, timeout int, app string) string {
	return fmt.Sprintf(connectionString, host, port, name, user, password, timeout, app)
}

// DbMigrate initializes db schema with needed tables
func DbMigrate(db *gorm.DB) {
	//db.DropTableIfExists(&Payment{}, &Attributes{}, &BeneficiaryParty{}, &DebtorParty{}, &SponsorParty{}, &ChargesInformation{}, &Charge{}, &Forex{})
	db.CreateTable(&Payment{}, &Attributes{}, &BeneficiaryParty{}, &DebtorParty{}, &SponsorParty{}, &ChargesInformation{}, &Charge{}, &Forex{})
}

// DbClose closes the connection to the database
func DbClose(db *gorm.DB) {
	if db != nil {
		db.Close()
	}
}

// GetPaymentByID ...
func (r *paymentRepository) GetPayment(id string) (Payment, error) {
	p := Payment{}
	err := r.db.Debug().Model(&p).Where("id = ?", id).Preload("Attributes.BeneficiaryParty").Preload("Attributes.ChargesInformation.SenderCharges").Preload("Attributes.DebtorParty").Preload("Attributes.Forex").Preload("Attributes.SponsorParty").Find(&p).Error
	if err != nil {
		return p, ErrNotFound.FromError(err)
	}
	return p, nil
}

// CreatePayment ...
func (r *paymentRepository) CreatePayment(p Payment) (string, error) {
	paymentID := uuid.NewV4()
	p.ID = paymentID
	err := r.db.Debug().Save(&p).Error
	if err != nil {
		return "", err
	}
	return paymentID.String(), nil
}

// UpdatePayment ...
func (r *paymentRepository) UpdatePayment(id string, p Payment) error {
	pid, err := uuid.FromString(id)
	if err != nil {
		return err
	}
	p.ID = pid
	pa := &Payment{}
	if err := r.db.Debug().Model(&p).Where("id = ?", p.ID).Find(&pa).Error; err != nil {
		return ErrNotFound.FromError(err)
	}

	err = r.db.Debug().Model(&p).Save(&p).Error
	if err != nil {
		return err
	}
	return nil
}

// DeletePayment ...
func (r *paymentRepository) DeletePayment(id string) error {
	pa := &Payment{}
	if err := r.db.Debug().Model(pa).Where("id = ?", id).Find(pa).Error; err != nil {
		return ErrNotFound.FromError(err)
	}
	// Delete payment by ID `Soft Delete`
	if err := r.db.Debug().Model(pa).Where("id = ?", id).Delete(pa).Error; err != nil {
		return err
	}
	return nil
}

// GetListOfPayments ...
func (r *paymentRepository) GetListOfPayments() ([]Payment, error) {
	var payments []Payment
	err := r.db.Debug().Find(&payments).Error
	if err != nil {
		return nil, err
	}
	for i, p := range payments {
		payments[i], _ = r.GetPayment(p.ID.String())
	}
	return payments, nil
}
