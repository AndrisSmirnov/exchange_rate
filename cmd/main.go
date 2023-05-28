package main

import (
	"context"
	"exchange_rate/pkg"
	"exchange_rate/pkg/controllers"
	"exchange_rate/pkg/domain"
	"exchange_rate/pkg/infrastructure/currency"
	"exchange_rate/pkg/infrastructure/mail"
	"exchange_rate/pkg/repository/file"
	"exchange_rate/pkg/services"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

type App struct {
	context         context.Context
	Controllers     *controllers.Controllers
	Services        *pkg.Services
	CurrencyGrabber currency.ICurrency
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	App := App{
		context: ctx,
	}

	if err := App.initServices(); err != nil {
		log.Fatalln(err)
	}

	if err := App.initControllers(); err != nil {
		log.Fatalln(err)
	}

	if err := App.launch(); err != nil {
		log.Fatalln(err)
	}

	go syscallWait(cancel)
	<-ctx.Done()
}

func syscallWait(cancelFunc func()) {
	syscallCh := make(chan os.Signal, 1)
	signal.Notify(syscallCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-syscallCh

	cancelFunc()
}

func (app *App) initServices() error {
	repoEmail, repoCurrency, err := file.NewFileSystemRepository(os.Getenv("FILE_DATABASE"))
	if err != nil {
		return err
	}

	repoCurrency.SetCurrencyRate(app.context, *domain.NewCurrencyRate(domain.GetMarketBTCUAH(), float64(1040000.22)))

	currency, err := currency.NewCurrency(app.context, repoCurrency)
	if err != nil {
		return err
	}

	mailSender, err := mail.NewEmailService()
	if err != nil {
		return err
	}

	userMailService := services.NewUserEmailService(app.context, repoEmail)

	notifierService := services.NewNotificationService(
		app.context, repoEmail, repoCurrency, mailSender,
	)

	services := pkg.NewServices(notifierService, repoCurrency, userMailService)

	app.Services = services
	app.CurrencyGrabber = currency

	return nil
}

func (app *App) initControllers() error {
	controllers, err := controllers.NewControllers(app.Services)
	if err != nil {
		return err
	}

	app.Controllers = controllers

	return nil
}

func (app *App) launch() error {
	app.Controllers.Start()
	return app.CurrencyGrabber.Start()
}
