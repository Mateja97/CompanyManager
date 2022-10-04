package kafka

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
)

type KafkaProducer struct {
	kafkaTopic string
	producer   sarama.SyncProducer
}
type KafkaMessage struct {
	Operation string
	ID        string
	Data      interface{}
}

func (kP *KafkaProducer) Init(kafkaBrokers []string, kafkaTopic string) error {
	kP.kafkaTopic = kafkaTopic
	var err error

	kP.producer, err = sarama.NewSyncProducer(kafkaBrokers, nil)
	if err != nil {
		return err
	}
	return nil
}

func (kP *KafkaProducer) SendMessage(event interface{}) error {
	json, err := json.Marshal(event)
	if err != nil {
		log.Println("[ERROR] Marshaling event failed")
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: kP.kafkaTopic,
		Key:   nil,
		Value: sarama.StringEncoder(json),
	}
	_, _, err = kP.producer.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}
func (kC *KafkaProducer) Stop() error {
	log.Println("Producer closed stopped")
	if err := kC.producer.Close(); err != nil {
		log.Println("[ERROR] Producer close failed")
		return err
	}
	return nil
}
