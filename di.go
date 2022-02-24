package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
	"os"
	"scoreBoard/api"
	"scoreBoard/api/handler"
	"scoreBoard/internal/config"
	scoreErrors "scoreBoard/internal/errors"
	"scoreBoard/internal/repository"
	"scoreBoard/pkg/amqp"
)

func injector(config *config.Config) (*http.Server, *scoreErrors.Errors) {
	dns := os.Getenv("DATABASE_URL")
	if dns != "" {
		config.Storage.Postgres.DNS = os.Getenv("DATABASE_URL")
	}

	var amqpConf = amqp.New(amqp.Config{
		Url:         config.Amqp.Url,
		ChannelName: config.Amqp.ChannelName,
	})

	db, err := provideDatabase(config.Storage.Postgres.DNS)
	if err != nil {
		return nil, err
	}
	storage := repository.ProvideStorage(db, amqpConf)

	_api := handler.ProvideAPI(config, handler.WithRepository(storage), handler.WithAmqp(amqpConf))

	//stats consumer to read msg from amqp server and updates in database
	go storage.Score().StartConsumingScoreFromAmqp()

	var router = api.NewRouter(_api)
	port := os.Getenv("PORT")
	Addr := fmt.Sprintf(":%s", port)
	return &http.Server{Handler: router, Addr: Addr}, nil
}

func provideDatabase(dns string) (*gorm.DB, *scoreErrors.Errors) {
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return nil, scoreErrors.Error(scoreErrors.Code("10.1"), err)
	}

	//creats connection pool
	//sqlDb, err := db.DB()
	//sqlDb.SetConnMaxLifetime(time.Hour * 1) //sets max connection time
	//sqlDb.SetMaxIdleConns(10)
	//sqlDb.SetConnMaxIdleTime(time.Minute * 1)
	//sqlDb.SetMaxOpenConns(15)
	return db, nil
}
