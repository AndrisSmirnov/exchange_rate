package currency

import (
	"context"
	"exchange_rate/pkg/domain"
	"exchange_rate/pkg/infrastructure/currency/binance"
)

type currency struct {
	binance binance.IBinance
}

type ICurrency interface {
	ICurrencyLaunch
	SubscribeNewCurrency(string) error
	UnsubscribeNewCurrency(string) error
}

type ICurrencyLaunch interface {
	Start() error
	Stop() error
}

func NewCurrency(
	ctx context.Context,
	repo domain.ICurrencyRateRepository,
) (ICurrency, error) {
	binance, err := binance.NewBinanceCurrencyRates(ctx, repo)
	if err != nil {
		return nil, err
	}

	return &currency{
		binance: binance,
	}, nil
}

func (c *currency) Start() error {
	return c.binance.Start()
}

func (c *currency) Stop() error {
	return c.binance.CloseConnection()
}

func (c *currency) SubscribeNewCurrency(data string) error {
	return c.binance.SubscribeToBinanceNewCurrency(data)
}

func (c *currency) UnsubscribeNewCurrency(data string) error {
	return c.binance.UnsubscribeToBinanceNewCurrency(data)
}
