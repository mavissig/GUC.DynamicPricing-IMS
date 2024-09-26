package kafka_server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/mavissig/GUC.DynamicPricing-IMS/internal/ims/domain"
	"github.com/mavissig/GUC.DynamicPricing-IMS/internal/ims/transport"
	"log"
)

type DataUC interface {
	DataSet(data *domain.Data) error
}

type Server struct {
	cfg    *transport.Config
	dataUC DataUC
}

func New(cfg *transport.Config, dataUC DataUC) *Server {
	return &Server{
		cfg:    cfg,
		dataUC: dataUC,
	}
}

func (s *Server) Run(ctx context.Context) {
	if err := s.workerStore(
		ctx,
		"ims-store",
		[]string{
			"outData",
		},
	); err != nil {
		log.Println(err)
	}
}

func (s *Server) workerStore(ctx context.Context, consumerGroup string, topics []string) error {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     "kafka:9092",
		"broker.address.family": "v4",
		"group.id":              consumerGroup,
		"auto.offset.reset":     "earliest",
	})
	if err != nil {
		return errors.New(fmt.Sprintf("[IMS][TRANSPORT][KAFKA-SERVER][NEW] ERROR %s", err))
	}

	if err := consumer.SubscribeTopics(topics, nil); err != nil {
		return errors.New(fmt.Sprintf("[IMS][TRANSPORT][KAFKA-SERVER][WORKER-STORE] ERROR %s", err))
	}

	log.Println("[IMS][TRANSPORT][KAFKA-SERVER][WORKER-STORE] INFO worker start")

	for {
		select {
		case <-ctx.Done():
			log.Println("[IMS][TRANSPORT][KAFKA-SERVER][WORKER-STORE]: worker stopped")
			consumer.Close()
			return nil
		default:
			msg, err := consumer.ReadMessage(-1)
			if err != nil {
				log.Println("[IMS][TRANSPORT][KAFKA-SERVER][WORKER-STORE] ERROR ", err)
				continue
			}
			log.Printf("[IMS][TRANSPORT][KAFKA-SERVER][WORKER-STORE] INFO message received")

			data := domain.Data{}

			err = json.Unmarshal(msg.Value, &data)
			if err != nil {
				log.Println("[IMS][TRANSPORT][KAFKA-SERVER][WORKER-STORE] ERROR ", err)
				continue
			}

			log.Printf("[IMS][TRANSPORT][KAFKA-SERVER][WORKER-STORE] INFO message unmarshal: %v\n", data)

			err = s.dataUC.DataSet(&data)
			if err != nil {
				log.Println("[IMS][TRANSPORT][KAFKA-SERVER][WORKER-STORE] ERROR ", err)
				continue
			}
		}
	}
}
