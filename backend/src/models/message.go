package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Text      string             `bson:"text" json:"text"`
	SenderID    primitive.ObjectID `bson:"sender" json:"sender_id"`
	RecipientID  primitive.ObjectID  `bson:"receiver" json:"recipient_id"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
}
