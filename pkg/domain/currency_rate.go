package domain

import (
	"context"
)

func GetMarketBTCUAH() Market {
	return Market{
		BaseCurrency:  "btc",
		QuoteCurrency: "uah",
	}
}

type CurrencyRate struct {
	Market
	Rate float64 `json:"rate"`
}

func NewCurrencyRate(market Market, rate float64) *CurrencyRate {
	return &CurrencyRate{
		Market: market,
		Rate:   rate,
	}
}

type ICurrencyRateRepository interface {
	ICurrencyRateRepositoryGet
	ICurrencyRateRepositorySet
}

type ICurrencyRateRepositorySet interface {
	SetCurrencyRate(ctx context.Context, rate CurrencyRate) error
}

type ICurrencyRateRepositoryGet interface {
	GetCurrencyRate(ctx context.Context, market Market) (*CurrencyRate, error)
}
