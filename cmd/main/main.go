package main

import (
	"context"
	"fmt"
	"github.com/drakenchef/RIP/MyMinio"
	"github.com/drakenchef/RIP/internal/app/config"
	"github.com/drakenchef/RIP/internal/app/dsn"
	"github.com/drakenchef/RIP/internal/app/handler"
	app "github.com/drakenchef/RIP/internal/app/pkg"
	"github.com/drakenchef/RIP/internal/app/redis"
	"github.com/drakenchef/RIP/internal/app/repository"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @title BITOP
// @version 1.0
// @description Bmstu Open IT Platform

// @contact.name API Support
// @contact.url https://vk.com/bmstu_schedule
// @contact.email bitop@spatecon.ru

// @license.name AS IS (NO WARRANTY)

// @host 127.0.0.1
// @schemes https http
// @BasePath /

func main() {
	logger := logrus.New()
	minioClient := MyMinio.NewMinioClient(logger)
	router := gin.Default()
	conf, err := config.NewConfig(logger)
	if err != nil {
		logger.Fatalf("Error with configuration reading: %s", err)
	}
	ctx := context.Background()
	redisClient, errRedis := redis.New(ctx, conf.Redis)
	if errRedis != nil {
		logger.Fatalf("Errof with redis connect: %s", err)
	}
	postgresString, errPost := dsn.FromEnv()
	if errPost != nil {
		logger.Fatalf("Error of reading postgres line: %s", errPost)
	}
	fmt.Println(postgresString)
	rep, errRep := repository.NewRepository(postgresString, logger)
	if errRep != nil {
		logger.Fatalf("Error from repository: %s", err)
	}
	hand := handler.NewHandler(logger, rep, minioClient, conf, redisClient)
	application := app.NewApp(conf, router, logger, hand)
	application.RunApp()
}
