package binance

import (
	"context"
	"exchange_rate/pkg/domain"
	"fmt"

	"github.com/gorilla/websocket"
)

type binance struct {
	ctx    context.Context
	config *config
	conn   *websocket.Conn
	repo   domain.ICurrencyRateRepository
}

type IBinance interface {
	IBinanceStart
	SubscribeToBinanceNewCurrency(string) error
	UnsubscribeToBinanceNewCurrency(string) error
}

type IBinanceStart interface {
	Start() error
	CloseConnection() error
}

func NewBinanceCurrencyRates(
	ctx context.Context,
	repo domain.ICurrencyRateRepository,
) (IBinance, error) {
	config, err := newBinanceConfig()
	if err != nil {
		return nil, err
	}

	return &binance{
		ctx:    ctx,
		config: config,
		repo:   repo,
	}, nil
}

func (b *binance) Start() error {
	connectionLink := fmt.Sprintf("%s%s", b.config.wsBaseURL, b.config.currencyRates)
	c, _, err := websocket.DefaultDialer.Dial(
		connectionLink,
		nil,
	)
	if err != nil {
		return ErrStartBinance
	}

	b.conn = c

	go b.reader()

	return nil
}

func (b *binance) CloseConnection() error {
	err := b.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return errorCloseConnection(err)
	}

	return nil
}
