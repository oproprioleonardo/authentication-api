package profile

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

func NewUserProfileMongoGateway(coll *mongo.Collection) *MongoGateway {
	return &MongoGateway{coll: coll}
}

func (r MongoGateway) Create(ctx context.Context, profile *entity.UserProfile) error {
	id, err := primitive.ObjectIDFromHex(profile.ID)
	if err != nil {
		return err
	}

	input := &db.UserProfile{
		ID:         id,
		UUID:       profile.UUID,
		Name:       profile.Name,
		OnlineMode: profile.OnlineMode,
		Email:      profile.Email,
		Phone:      profile.Phone,
		CreatedAt:  profile.CreatedAt,
	}

	if _, err := r.coll.InsertOne(ctx, input); err != nil {
		return err
	}
	return nil
}

func (r MongoGateway) FindBy(ctx context.Context, filter interface{}) (*entity.UserProfile, error) {
	var (
		profile *db.UserProfile
		result  *entity.UserProfile
		err     error
	)
	err = r.coll.FindOne(ctx, filter).Decode(&profile)
	if err != nil {
		return nil, err
	}

	result = &entity.UserProfile{
		ID:         profile.ID.Hex(),
		UUID:       profile.UUID,
		Name:       profile.Name,
		OnlineMode: profile.OnlineMode,
		Email:      profile.Email,
		Phone:      profile.Phone,
		CreatedAt:  profile.CreatedAt,
	}
	return result, nil
}

func (r MongoGateway) FindById(ctx context.Context, id string) (*entity.UserProfile, error) {
	identifier, _ := primitive.ObjectIDFromHex(id)
	return r.FindBy(ctx, bson.D{{Key: "_id", Value: identifier}})
}

func (r MongoGateway) FindByName(ctx context.Context, name string) (*entity.UserProfile, error) {
	return r.FindBy(ctx, bson.D{{Key: "name", Value: name}})
}

func (r MongoGateway) FindByUUID(ctx context.Context, uuid string) (*entity.UserProfile, error) {
	return r.FindBy(ctx, bson.D{{Key: "uuid", Value: uuid}})
}

func (r MongoGateway) FindAll(ctx context.Context, filter interface{}, page int64, perPage int64) ([]*entity.UserProfile, error) {
	var (
		results  []*entity.UserProfile
		profiles []*db.UserProfile
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

	err = cursor.All(ctx, &profiles)
	if err != nil {
		return results, err
	}

	results = make([]*entity.UserProfile, len(profiles))

	for i, profile := range profiles {
		results[i] = &entity.UserProfile{
			ID:         profile.ID.Hex(),
			UUID:       profile.UUID,
			Name:       profile.Name,
			OnlineMode: profile.OnlineMode,
			Email:      profile.Email,
			Phone:      profile.Phone,
			CreatedAt:  profile.CreatedAt,
		}
	}

	return results, nil
}

func (r MongoGateway) Update(ctx context.Context, profile *entity.UserProfile) error {
	attributes := bson.D{
		{Key: "uuid", Value: profile.UUID},
		{Key: "name", Value: profile.Name},
		{Key: "onlineMode", Value: profile.OnlineMode},
		{Key: "email", Value: profile.Email},
		{Key: "phone", Value: profile.Phone},
		{Key: "createdAt", Value: profile.CreatedAt},
	}

	update := bson.D{{Key: "$set", Value: attributes}}
	objId, _ := primitive.ObjectIDFromHex(profile.ID)

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
