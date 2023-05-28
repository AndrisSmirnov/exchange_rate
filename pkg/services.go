package pkg

import "exchange_rate/pkg/domain"

type Services struct {
	NotificationService domain.INotificationService
	CurrencyService     domain.ICurrencyRateRepository
	EmailService        domain.IEmailService
}

func NewServices(
	notification domain.INotificationService,
	currency domain.ICurrencyRateRepository,
	email domain.IEmailService,
) *Services {
	return &Services{
		NotificationService: notification,
		CurrencyService:     currency,
		EmailService:        email,
	}
}
