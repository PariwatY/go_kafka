package main

import (
	"producer/controllers"
	"producer/services"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	// Use viper for get config from config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	// Config producer to call kafka servers
	producer, err := sarama.NewSyncProducer(viper.GetStringSlice("kafka.servers"), nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	eventProducer := services.NewEventProducer(producer)
	accountService := services.NewAccountService(eventProducer)
	accountController := controllers.NewAccountController(accountService)

	app := echo.New()

	// Open New Account
	app.POST("/openAccount", accountController.OpenAccount)
	// Deposite Fund
	app.POST("/depositFund", accountController.DepositFund)
	// Withdraw Fund
	app.POST("/withdrawFund", accountController.WithdrawFund)
	// Close Account
	app.POST("/closeAccount", accountController.CloseAccount)

	app.Start(":8081")
}
