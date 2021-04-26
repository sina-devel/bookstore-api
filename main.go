package main

import (
	"github.com/kianooshaz/bookstore-api/internal/config"
	"log"
)

func init() {
	var cfg *config.Config

	if err := config.ReadFile(cfg); err != nil {
		log.Fatalln(err)
	}

	if err := config.ReadEnv(cfg); err != nil {
		log.Fatalln(err)
	}

	config.SetConfig(cfg)
}

func main() {

}
