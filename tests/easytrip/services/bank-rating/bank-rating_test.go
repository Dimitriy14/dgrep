package bankRatingService

import (
	"reflect"
	"testing"

	"github.com/oreuta/easytrip/mocks"
	"github.com/oreuta/easytrip/models"
)

func TestGetBankRates(t *testing.T) {
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

	expect := []models.CurrencyBank{
		{
			BankName:  "Піреус Банк",
			CodeAlpha: "USD",
			RateBuy:   229,
			RateSale:  22,
		},
		{
			BankName:  "ПриватБанк",
			CodeAlpha: "USD",
			RateBuy:   66.6,
			RateSale:  33.3,
		},
	}

	mc := mocks.BankUAClientMock{
		Unpacked: banks,
	}

	ser := New(mc)
	res, err := ser.GetBankRates(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if !reflect.DeepEqual(expect, res) {
		t.Errorf("Want: %v, but get %v", expect, res)
	}
}
