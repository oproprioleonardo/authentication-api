package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuthUser struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	UserProfileID     primitive.ObjectID `bson:"userProfileId,omitempty"`
	LastSessionID     primitive.ObjectID `bson:"lastSessionId,omitempty"`
	AuthByLastSession bool               `bson:"authByLastSession"`
	Password          string             `bson:"password"`
	Secret            string             `bson:"secret"`
}
