package main

import (
	services_provider "backend/pkg/services/provider"
	"context"
	"errors"
	"github.com/adshao/go-binance/v2/futures"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	futures.WebsocketKeepalive = true
	futures.WebsocketTimeout = time.Second * 600
}

func main() {
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	providerService := services_provider.NewProviderService(ctx)

	if err := run(providerService, ctx, cancelCtx); err != nil {
		providerService.LoggerService().Error().Printf("server encountered an error: %v", err)
	}
}

func run(providerService *services_provider.ProviderService, parentCtx context.Context, parentCancelCtx context.CancelFunc) error {
	// if err := providerService.TradeService().Init(); err != nil {
	// 	providerService.LoggerService().Error().Fatalf("failed to init trades: %v", err)
	// }

	providerService.LoggerService().Info().Printf("starting server: %s", providerService.ConfigService().GetConfig().HttpServer.Address)

	providerService.UserService().Init()
	providerService.BalanceService().InitBalance(providerService.UserService().GetBalance())
	providerService.BotService().InitGuard()

	go providerService.BotService().RunGuardChannel()
	go providerService.BotService().RunAddDealChannel()
	go providerService.OrderService().RunOrderChannel()
	go providerService.BotService().RunDealChannel()
	go providerService.BotService().RunChannel()
	go providerService.WebsocketService().Start()
	go providerService.ExchangeWebsocketService().Start()

	server := &http.Server{
		Addr:         providerService.ConfigService().GetConfig().HttpServer.Address,
		Handler:      providerService.RouterService().GetRouter(),
		ReadTimeout:  providerService.ConfigService().GetConfig().HttpServer.ReadTimeout,
		WriteTimeout: providerService.ConfigService().GetConfig().HttpServer.WriteTimeout,
		IdleTimeout:  providerService.ConfigService().GetConfig().HttpServer.IdleTimeout,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			providerService.LoggerService().Error().Printf("failed to start server: %v", err)
		}
	}()

	<-stop
	providerService.LoggerService().Info().Printf("shutting down server...")

	parentCancelCtx()
	providerService.Shutdown()

	shutdownContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownContext); err != nil {
		return err
	}

	providerService.LoggerService().Info().Printf("server stopped")

	return nil
}
