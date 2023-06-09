package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Role     string             `bson:"role"`
}

type Friend struct {
	ID primitive.ObjectID `bson:"_id"`
	userA primitive.ObjectID  `bson:"userA"`
	userB primitive.ObjectID `bson:"userB"`
}
