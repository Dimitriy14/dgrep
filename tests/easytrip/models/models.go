package models

//MainRequest gives fields for request from main page
type MainRequest struct {
	Currency []string
	Option   string
	Bank     []string
}

type FullRequest struct {
	Start  MainRequest
	Link   string
	Method string
}

//CurrencyBank gives definition of banks
type CurrencyBank struct {
	BankName  string
	CodeAlpha string
	RateBuy   float64 `json:",string"`
	RateSale  float64 `json:",string"`
}

//CurrencyBanks is an array of CurrencyBank
type CurrencyBanks []CurrencyBank

func Bank() map[string]string {
	banksMap := map[string]string{
		"privat": "ПриватБанк",
		"otp":    "ОТП Банк",
		"pireus": "Піреус Банк",
		"kredo":  "Кредобанк",
	}
	return banksMap
}

func Currency() map[string]string {
	currencyMap := map[string]string{
		"usd": "USD",
		"eur": "EUR",
	}
	return currencyMap
}

// Registration

type User struct {
	Name     string `valid:"Required;MinSize(3);Alpha"`
	Login    string `valid:"Required;MinSize(3)"`
	Password string `valid:"Required;MinSize(4)"`
}
type HistoryStruct struct {
	Link     string
	Banks    string
	Currency string
	Option   string
}
