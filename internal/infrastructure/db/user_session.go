package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserSession struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	ProfileId   primitive.ObjectID `bson:"profileId,omitempty"`
	Active      bool               `bson:"active"`
	LastServer  string             `bson:"lastServer"`
	Ip          string             `bson:"lastIp"`
	FinalizedAt int64              `bson:"finalizedAt"`
	StartedAt   int64              `bson:"startedAt"`
}
