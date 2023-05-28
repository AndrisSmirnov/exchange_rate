package http

import (
	"exchange_rate/pkg/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HTTPController) CurrencyController(router *gin.Engine) {

	currencyGroup := router.Group("/rate")

	currencyGroup.Use()
	{
		currencyGroup.GET("/",
			func(c *gin.Context) {
				context := c.Request.Context()

				// mock to create request for BTCUAH Market
				// in the future it will be for many currency
				market := domain.GetMarketBTCUAH()

				response, err := h.services.CurrencyService.GetCurrencyRate(context, market)
				if err != nil {
					c.JSON(400, nil)
					return
				}

				c.JSON(http.StatusOK, response.Rate)
			},
		)
	}
}
