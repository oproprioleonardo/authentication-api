package auth

import "go.mongodb.org/mongo-driver/bson/primitive"

type Search struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" query:"id"`
	UserProfileID primitive.ObjectID `bson:"userProfileId,omitempty" query:"profileId"`
	LastSessionID primitive.ObjectID `bson:"lastSessionId,omitempty" query:"lastSessionId"`
}
