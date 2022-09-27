package main

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
)

type Config struct {
	User string `env:"USER"`
}

func main() {
	var cfg Config
	// допишите код здесь

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current user is %s\n", cfg.User)
}
