package configs

import "fmt"

type Config struct {
	Redis    RedisConfig    `yaml:"redis" envconfig:"REDIS"`
	Postgres PostgresConfig `yaml:"postgres" envconfig:"POSTGRES"`
}

type PostgresConfig struct {
	Host     string `yaml:"host" envconfig:"HOST"`
	Port     string `yaml:"port" envconfig:"PORT"`
	User     string `yaml:"user" envconfig:"USER"`
	Password string `yaml:"password" envconfig:"PASSWORD"`
	Database string `yaml:"database" envconfig:"DATABASE"`
}

func (c PostgresConfig) ToDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.User, c.Password, c.Host, c.Port, c.Database)
}

type RedisConfig struct {
	Host     string `yaml:"host" envconfig:"HOST"`
	Port     string `yaml:"port" envconfig:"PORT"`
	DB       int    `yaml:"db" envconfig:"DB"`
	Password string `yaml:"password,omitempty" envconfig:"PASSWORD,omitempty"`
}
