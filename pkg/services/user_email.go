package services

import (
	"context"

	"exchange_rate/pkg/domain"
)

type userEmailService struct {
	ctx           context.Context
	emailUserRepo domain.IEmailRepository
}

func NewUserEmailService(
	ctx context.Context,
	emailRepo domain.IEmailRepository,
) domain.IEmailService {
	return &userEmailService{
		ctx:           ctx,
		emailUserRepo: emailRepo,
	}
}

func (e *userEmailService) NewEmailUser(ctx context.Context, emailUser *domain.UserEmail) error {
	exist, err := e.emailUserRepo.EmailExist(ctx, emailUser.Email.ToString())
	if err != nil {
		return err
	}

	if exist {
		return domain.ErrAlreadyExist
	}

	return e.emailUserRepo.SaveEmail(ctx, emailUser)
}
