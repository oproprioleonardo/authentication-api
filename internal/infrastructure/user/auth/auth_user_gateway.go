package auth

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

type (
	MongoGateway struct {
		coll *mongo.Collection
	}
)

func NewAuthUserMongoGateway(coll *mongo.Collection) *MongoGateway {
	return &MongoGateway{coll: coll}
}

func (r MongoGateway) Create(ctx context.Context, user *entity.AuthUser) error {
	id, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return err
	}
	profileId, err := primitive.ObjectIDFromHex(user.UserProfileID)
	if err != nil {
		return err
	}
	lastSessionId, err := primitive.ObjectIDFromHex(user.LastSessionID)
	if err != nil {
		lastSessionId = primitive.NilObjectID
	}

	input := &db.AuthUser{
		ID:                id,
		UserProfileID:     profileId,
		LastSessionID:     lastSessionId,
		AuthByLastSession: user.AuthByLastSession,
		Password:          user.Password,
		Secret:            user.Secret,
	}

	if _, err := r.coll.InsertOne(ctx, input); err != nil {
		return err
	}
	return nil
}

func (r MongoGateway) FindBy(ctx context.Context, filter interface{}) (*entity.AuthUser, error) {
	var (
		user   *db.AuthUser
		result *entity.AuthUser
		err    error
	)
	err = r.coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	result = &entity.AuthUser{
		ID:                user.ID.Hex(),
		UserProfileID:     user.UserProfileID.Hex(),
		LastSessionID:     user.LastSessionID.Hex(),
		AuthByLastSession: user.AuthByLastSession,
		Password:          user.Password,
		Secret:            user.Secret,
	}
	return result, nil
}

func (r MongoGateway) FindById(ctx context.Context, id string) (*entity.AuthUser, error) {
	identifier, _ := primitive.ObjectIDFromHex(id)
	return r.FindBy(ctx, bson.D{{Key: "_id", Value: identifier}})
}

func (r MongoGateway) FindByProfileId(ctx context.Context, id string) (*entity.AuthUser, error) {
	identifier, _ := primitive.ObjectIDFromHex(id)
	return r.FindBy(ctx, bson.D{{Key: "userProfileId", Value: identifier}})
}

func (r MongoGateway) FindAll(ctx context.Context, filter interface{}, page int64, perPage int64) ([]*entity.AuthUser, error) {
	var (
		results []*entity.AuthUser
		users   []*db.AuthUser
		cursor  *mongo.Cursor
		err     error
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

	err = cursor.All(ctx, &users)
	if err != nil {
		return results, err
	}

	for _, user := range users {
		results = append(results, &entity.AuthUser{
			ID:                user.ID.Hex(),
			UserProfileID:     user.UserProfileID.Hex(),
			LastSessionID:     user.LastSessionID.Hex(),
			AuthByLastSession: user.AuthByLastSession,
			Password:          user.Password,
			Secret:            user.Secret,
		})
	}

	return results, nil
}

func (r MongoGateway) Update(ctx context.Context, user *entity.AuthUser) error {
	attributes := bson.D{
		{Key: "userProfileId", Value: user.UserProfileID},
		{Key: "lastSessionId", Value: user.LastSessionID},
		{Key: "authByLastSession", Value: user.AuthByLastSession},
		{Key: "password", Value: user.Password},
		{Key: "secret", Value: user.Secret},
	}

	update := bson.D{{Key: "$set", Value: attributes}}
	objId, _ := primitive.ObjectIDFromHex(user.ID)

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

func (r MongoGateway) DeleteByUUID(ctx context.Context, uuid string) error {
	return r.DeleteBy(ctx, bson.D{{Key: "uuid", Value: uuid}})
}
