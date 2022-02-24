package handler

import (
	"scoreBoard/internal/config"
	repo "scoreBoard/internal/repository"
	"scoreBoard/pkg/amqp"
)

type Repository interface {
	Game() repo.GameDAO
	Player() repo.PlayerDAO
	Score() repo.ScoreDAO
}

// API defines a holder type to hold service dependencies for endpoints
type API struct {
	config  *config.Config // config is the global configuration object
	Storage Repository
	Amqp    amqp.Amqp
}

//ProvideAPI assembles the factory methods
func ProvideAPI(config *config.Config, fns ...func(*API)) *API {
	var api = &API{
		config: config,
	}

	for _, fn := range fns {
		if fn != nil {
			fn(api)
		}
	}

	return api
}

////WithMySql is a factory function to create sql connection
//func WithMySql(client *mysql.Postgres) func(*API) {
//	return func(api *API) {
//		api.Postgres = client
//	}
//}

//WithRepository is a factory function to create repository methods
func WithRepository(client Repository) func(*API) {
	return func(api *API) {
		api.Storage = client
	}
}

// WithAmqp injects pkg/amqp/Amqp client provider function in API
func WithAmqp(client amqp.Amqp) func(*API) {
	return func(api *API) {
		api.Amqp = client
	}
}
