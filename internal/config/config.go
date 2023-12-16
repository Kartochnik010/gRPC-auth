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
	ErrNoConfigPath = errors.New("no path to config file")
	ErrNoConfigFile = errors.New("no config file")
)

type Config struct {
	Env         string        `yaml:"env" env-required:"true"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"grpc" env-required:"true"`
}
type GRPCConfig struct {
	Port    int           `yaml:"port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
}

// env: "local"
// storage_path: "./storage/sso.db"
// token_ttl: 5m
// grpc:
//
//	port: 44044
//	timeout: 5s
func NewConfig() (*Config, error) {
	path, err := fetchConfigPath()
	if err != nil {
		return nil, err
	}

	if utils.FileNotExists(path) {
		return nil, ErrNoConfigFile
	}
	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, fmt.Errorf("error parsing config: %s", err)
	}
	return &cfg, nil
}

func MustLoadByPath(path string) *Config {
	if utils.FileNotExists(path) {
		panic("config file not found. create config file as ./config/config.yml from root directory of the project")
	}
	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("error parsing config: " + err.Error())
	}
	return &cfg
}

func MustLoad() *Config {
	path, err := fetchConfigPath()
	if err != nil {
		panic("config file not found")
	}
	return MustLoadByPath(path)
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
