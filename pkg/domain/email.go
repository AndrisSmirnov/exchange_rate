package domain

import (
	"context"
	"exchange_rate/pkg/domain/vo"
)

type UserEmail struct {
	Email vo.Email `json:"email"`
}

func (e *UserEmail) Validate() error {
	if err := e.Email.Validate(); err != nil {
		return ErrBadRequest
	}

	return nil
}

func NewUserEmail(email string) *UserEmail {
	return &UserEmail{
		Email: vo.Email(email),
	}
}

type IEmailService interface {
	NewEmailUser(ctx context.Context, eu *UserEmail) error
}

type IEmailRepository interface {
	SaveEmail(ctx context.Context, eu *UserEmail) error
	EmailExist(ctx context.Context, email string) (bool, error)
	GetByEmail(ctx context.Context, email string) (*UserEmail, error)
	GetAllEmails(ctx context.Context) ([]string, error)
}
