package mail

type Config struct {
	address  string
	appKey   string
	subject  string
	smtpHost string
	smtpPort string
}

func newConfig(address, appKey, subject, smtpHost, smtpPort string) (*Config, error) {
	if address == "" || appKey == "" || subject == "" || smtpHost == "" || smtpPort == "" {
		return nil, ErrMailConfig
	}

	return &Config{
		address:  address,
		appKey:   appKey,
		subject:  subject,
		smtpHost: smtpHost,
		smtpPort: smtpPort,
	}, nil
}
