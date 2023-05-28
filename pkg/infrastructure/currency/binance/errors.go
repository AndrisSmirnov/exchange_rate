package binance

import (
	"errors"
	"fmt"
)

var (
	ErrInitBinance  = errors.New("error initialization binance config")
	ErrStartBinance = errors.New("error while start binance")
)

func errorCloseConnection(err error) error {
	return errors.New(fmt.Sprintf("Error closing Binance WebSocket connection:%s", err.Error()))
}
