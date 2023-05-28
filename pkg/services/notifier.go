package services

import (
	"context"

	"exchange_rate/pkg/domain"
)

// Here we define mail service interface that we need for sending emails.
type IMailService interface {
	SendEmail(ctx context.Context, data *domain.CurrencyRate, receivers ...string) error
}

type notificationService struct {
	ctx                context.Context
	emailUserRepo      domain.IEmailRepository
	currencyRepository domain.ICurrencyRateRepository
	mailService        IMailService
}

func NewNotificationService(
	ctx context.Context,
	emailRepo domain.IEmailRepository,
	currencyRepository domain.ICurrencyRateRepository,
	mailService IMailService,
) domain.INotificationService {
	return &notificationService{
		ctx:                ctx,
		emailUserRepo:      emailRepo,
		mailService:        mailService,
		currencyRepository: currencyRepository,
	}
}

// Notify users via email due to our business logic.
func (n *notificationService) Notify(ctx context.Context, _ *domain.Notification) error {
	marketBTCUAH := domain.GetMarketBTCUAH()

	currency, err := n.currencyRepository.GetCurrencyRate(ctx, marketBTCUAH)
	if err != nil {
		return err
	}

	emails, err := n.emailUserRepo.GetAllEmails(ctx)
	if err != nil {
		return err
	}

	return n.mailService.SendEmail(ctx, currency, emails...)
}
