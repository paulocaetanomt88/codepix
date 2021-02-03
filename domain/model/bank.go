package model

import (
	"github.com/asaskevich/govalidator"
	"time"
	uuid "github.com/satori/go.uuid"
)


// É possível programar orientado a objetos, mas não da forma mais comum,
// pois Go não utiliza classes e sim estruturas
type Bank struct {
	Base `valid:"required"`
	Code string `json:"code" valid:"notnull"`
	Name string `json:"name" valid:"notnull"`
	Accounts []*Account `valid:"-"`
}

// método de validação do banco
func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)

	if err != nil {
		return err
	}

	return nil
}

// função para criar novo banco
func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid()
	if err != nil {
		return nil, err
	}

	return &bank, nil
}
