package file

import (
	"context"
	"exchange_rate/pkg/domain"
	"fmt"
	"os"
	"strings"
	"sync"
)

type fileSystemRepository struct {
	filePath      string
	indexEmail    map[string]struct{}
	indexCurrency map[string]*domain.CurrencyRate
	// file mutex
	fm sync.RWMutex
	// index mutex
	im sync.RWMutex
}

func NewFileSystemRepository(filePath string) (
	domain.IEmailRepository,
	domain.ICurrencyRateRepository,
	error) {
	fileSystem := &fileSystemRepository{
		filePath:      filePath,
		indexCurrency: map[string]*domain.CurrencyRate{},
		fm:            sync.RWMutex{},
		im:            sync.RWMutex{},
	}

	if err := fileSystem.loadEmailIndex(); err != nil {
		return nil, nil, err
	}

	return fileSystem, fileSystem, nil
}

func (f *fileSystemRepository) SaveEmail(_ context.Context, eu *domain.UserEmail) error {
	f.fm.Lock()
	defer f.fm.Unlock()

	file, err := os.OpenFile(f.filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("failed to open data file: %w", err)
	}
	defer file.Close()

	if _, err = file.WriteString(addEndOfTheLine(eu.Email.ToString())); err != nil {
		return fmt.Errorf("failed write to file")
	}

	f.im.Lock()
	defer f.im.Unlock()

	f.indexEmail[eu.Email.ToString()] = struct{}{}

	return nil
}

func (f *fileSystemRepository) GetByEmail(
	_ context.Context,
	email string,
) (*domain.UserEmail, error) {
	f.im.RLock()
	defer f.im.RUnlock()

	_, ok := f.indexEmail[email]
	if !ok {
		return nil, domain.ErrNotFound
	}

	return domain.NewUserEmail(email), nil
}

func (f *fileSystemRepository) GetAllEmails(
	_ context.Context,
) ([]string, error) {
	f.fm.RLock()
	defer f.fm.RUnlock()

	if _, err := os.Stat(f.filePath); os.IsNotExist(err) {
		return []string{}, nil
	}

	data, err := os.ReadFile(f.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file by path: %s", f.filePath)
	}

	rows := strings.Split(string(data), "\n")

	emails := make([]string, 0, len(rows))

	for _, row := range rows {
		if row == "" {
			continue
		}

		emails = append(emails, row)
	}

	return emails, nil
}

func (f *fileSystemRepository) EmailExist(_ context.Context, email string) (bool, error) {
	f.im.RLock()
	defer f.im.RUnlock()

	_, ok := f.indexEmail[email]

	return ok, nil
}

func (f *fileSystemRepository) GetCurrencyRate(
	_ context.Context,
	market domain.Market,
) (*domain.CurrencyRate, error) {
	f.im.RLock()
	defer f.im.RUnlock()

	rate, ok := f.indexCurrency[market.ToString()]
	if !ok {
		return nil, domain.ErrNotFound
	}

	rate.Market = market

	return rate, nil
}

func (f *fileSystemRepository) SetCurrencyRate(_ context.Context, rate domain.CurrencyRate) error {
	f.im.Lock()
	defer f.im.Unlock()

	f.indexCurrency[rate.Market.ToString()] = &rate

	return nil
}

func addEndOfTheLine(data string) string {
	return data + "\n"
}
