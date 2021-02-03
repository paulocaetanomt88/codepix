package model

import (
	"github.com/asaskevich/govalidator"
	"time"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Base struct {
	ID string `json:"id" valid:"uuid"` // Go trabalha com tags que ajuda na serialização, que é uma forma mais fácil de converter para json
	CreatedAt time.Time `json:"createdAt" valid:"-"`
	UpdatedAt time.Time `json:"updatedAt" valid:"-"`
}