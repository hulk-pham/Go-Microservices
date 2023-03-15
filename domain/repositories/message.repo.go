package repositories

import (
	"errors"
	"hulk/go-webservice/domain/entities"
	"hulk/go-webservice/infrastructure/persist"
)

type MessageRepository struct{}

func (m *MessageRepository) GetMessages() []entities.Message {
	var messages []entities.Message
	persist.DB.Find(&messages)
	return messages
}

func (m *MessageRepository) GetMessage(messageId string) (message entities.Message, err error) {
	r := persist.DB.First(&message, messageId)
	if r.RowsAffected > 0 {
		err = errors.New("Not found")
		return
	}
	return
}
