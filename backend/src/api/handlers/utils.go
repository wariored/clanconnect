package handlers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func ParseToPrimitive(str string) primitive.ObjectID {
	parsedId, _ := primitive.ObjectIDFromHex(str)
	return parsedId
}
