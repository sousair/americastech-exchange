package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	app_usecases "github.com/sousair/americastech-exchange/internal/application/usecases"
	gorm_models "github.com/sousair/americastech-exchange/internal/infra/database/models"
	gorm_repositories "github.com/sousair/americastech-exchange/internal/infra/database/repositories"
	binance_exchange "github.com/sousair/americastech-exchange/internal/infra/providers/exchange"
	http_handlers "github.com/sousair/americastech-exchange/internal/presentation/http/handlers"
	http_middlewares "github.com/sousair/americastech-exchange/internal/presentation/http/middlewares"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		binanceApiKey     = os.Getenv("BINANCE_API_KEY")
		binanceSecret     = os.Getenv("BINANCE_API_SECRET")
		binanceApiBaseUrl = os.Getenv("BINANCE_API_BASE_URL")

		postgresConnectionUrl = os.Getenv("POSTGRES_CONNECTION_URL")
	)

	db, err := gorm.Open(postgres.Open(postgresConnectionUrl), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&gorm_models.Order{})

	orderRepository := gorm_repositories.NewOrderRepository(db)
	binanceExchangeProvider := binance_exchange.NewBinanceExchangeProvider(binanceApiKey, binanceSecret, binanceApiBaseUrl)

	createUserUC := app_usecases.NewCreateOrderUseCase(orderRepository, binanceExchangeProvider)

	createOrderHandler := http_handlers.NewCreateOrderHandler(createUserUC).Handle

	userAuthMiddleware := http_middlewares.UserAuthMiddleware

	e := echo.New()

	e.POST("/orders", userAuthMiddleware(createOrderHandler))

	e.Logger.Fatal(e.Start(":7070"))
}
