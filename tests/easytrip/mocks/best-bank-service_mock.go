package mocks

import (
	"github.com/oreuta/easytrip/models"
)

type BestServiceMock struct {
	BBSale []models.CurrencyBank
	BBBuy  []models.CurrencyBank
	Err    error
}

func (m BestServiceMock) GetBestBanks(r models.MainRequest) (bBSale []models.CurrencyBank, bBBuy []models.CurrencyBank, err error) {
	return m.BBSale, m.BBBuy, m.Err
}
