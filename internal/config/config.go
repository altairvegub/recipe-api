package config

import "github.com/kelseyhightower/envconfig"

var (
	RecipeApiPrefix = "RECIPE_API_PREFIX"
)

type HTTPServer struct {
	Port int `default:"8080"`
}

type Config struct {
	HTTPServer
}

func LoadConfigs(svcPrefix string) Config {
	var cfg Config
	envconfig.MustProcess(svcPrefix, &cfg)
	return cfg
}
