package driver

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env     string `env:"ENVIRONMENT" envDefault:"production"`
	Port    int    `env:"PORT" envDefault:"8081"`
	Timeout int    `env:"TIMEOUT" envDefault:"10"`
	DBHost  string `env:"DB_HOST" envDefault:"localhost"`
	DBPort  string `env:"DB_PORT" envDefault:"3306"`
	DBUser  string `env:"DB_USER" envDefault:"test"`
	DBPass  string `env:"DB_PASS" envDefault:"test"`
	DBName  string `env:"DB_NAME" envDefault:"zero_system"`
}

// 環境変数の構造体を返却
//
// @return *Config 環境変数の構造体
//
// @return error エラー
func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
