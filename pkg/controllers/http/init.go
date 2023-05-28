package http

import (
	"exchange_rate/pkg"

	"github.com/gin-gonic/gin"
)

type HTTPController struct {
	services *pkg.Services
	handlers *gin.Engine
}

func NewHttpControllers(
	services *pkg.Services,
) (*HTTPController, error) {
	if services == nil {
		return nil, ErrInitHTTPController
	}

	return &HTTPController{
		services: services,
		handlers: gin.Default(),
	}, nil
}

func (h *HTTPController) InitControllers() *gin.Engine {
	h.CurrencyController(h.handlers)
	h.EmailController(h.handlers)

	return h.handlers
}
