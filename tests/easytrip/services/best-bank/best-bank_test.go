package bestBankService

import (
	"reflect"
	"testing"

	"github.com/oreuta/easytrip/mocks"
	"github.com/oreuta/easytrip/models"
)

func TestGetBestBanks(t *testing.T) {

	banks := []models.CurrencyBank{
		{
			BankName:  "ПриватБанк",
			CodeAlpha: "USD",
			RateBuy:   66.6,
			RateSale:  33.3,
		},
		{
			BankName:  "Піреус Банк",
			CodeAlpha: "USD",
			RateBuy:   229,
			RateSale:  22,
		},
	}

	req := models.MainRequest{
		Currency: []string{"usd"},
		Option:   "sale",
		Bank:     []string{"privat", "pireus"},
	}

	expect1 := []models.CurrencyBank{
		{
			BankName:  "Піреус Банк",
			CodeAlpha: "USD",
			RateBuy:   229,
		},
		{
			BankName:  "ПриватБанк",
			CodeAlpha: "USD",
			RateBuy:   66.6,1
		},
	}

	expect2 := []models.CurrencyBank{
		{
			BankName:  "Піреус Банк",
			CodeAlpha: "USD",
			RateSale:  22,
		},
		{
			BankName:  "ПриватБанк",
			CodeAlpha: "USD",
			RateSale:  33.3,
		},
	}

	mc := mocks.BankUAClientMock{
		Unpacked: banks,
	}

	ser := New(mc)
	res1, res2, err := ser.GetBestBanks(req)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if !reflect.DeepEqual(expect1, res1) {
		t.Errorf("Want: %v, but get %v", expect1, res1)
	}

	if !reflect.DeepEqual(expect2, res2) {
		t.Errorf("Want: %v, but get %v", expect2, res2)
	}
}
