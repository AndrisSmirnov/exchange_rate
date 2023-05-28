package vo

import "net/mail"

type Email string

func (e Email) ToString() string {
	return string(e)
}

func (e Email) Validate() error {
	_, err := mail.ParseAddress(e.ToString())
	if err != nil {
		return ErrNotValidEmail
	}

	return nil
}
