package binance

import (
	"encoding/json"
	"exchange_rate/pkg/domain"
	"log"
	"strings"
)

type BinanceCurrencyResponse struct {
	Stream       string       `json:"stream"`
	ResponseData CurrencyData `json:"data"`
	ID           int          `json:"id"`
}

type CurrencyData struct {
	EventType                   string  `json:"e"`
	EventTime                   uint64  `json:"E"`
	Symbol                      string  `json:"s"`
	PriceChange                 float64 `json:"p,string"`
	PriceChangePercent          float64 `json:"P,string"`
	WeightedAveragePrice        float64 `json:"w,string"`
	FirstTradePrice             float64 `json:"x,string"`
	LastPrice                   float64 `json:"c,string"`
	LastQuantity                float64 `json:"Q,string"`
	BestBidPrice                float64 `json:"b,string"`
	BestBidQuantity             float64 `json:"B,string"`
	BestAskPrice                float64 `json:"a,string"`
	BestAskQuantity             float64 `json:"A,string"`
	OpenPrice                   float64 `json:"o,string"`
	HighPrice                   float64 `json:"h,string"`
	LowPrice                    float64 `json:"l,string"`
	TotalTradedBaseAssetVolume  float64 `json:"v,string"`
	TotalTradedQuoteAssetVolume float64 `json:"q,string"`
	StatisticsOpenTime          int     `json:"O"`
	StatisticsCloseTime         int     `json:"C"`
	FirstTradeID                int     `json:"F"`
	LastTradeID                 int     `json:"L"`
	TotalNumberOfTrades         int     `json:"n"`
}

func (b *binance) reader() {
	for {
		_, message, err := b.conn.ReadMessage()
		if err != nil {
			log.Println("Error reading WebSocket message:", err)
			return
		}

		// Parse the message and extract the price
		// The message format will depend on the specific WebSocket endpoint being used

		response := BinanceCurrencyResponse{}
		err = json.Unmarshal(message, &response)
		if err != nil {
			log.Println("Error parsing WebSocket message:", err)
			return
		}

		if response.ID != b.config.binanceID {
			b.repo.SetCurrencyRate(b.ctx, *domain.NewCurrencyRate(
				*domain.NewMarket(strings.Split(response.Stream, "@ticker")[0], ""),
				response.ResponseData.LastPrice,
			))
		}
	}
}
