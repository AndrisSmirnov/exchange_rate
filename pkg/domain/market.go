package domain

import "fmt"

type Market struct {
	BaseCurrency  string `json:"base_currency"`
	QuoteCurrency string `json:"quote_currency"`
}

func NewMarket(base, quote string) *Market {
	return &Market{
		BaseCurrency:  base,
		QuoteCurrency: quote,
	}
}

func (m Market) ToString() string {
	return fmt.Sprintf("%s%s", m.BaseCurrency, m.QuoteCurrency)
}
