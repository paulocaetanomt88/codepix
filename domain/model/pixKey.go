package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

// interface que normatiza como poderá ser implementada por outras aplicações e ter acesso ao banco de dados
type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

type PixKey struct {
	Base 	`valid:"required"`
	Kind 	string 	`json:"kind" valid:"notnull"`
	Key 	string 	`json:"key" valid:"notnull"`
	AccountID string `json:"accountID" valid:"notnull"`
	Account *Account `valid:"-"`
	Status 	string 	`json:"status" valid:"notnull"`
}

// método de validação da chave Pix fictícia
func (pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixKey)

	// validações específicas de regras de negócio para Kind
	if pixKey.Kind != "email" && pixKey.Kind != "cpf" {
		return errors.New("invalid type of key")
	}
	if pixKey.Status != "active" && pixKey.Status != "inactive" {
		return errors.New("invalid type of key")
	}

	if err != nil {
		return err
	}

	return nil
}

// função para criar nova chave pix fictícia
func NewPixKey(kind string, account *Account, key string) (*PixKey, error) {
	pixKey := PixKey{
		Kind: kind,
		Key: key,
		Account: account,
		Status: "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}

	return &pixKey, nil
}
