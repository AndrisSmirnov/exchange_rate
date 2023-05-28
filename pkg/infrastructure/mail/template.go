package mail

import (
	"bytes"
	"exchange_rate/pkg/domain"
	"fmt"
	"html/template"
	"strings"
	"time"
)

type TemplateFillData struct {
	CurrentTime   string
	CurrentPrise  string
	BaseCurrency  string
	QuoteCurrency string
}

func (e *EmailSender) crateTemplate(rate *domain.CurrencyRate) []byte {
	t, _ := template.ParseFiles("./files/template.html")
	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("%s \n%s\n\n", e.config.subject, mimeHeaders)))

	templateFillData := TemplateFillData{
		CurrentTime:   time.Now().Format("02.01.2006 15:04"),
		CurrentPrise:  fmt.Sprintf("%.2f", rate.Rate),
		BaseCurrency:  strings.ToUpper(rate.Market.BaseCurrency),
		QuoteCurrency: strings.ToUpper(rate.Market.QuoteCurrency),
	}

	t.Execute(&body, templateFillData)

	return body.Bytes()
}
