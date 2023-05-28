package binance

import (
	"os"
	"strconv"
)

type config struct {
	wsBaseURL     string
	currencyRates string
	binanceID     int
}

func newBinanceConfig() (*config, error) {
	binanceWsBaseURL := os.Getenv("BINANCE_WS_BASE")
	if binanceWsBaseURL == "" {
		return nil, ErrInitBinance
	}

	binanceCurrencyRates := os.Getenv("BINANCE_WS_CURRENCY_RATES")
	if binanceCurrencyRates == "" {
		return nil, ErrInitBinance
	}

	binanceIDStr := os.Getenv("BINANCE_ID")
	if binanceIDStr == "" {
		return nil, ErrInitBinance
	}
	binanceID, err := strconv.Atoi(binanceIDStr)
	if err != nil {
		return nil, ErrInitBinance
	}

	return &config{
		wsBaseURL:     binanceWsBaseURL,
		currencyRates: binanceCurrencyRates,
		binanceID:     binanceID,
	}, nil
}
