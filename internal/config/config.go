package config

import "github.com/kelseyhightower/envconfig"

var (
	RecipeApiPrefix = "RECIPE_API_PREFIX"
)

type HTTPServer struct {
	Port int `default:"8080" envconfig:"HTTP_SERVER_PORT"`
}

type PostgresConfig struct {
	Host     string `default:"localhost"`
	Port     int    `default:"5432"`
	Database string `default:"recipe"`
	Password string `default:"password"`
	Username string `default:"recipe-user"`
}

type Config struct {
	HTTPServer
	PostgresConfig
}

func LoadConfigs(svcPrefix string) Config {
	var cfg Config
	envconfig.MustProcess(svcPrefix, &cfg)
	return cfg
}
