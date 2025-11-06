package infra

import (
	"context"
	"cqrs-base/internal/domain"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

type UserConsumer struct {
	reader *kafka.Reader
	db     *sql.DB
}

func NewUserConsumer(reader *kafka.Reader, db *sql.DB) *UserConsumer {
	return &UserConsumer{reader, db}
}

func (c *UserConsumer) Start(ctx context.Context) error {
	log.Println("🚀 Starting user consumer...")
	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Println("Kafka read error:", err)
			continue
		}

		var event domain.UserEvent
		if err := json.Unmarshal(m.Value, &event); err != nil {
			log.Println("JSON decode error:", err)
			continue
		}

		switch event.EventType {
		case "UserCreated":
			_, err = c.db.ExecContext(ctx,
				`INSERT INTO users_read (id, name, email, created_at, updated_at)
				 VALUES ($1, $2, $3, $4, $5)
				 ON CONFLICT (id) DO NOTHING`,
				event.UserID, event.Name, event.Email, event.Time, event.Time)

		case "UserUpdated":
			_, err = c.db.ExecContext(ctx,
				`UPDATE users_read SET name=$1, email=$2, updated_at=$3 WHERE id=$4`,
				event.Name, event.Email, event.Time, event.UserID)

		case "UserDeleted":
			_, err = c.db.ExecContext(ctx, `DELETE FROM users_read WHERE id=$1`, event.UserID)
		}

		if err != nil {
			log.Println("DB error:", err)
		}
	}
}
