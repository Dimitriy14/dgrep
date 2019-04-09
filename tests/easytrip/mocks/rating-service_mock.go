package mocks

import (
	"github.com/oreuta/easytrip/models"
)

type RatesServiceMock struct {
	Banks []models.CurrencyBank
	Err   error
}

func (m RatesServiceMock) GetBankRates(r models.MainRequest) (banks []models.CurrencyBank, err error) {
	return m.Banks, m.Err
}
