package config

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/Kartochnik010/go-sso/utils"
	"github.com/ilyakaznacheev/cleanenv"
)

var (
	ErrPortUnset    = errors.New("port unset")
	ErrNoConfigPath = errors.New("no path to config file")
	ErrNoConfigFile = errors.New("no config file")
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-default:"1h"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}
type GRPCConfig struct {
	Port    int           `yaml:"port" env-default:"44044"`
	Timeout time.Duration `yaml:"timeout" env-default:"5s"`
}

func NewConfig() (*Config, error) {
	path, err := fetchConfigPath()
	if err != nil {
		return nil, err
	}

	if utils.FileNotExists(path) {
		return nil, ErrNoConfigFile
	}
	var cfg *Config
	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		return nil, fmt.Errorf("error parsing config: %s", err)
	}
	return cfg, nil
}

func fetchConfigPath() (string, error) {
	var res string
	flag.StringVar(&res, "config", "./config/config.yml", "path to config file")
	flag.Parse()
	if res == "" {
		return "", ErrNoConfigFile
	}
	return res, nil
}
