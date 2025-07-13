package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/sydney-health-clone/backend/shared/logger"
	"go.uber.org/zap"
)

type Producer struct {
	writer *kafka.Writer
	topic  string
}

func NewProducer(brokers []string, topic string) *Producer {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireOne,
		Compression:  kafka.Snappy,
	}
	
	return &Producer{
		writer: writer,
		topic:  topic,
	}
}

func (p *Producer) SendMessage(ctx context.Context, key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}
	
	msg := kafka.Message{
		Key:   []byte(key),
		Value: data,
	}
	
	err = p.writer.WriteMessages(ctx, msg)
	if err != nil {
		logger.Error("Failed to send Kafka message",
			zap.String("topic", p.topic),
			zap.String("key", key),
			zap.Error(err),
		)
		return err
	}
	
	logger.Debug("Kafka message sent",
		zap.String("topic", p.topic),
		zap.String("key", key),
	)
	
	return nil
}

func (p *Producer) Close() error {
	return p.writer.Close()
}

// Message types for different topics

type ClaimMessage struct {
	ClaimID      string `json:"claim_id"`
	MemberID     string `json:"member_id"`
	Status       string `json:"status"`
	Action       string `json:"action"`
	Timestamp    int64  `json:"timestamp"`
}

type MessageNotification struct {
	ConversationID string `json:"conversation_id"`
	MessageID      string `json:"message_id"`
	MemberID       string `json:"member_id"`
	SenderName     string `json:"sender_name"`
	Preview        string `json:"preview"`
	Timestamp      int64  `json:"timestamp"`
}

type AuditEvent struct {
	EventType   string                 `json:"event_type"`
	EntityType  string                 `json:"entity_type"`
	EntityID    string                 `json:"entity_id"`
	UserID      string                 `json:"user_id"`
	IPAddress   string                 `json:"ip_address"`
	UserAgent   string                 `json:"user_agent"`
	EventData   map[string]interface{} `json:"event_data"`
	Timestamp   int64                  `json:"timestamp"`
}

type MemberUpdate struct {
	MemberID     string   `json:"member_id"`
	UpdateType   string   `json:"update_type"`
	UpdatedFields []string `json:"updated_fields"`
	Timestamp    int64    `json:"timestamp"`
}