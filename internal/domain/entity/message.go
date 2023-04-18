package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        string
	Role      string
	Content   string
	Tokens    int
	Model     *Model
	CreatedAt time.Time
}

func NewMessage(role, content string, model *Model) (*Message, error) {
	msg := &Message{
		ID:        uuid.New().String(),
		Role:      role,
		Content:   content,
		Model:     model,
		CreatedAt: time.Now(),
	}
	return msg, nil
}

func (m *Message) Validate() error {

	if m.Role != "user" && m.Role != "System" && m.Role != "assistant" {
		return errors.New("invalid role")
	}
	return nil
}
