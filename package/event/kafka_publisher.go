package infra

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaEventPublisher struct {
	writer *kafka.Writer
}

// ✅ Constructor
func NewKafkaEventPublisher(brokers []string, topic string) *KafkaEventPublisher {
	return &KafkaEventPublisher{
		writer: &kafka.Writer{
			Addr:         kafka.TCP(brokers...),
			Topic:        topic,
			Balancer:     &kafka.LeastBytes{},
			RequiredAcks: kafka.RequireAll,
		},
	}
}

// ✅ Method PublishEvent
func (p *KafkaEventPublisher) PublishEvent(ctx context.Context, eventType string, payload interface{}) error {
	message := map[string]interface{}{
		"event_type": eventType,
		"timestamp":  time.Now().UTC(),
		"payload":    payload,
	}

	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = p.writer.WriteMessages(ctx, kafka.Message{
		Value: bytes,
	})
	if err != nil {
		log.Printf("[Kafka] Failed to publish event %s: %v", eventType, err)
		return err
	}

	log.Printf("[Kafka] Event %s published successfully ✅", eventType)
	return nil
}

// ✅ Jangan lupa tutup writer saat shutdown
func (p *KafkaEventPublisher) Close() error {
	return p.writer.Close()
}
