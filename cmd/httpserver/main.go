package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adshao/go-binance/v2"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	app_usecases "github.com/sousair/americastech-exchange/internal/application/usecases"
	"github.com/sousair/americastech-exchange/internal/core/usecases"
	gorm_models "github.com/sousair/americastech-exchange/internal/infra/database/models"
	gorm_repositories "github.com/sousair/americastech-exchange/internal/infra/database/repositories"
	binance_exchange "github.com/sousair/americastech-exchange/internal/infra/providers/exchange/binance"
	http_handlers "github.com/sousair/americastech-exchange/internal/presentation/http/handlers"
	http_middlewares "github.com/sousair/americastech-exchange/internal/presentation/http/middlewares"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		env           = os.Getenv("ENV")
		binanceApiKey = os.Getenv("BINANCE_API_KEY")
		binanceSecret = os.Getenv("BINANCE_API_SECRET")

		postgresConnectionUrl = os.Getenv("POSTGRES_CONNECTION_URL")
	)

	if env == "development" {
		binance.UseTestnet = true
	}

	db, err := gorm.Open(postgres.Open(postgresConnectionUrl), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&gorm_models.Order{})

	orderRepository := gorm_repositories.NewOrderRepository(db)
	binanceExchangeProvider := binance_exchange.NewBinanceExchangeProvider(binanceApiKey, binanceSecret)

	updateOrderUC := app_usecases.NewUpdateOrderFillUseCase(orderRepository)
	createOrderUC := app_usecases.NewCreateOrderUseCase(orderRepository, binanceExchangeProvider)
	getOrdersUC := app_usecases.NewGetOrdersUseCase(orderRepository)

	go func() {
		for orderEvent := range binanceExchangeProvider.UpdateOrderEventChan {
			err := updateOrderUC.Update(usecases.UpdateOrderFillParams{
				ExternalID: orderEvent.ExternalID,
				Status:     orderEvent.Status,
				Price:      orderEvent.Price,
			})

			if err != nil {
				fmt.Printf("Error updating order: %v\n", err)
			}

			fmt.Printf("Order updated: %s\n", orderEvent.ExternalID)
		}
	}()

	createOrderHandler := http_handlers.NewCreateOrderHandler(createOrderUC).Handle
	getOrdersHandler := http_handlers.NewGetOrdersHandler(getOrdersUC).Handle

	userAuthMiddleware := http_middlewares.UserAuthMiddleware

	e := echo.New()

	e.POST("/orders", userAuthMiddleware(createOrderHandler))
	e.GET("/orders", userAuthMiddleware(getOrdersHandler))

	e.Logger.Fatal(e.Start(":7070"))
}
