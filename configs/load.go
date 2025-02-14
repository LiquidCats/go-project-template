package configs

import "github.com/kelseyhightower/envconfig"

func Load(prefix string) (Config, error) {
	var config Config

	err := envconfig.Process(prefix, &config)

	return config, err
}
