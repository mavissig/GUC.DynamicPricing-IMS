// @title           DynamicPricing-API
// @version         1.0
// @description     API для работы с динамическим ценообразованием

// @contact.name   API Support(Егор)
// @contact.url    https://t.me/senior_stepik

// @license.name  MIT
// @license.url   http://opensource.org/licenses/MIT

// @host      127.0.0.1:8080
// @BasePath  /api/v1

// https://github.com/mavissig/GUC.DynamicPricing-IMS

package main

import (
	"context"
	"fmt"
	"github.com/mavissig/GUC.DynamicPricing-IMS/internal/ims/domain"
	"github.com/mavissig/GUC.DynamicPricing-IMS/internal/ims/repository"
	"github.com/mavissig/GUC.DynamicPricing-IMS/internal/ims/repository/redis"
	"github.com/mavissig/GUC.DynamicPricing-IMS/internal/ims/transport"
	kafkaSrv "github.com/mavissig/GUC.DynamicPricing-IMS/internal/ims/transport/kafka-server"
	"os"
	"os/signal"
)

func main() {
	defer fmt.Println("Сервер успешно остановлен")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	transportCfg := transport.LoadConfig()
	repositoryCfg := repository.LoadConfig()

	r := redis.New(repositoryCfg)

	useCase := domain.New(r)

	kafkaServer := kafkaSrv.New(transportCfg, useCase)
	kafkaServer.Run(ctx)

	//httpServer := httpSrv.New(transportCfg, useCase)
	//go httpServer.Run()

	<-quit

	fmt.Println("Попытка завершить сервер...")
}
