package bestBankService

import (
	"fmt"
	"sort"
	"strings"

	"github.com/oreuta/easytrip/repository"

	"github.com/oreuta/easytrip/clients"
	"github.com/oreuta/easytrip/models"
)

var nameOfBanks = map[string]string{
	"privat": "ПриватБанк",
	"pireus": "Піреус Банк",
	"otp":    "ОТП Банк",
	"kredo":  "Кредобанк",
}

var nameOfCurrency = map[string]string{
	"usd": "USD",
	"eur": "EUR",
}
var nameOfOption = map[string]string{
	"sale": "sale",
	"buy":  "buy",
}

//BestBankServiceInterface ..

type BestBankServiceInterface interface {
	GetBestBanks(r models.MainRequest) (bBSale []models.CurrencyBank, bBBuy []models.CurrencyBank, err error)
}

//BestBankService ..
type BestBankService struct {
	Client clients.BankUAClient
}

//New ..
func New(newClient clients.BankUAClient) BestBankServiceInterface {
	return &BestBankService{Client: newClient}
}

func (b BestBankService) GetBestBanks(data models.MainRequest) (bBSale, bBBuy []models.CurrencyBank, err error) {
	banks, err := b.Client.GetCurrBank()
	if err != nil {
		fmt.Errorf("Method Get in Client BankUACient: %v", err)
		banks, _ := repository.JsnChanger()
		_ = banks
	}

	banks = FilterCurrency(data, FilterBank(data, banks))
	if data.Option != nameOfOption["buy"] {
		bBSale = BestSale(banks)
	}
	if data.Option != nameOfOption["sale"] {
		bBBuy = BestBuy(banks)
	}
	return bBSale, bBBuy, err
}

func FilterBank(data models.MainRequest, inpBanks []models.CurrencyBank) (OutpBanks []models.CurrencyBank) {
	s := strings.Join(data.Bank, "")

	for key := range nameOfBanks {
		if strings.Contains(s, key) {
			for _, value := range inpBanks {
				if value.BankName == nameOfBanks[key] {
					OutpBanks = append(OutpBanks, value)
				}
			}
		}
	}
	return
}
func FilterCurrency(data models.MainRequest, inpBanks []models.CurrencyBank) (OutpBanks []models.CurrencyBank) {
	s := strings.Join(data.Currency, "")
	for key := range nameOfCurrency {
		if strings.Contains(s, key) {
			for _, value := range inpBanks {
				if value.CodeAlpha == nameOfCurrency[key] {
					OutpBanks = append(OutpBanks, value)
				}
			}
		}
	}
	return
}

//min
func BestSale(inpBanks []models.CurrencyBank) (OutpBanks []models.CurrencyBank) {
	banks := append([]models.CurrencyBank(nil), inpBanks...)
	sort.Slice(banks, func(i, j int) bool {
		return banks[i].RateSale < banks[j].RateSale
	})

	for key := range nameOfCurrency {
		tmp := 0.0
		i := 0
		for _, value := range banks {
			if value.CodeAlpha == nameOfCurrency[key] {
				if i == 0 {
					OutpBanks = append(OutpBanks, value)
					tmp = value.RateSale
					i++
				} else if value.RateSale == tmp {
					OutpBanks = append(OutpBanks, value)
				}

			}
		}
	}
	return OutpBanks
}

//max
func BestBuy(inpBanks []models.CurrencyBank) (OutpBanks []models.CurrencyBank) {
	banks := append([]models.CurrencyBank(nil), inpBanks...)
	sort.Slice(banks, func(i, j int) bool {
		return banks[i].RateBuy > banks[j].RateBuy
	})

	for key := range nameOfCurrency {
		tmp := 0.0
		i := 0
		for _, value := range banks {
			if value.CodeAlpha == nameOfCurrency[key] {
				if i == 0 {
					OutpBanks = append(OutpBanks, value)
					tmp = value.RateBuy
					i++
				} else if value.RateBuy == tmp {
					OutpBanks = append(OutpBanks, value)
				}

			}
		}
	}
	return OutpBanks
}
