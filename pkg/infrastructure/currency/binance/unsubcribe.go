package binance

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func newUnsubscribe(sub string, id int) string {
	subReq := fmt.Sprintf(`{
"method": "UNSUBSCRIBE",
"params":
[
"%s@ticker"
],
"id": %d
}`, sub, id)

	return subReq
}

func (b *binance) UnsubscribeToBinanceNewCurrency(data string) error {
	byteUnsubscribeData := []byte(newUnsubscribe(data, b.config.binanceID))
	if err := b.conn.WriteMessage(websocket.TextMessage, byteUnsubscribeData); err != nil {
		return err
	}

	return nil
}
