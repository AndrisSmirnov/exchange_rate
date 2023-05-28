package http

import (
	"exchange_rate/pkg/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HTTPController) EmailController(router *gin.Engine) {

	emailGroup := router.Group("/")

	emailGroup.Use()
	{
		subGroup := router.Group("/subscribe")
		subGroup.POST("/",
			func(c *gin.Context) {
				context := c.Request.Context()

				email := domain.NewUserEmail(c.PostForm("email"))

				if err := email.Validate(); err != nil {
					c.JSON(http.StatusBadRequest, nil)
					return
				}

				if err := h.services.EmailService.NewEmailUser(context, email); err != nil {
					if err == domain.ErrAlreadyExist {
						c.JSON(http.StatusConflict, nil)
						return
					}
					c.JSON(http.StatusInternalServerError, nil)
					return
				}

				c.JSON(http.StatusOK, nil)
			},
		)

		senderGroup := router.Group("/sendEmails")
		senderGroup.POST("/",
			func(c *gin.Context) {
				go func() {
					if err := h.services.NotificationService.Notify(
						c.Request.Context(),
						domain.NewNotification(),
					); err != nil {
						log.Printf("error on sending emails: %v\n", err)
					}
				}()

				c.JSON(http.StatusOK, nil)
			},
		)

	}
}
