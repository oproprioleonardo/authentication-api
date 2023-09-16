package session

import (
	"context"
	"github.com/skyepic/privateapi/internal/domain/user/entity"
	"github.com/skyepic/privateapi/internal/infrastructure/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoGateway struct {
	coll *mongo.Collection
}

func NewUserSessionMongoGateway(coll *mongo.Collection) *MongoGateway {
	return &MongoGateway{coll: coll}
}

func (r MongoGateway) Create(ctx context.Context, session *entity.UserSession) error {
	id, err := primitive.ObjectIDFromHex(session.ID)
	if err != nil {
		return err
	}
	profileId, err := primitive.ObjectIDFromHex(session.ProfileId)
	if err != nil {
		return err
	}

	input := &db.UserSession{
		ID:          id,
		ProfileId:   profileId,
		Active:      session.Active,
		LastServer:  session.LastServer,
		Ip:          session.Ip,
		FinalizedAt: session.FinalizedAt,
		StartedAt:   session.StartedAt,
	}

	if _, err := r.coll.InsertOne(ctx, input); err != nil {
		return err
	}
	return nil
}

func (r MongoGateway) FindBy(ctx context.Context, filter interface{}) (*entity.UserSession, error) {
	var (
		session *db.UserSession
		result  *entity.UserSession
		err     error
	)
	err = r.coll.FindOne(ctx, filter).Decode(&session)
	if err != nil {
		return nil, err
	}

	result = &entity.UserSession{
		ID:          session.ID.Hex(),
		ProfileId:   session.ProfileId.Hex(),
		Active:      session.Active,
		LastServer:  session.LastServer,
		Ip:          session.Ip,
		FinalizedAt: session.FinalizedAt,
		StartedAt:   session.StartedAt,
	}
	return result, nil
}

func (r MongoGateway) FindById(ctx context.Context, id string) (*entity.UserSession, error) {
	identifier, _ := primitive.ObjectIDFromHex(id)
	return r.FindBy(ctx, bson.D{{Key: "_id", Value: identifier}})
}

func (r MongoGateway) FindAll(ctx context.Context, filter interface{}, page int64, perPage int64) ([]*entity.UserSession, error) {
	var (
		results  []*entity.UserSession
		sessions []*db.UserSession
		cursor   *mongo.Cursor
		err      error
	)
	skips := (page * perPage) - perPage
	opt := options.Find().SetLimit(perPage).SetSkip(skips)
	if filter != nil {
		cursor, err = r.coll.Find(ctx, filter, opt)
	} else {
		cursor, err = r.coll.Find(ctx, bson.D{}, opt)
	}

	if err != nil {
		return results, err
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Println(err)
		}
	}(cursor, ctx)

	err = cursor.All(ctx, &sessions)
	if err != nil {
		return results, err
	}

	results = make([]*entity.UserSession, len(sessions))

	for i, session := range sessions {
		results[i] = &entity.UserSession{
			ID:          session.ID.Hex(),
			ProfileId:   session.ProfileId.Hex(),
			Active:      session.Active,
			LastServer:  session.LastServer,
			Ip:          session.Ip,
			FinalizedAt: session.FinalizedAt,
			StartedAt:   session.StartedAt,
		}
	}

	return results, nil
}

func (r MongoGateway) Update(ctx context.Context, session *entity.UserSession) error {
	attributes := bson.D{
		{Key: "profileId", Value: session.ProfileId},
		{Key: "ip", Value: session.Ip},
		{Key: "lastServer", Value: session.LastServer},
		{Key: "active", Value: session.Active},
		{Key: "finalizedAt", Value: session.FinalizedAt},
		{Key: "startedAt", Value: session.StartedAt},
	}

	update := bson.D{{Key: "$set", Value: attributes}}
	objId, _ := primitive.ObjectIDFromHex(session.ID)

	if _, err := r.coll.UpdateByID(ctx, objId, update); err != nil {
		return err
	}

	return nil
}

func (r MongoGateway) DeleteBy(ctx context.Context, filter interface{}) error {
	if _, err := r.coll.DeleteMany(ctx, filter); err != nil {
		return err
	}
	return nil
}

func (r MongoGateway) DeleteById(ctx context.Context, id string) error {
	objId, _ := primitive.ObjectIDFromHex(id)
	return r.DeleteBy(ctx, bson.D{{Key: "_id", Value: objId}})
}
