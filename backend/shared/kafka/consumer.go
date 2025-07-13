package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/sydney-health-clone/backend/shared/logger"
	"go.uber.org/zap"
)

type MessageHandler func(ctx context.Context, message kafka.Message) error

type Consumer struct {
	reader  *kafka.Reader
	topic   string
	groupID string
	handler MessageHandler
}

func NewConsumer(brokers []string, topic, groupID string, handler MessageHandler) *Consumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       topic,
		GroupID:     groupID,
		MinBytes:    1,
		MaxBytes:    10e6, // 10MB
		StartOffset: kafka.LastOffset,
	})
	
	return &Consumer{
		reader:  reader,
		topic:   topic,
		groupID: groupID,
		handler: handler,
	}
}

func (c *Consumer) Start(ctx context.Context) error {
	logger.Info("Starting Kafka consumer",
		zap.String("topic", c.topic),
		zap.String("groupID", c.groupID),
	)
	
	for {
		select {
		case <-ctx.Done():
			logger.Info("Stopping Kafka consumer",
				zap.String("topic", c.topic),
				zap.String("groupID", c.groupID),
			)
			return ctx.Err()
		default:
			msg, err := c.reader.FetchMessage(ctx)
			if err != nil {
				if err == context.Canceled {
					return nil
				}
				logger.Error("Failed to fetch message",
					zap.String("topic", c.topic),
					zap.Error(err),
				)
				continue
			}
			
			// Process message
			if err := c.handler(ctx, msg); err != nil {
				logger.Error("Failed to handle message",
					zap.String("topic", c.topic),
					zap.String("key", string(msg.Key)),
					zap.Error(err),
				)
				// Continue processing other messages
			}
			
			// Commit message
			if err := c.reader.CommitMessages(ctx, msg); err != nil {
				logger.Error("Failed to commit message",
					zap.String("topic", c.topic),
					zap.Error(err),
				)
			}
		}
	}
}

func (c *Consumer) Close() error {
	return c.reader.Close()
}

// Helper functions for common message types

func UnmarshalClaimMessage(data []byte) (*ClaimMessage, error) {
	var msg ClaimMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal claim message: %w", err)
	}
	return &msg, nil
}

func UnmarshalMessageNotification(data []byte) (*MessageNotification, error) {
	var msg MessageNotification
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal message notification: %w", err)
	}
	return &msg, nil
}

func UnmarshalAuditEvent(data []byte) (*AuditEvent, error) {
	var msg AuditEvent
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal audit event: %w", err)
	}
	return &msg, nil
}

func UnmarshalMemberUpdate(data []byte) (*MemberUpdate, error) {
	var msg MemberUpdate
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal member update: %w", err)
	}
	return &msg, nil
}