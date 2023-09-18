package main

import (
	"fmt"
	"github.com/drakenchef/RIP/internal/app/config"
	"github.com/drakenchef/RIP/internal/app/dsn"
	"github.com/drakenchef/RIP/internal/app/handler"
	"github.com/drakenchef/RIP/internal/app/pkg"
	"github.com/drakenchef/RIP/internal/app/repository"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	router := gin.Default()
	conf, err := config.NewConfig(logger)
	if err != nil {
		logger.Fatalf("Error with configuration reading: %s", err)
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
	hand := handler.NewHandler(logger, rep)
	application := pkg.NewApp(conf, router, logger, hand)
	application.RunApp()
}
