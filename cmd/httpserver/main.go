package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adshao/go-binance/v2"
	"github.com/go-playground/validator/v10"
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
		port          = os.Getenv("PORT")
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

	validator := validator.New()

	updateOrderUC := app_usecases.NewUpdateOrderFillUseCase(orderRepository)
	createOrderUC := app_usecases.NewCreateOrderUseCase(orderRepository, binanceExchangeProvider)
	getOrdersUC := app_usecases.NewGetOrdersUseCase(orderRepository)
	getOrderUC := app_usecases.NewGetOrderUseCase(orderRepository)
	cancelOrderUC := app_usecases.NewCancelOrderUseCase(orderRepository, binanceExchangeProvider)

	createOrderHandler := http_handlers.NewCreateOrderHandler(createOrderUC, validator).Handle
	getOrdersHandler := http_handlers.NewGetOrdersHandler(getOrdersUC).Handle
	getOrderHandler := http_handlers.NewGetOrderHandler(getOrderUC, validator).Handle
	cancelOrderHandler := http_handlers.NewCancelOrderHandler(cancelOrderUC, validator).Handle

	userAuthMiddleware := http_middlewares.UserAuthMiddleware

	e := echo.New()

	e.POST("/orders", userAuthMiddleware(createOrderHandler))
	e.GET("/orders/:order_id", userAuthMiddleware(getOrderHandler))
	e.GET("/orders", userAuthMiddleware(getOrdersHandler))
	e.PATCH("/orders/cancel/:order_id", userAuthMiddleware(cancelOrderHandler))

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

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
