package offlinerecord

import "go.mongodb.org/mongo-driver/bson/primitive"

type Search struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" query:"id,omitempty"`
	UUID string             `bson:"uuid,omitempty" query:"uuid,omitempty"`
	Name string             `bson:"name,omitempty" query:"name,omitempty"`
}
