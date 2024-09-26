package repository

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
}

func LoadConfig() *Config {
	for _, fileName := range []string{".env.local", ".env.storage"} {
		err := godotenv.Load(fileName)
		if err != nil {
			log.Println("[IMS][INFRASTRUCTURE][CONFIG] WARN ", err)
		}
	}

	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("[IMS]POSTGRES CONFIG: \n")

	fmt.Println()

	return cfg
}
