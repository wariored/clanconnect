package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID        primitive.ObjectID `bson:"_id"`
	Text      string             `bson:"text"`
	SenderID    *primitive.ObjectID             `bson:"sender"`
	RecipientID  *primitive.ObjectID             `bson:"receiver"`
	Timestamp time.Time          `bson:"timestamp"`
}

