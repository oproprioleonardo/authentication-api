package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserProfile struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UUID       string             `bson:"uuid"`
	Name       string             `bson:"name"`
	OnlineMode bool               `bson:"onlineMode"`
	Email      string             `bson:"email"`
	Phone      string             `bson:"phone"`
	CreatedAt  int64              `bson:"createdAt"`
}
