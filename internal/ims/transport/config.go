package transport

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	HTTP  *HttpConfig
	KAFKA *KafkaConfig
}

type HttpConfig struct {
	Port string `envconfig:"HTTP_ROUTER_ADDRESS" required:"true"`
}

type KafkaConfig struct {
	BootstrapServers string `envconfig:"KAFKA_BOOTSTRAP_SERVERS" required:"true"`
}

func LoadConfig() *Config {
	for _, fileName := range []string{".env.local", ".env.transport"} {
		err := godotenv.Load(fileName)
		if err != nil {
			log.Println("[IMS][TRANSPORT][CONFIG] WARN ", err)
		}
	}

	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		log.Fatalln(err)
	}

	fmt.Println()
	fmt.Printf("[IMS]HTTP CONFIG: \n")
	fmt.Printf("Port:        %s\n", cfg.HTTP.Port)

	fmt.Println()
	fmt.Printf("[IMS]KAFKA CONFIG: \n")
	fmt.Printf("BootstrapServers:        %s\n", cfg.KAFKA.BootstrapServers)

	return cfg
}
