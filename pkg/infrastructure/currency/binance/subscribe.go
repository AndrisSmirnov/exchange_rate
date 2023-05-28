package binance

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func newSubscribe(sub string, id int) string {
	subReq := fmt.Sprintf(`{
"method": "SUBSCRIBE",
"params":
[
"%s@ticker"
],
"id": %d
}`, sub, id)

	return subReq
}

func (b *binance) SubscribeToBinanceNewCurrency(data string) error {
	byteSubscribeData := []byte(newSubscribe(data, b.config.binanceID))
	if err := b.conn.WriteMessage(websocket.TextMessage, byteSubscribeData); err != nil {
		return err
	}

	return nil
}
