package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config contains the required app configuration parameters that are read from config file
type Config struct {
	Amqp    Amqp
	Storage Storage
}

type Amqp struct {
	Url         string //amqp://guest:guest@localhost:5672/
	ChannelName string //score
}

type Storage struct {
	Postgres Postgres
}
type Postgres struct {
	DNS string
}

func init() {
	// command line flags to quickly specify some common values
	flag.Bool("debug", false, "enable debug mode")
	flag.String("config", "config.yaml", "path to configuration file")

	flag.Parse() // parse the command-line

	// configure default values for some configuration parameters
	// these can be overridden and have least precedence
	viper.SetDefault("http.addr", "127.0.0.1")
	viper.SetDefault("http.port", "8080")
	viper.SetConfigType("yaml")
}

// Load ...
func Load() (*Config, error) {
	// configure viper to load values from environment
	viper.SetEnvPrefix("SCORE")
	viper.AutomaticEnv()

	// bind viper with the flag set
	if err := viper.BindPFlags(flag.CommandLine); err != nil {
		return nil, errors.Wrap(err, "failed to bind command-line flags")
	}

	// check if --config was provided...
	if path := viper.GetString("config"); path != "" {
		// ...and read in the values from the file
		viper.SetConfigFile(path)
	}

	var config = &Config{}
	if err := unmarshal(config); err != nil {
		return nil, errors.Wrap(err, "failed to un-marshal config values")
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		unmarshal(&config)
	})

	return config, nil
}

func unmarshal(rawVal interface{}) error {
	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			// 	do nothing
		default:
			return err
		}
	}
	// We need to unmarshal before the env binding to make sure that keys of maps are bound just like the struct fields
	// We silence errors here because we'll unmarshal a second time
	err := viper.Unmarshal(rawVal)
	return err
}
