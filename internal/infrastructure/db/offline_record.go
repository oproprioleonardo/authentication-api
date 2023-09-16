package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type OfflineRecord struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UUID       string             `bson:"uuid"`
	Name       string             `bson:"name"`
	OnlineMode bool               `bson:"onlineMode"`
	Registered bool               `bson:"registered"`
}
