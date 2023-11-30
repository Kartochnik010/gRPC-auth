package main

import (
	"log"

	"github.com/Kartochnik010/go-sso/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	_ = cfg
}
