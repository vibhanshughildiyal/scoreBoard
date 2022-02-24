package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"runtime"
	"scoreBoard/internal/config"
)

func init() {
	// maximise the number of operating system threads
	runtime.GOMAXPROCS(runtime.NumCPU())

	// initialize logger
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true})
}

func main() {
	if cfg, err := config.Load(); err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	} else {
		if server, err := injector(cfg); err != nil {
			log.Fatal().Err(err.Error).Msg("failed to initialize")
		} else {
			if e1 := server.ListenAndServe(); e1 != nil { // blocks when server is running
				log.Fatal().Err(e1).Msg("failed to start server")
			}
		}
	}
}
