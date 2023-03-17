package services

import (
	"context"
	"time"
	"wrapup/database"
	"wrapup/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageService interface {
	CreateMessage(message *models.Message) error
	GetMessagesByRecipient(recipientID primitive.ObjectID) ([]*models.Message, error)
	DeleteMessage(messageID primitive.ObjectID) error
}

type MessageServiceImpl struct {
	Db *database.Client
}

func (ms *MessageServiceImpl) CreateMessage(message *models.Message) error {
	ctx := context.Background()
	message.ID = primitive.NewObjectID()
	message.Timestamp = time.Now()

	_, err := ms.Db.Database("wrapup-base").Collection("messages").InsertOne(ctx, message)
	return err
}

func (ms *MessageServiceImpl) GetMessagesByRecipient(recipientID primitive.ObjectID) ([]*models.Message, error) {
	ctx := context.Background()
	messages, err := ms.Db.Database("wrapup-base").Collection("messages").Find(ctx, bson.M{"recipientID": recipientID})
	if err != nil {
		return nil, err
	}
	defer messages.Close(ctx)

	var result []*models.Message
	for messages.Next(ctx) {
		var message models.Message
		if err := messages.Decode(&message); err != nil {
			return nil, err
		}
		result = append(result, &message)
	}

	return result, nil
}

func (ms *MessageServiceImpl) DeleteMessage(messageID primitive.ObjectID) error {
	ctx := context.Background()
	_, err := ms.Db.Database("wrapup-base").Collection("messages").DeleteOne(ctx, bson.M{"_id": messageID})
	return err
}
