package domain

import "context"

type Notification struct {
	User string `json:"user"`
}

func NewNotification() *Notification {
	return &Notification{}
}

type INotificationService interface {
	Notify(ctx context.Context, n *Notification) error
}
