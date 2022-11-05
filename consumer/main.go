package main

import (
	"consumer/config"
	"consumer/repositories"
	"consumer/services"
	"context"
	"events"
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
)

func init() {
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

	// Create consumer setting to kafka server
	consumer, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka.servers"), viper.GetString("kafka.group"), nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// Initial Database Connection
	db := config.InitialDB()

	//Initial repositories
	accountRepo := repositories.NewAccountRepository(db)
	//Initial event handler
	accountEventHandler := services.NewAccountEventHandler(accountRepo)
	//Initial customer handler for getting sarama.ConsumerGroupHandler to use in Consume
	accountConsumerHandler := services.NewConsumerHandler(accountEventHandler)

	fmt.Println("======Account Consumer Started======")
	for {
		consumer.Consume(context.Background(), events.Topics, accountConsumerHandler)
	}

}
