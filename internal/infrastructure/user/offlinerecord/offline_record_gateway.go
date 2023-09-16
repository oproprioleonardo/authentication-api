package offlinerecord

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

func NewOfflineRecordMongoGateway(coll *mongo.Collection) *MongoGateway {
	return &MongoGateway{coll: coll}
}

func (r MongoGateway) Create(ctx context.Context, rec *entity.OfflineRecord) error {
	id, err := primitive.ObjectIDFromHex(rec.ID)
	if err != nil {
		return err
	}

	input := &db.OfflineRecord{
		ID:         id,
		UUID:       rec.UUID,
		Name:       rec.Name,
		OnlineMode: rec.OnlineMode,
		Registered: rec.Registered,
	}

	if _, err := r.coll.InsertOne(ctx, input); err != nil {
		return err
	}
	return nil
}

func (r MongoGateway) FindBy(ctx context.Context, filter interface{}) (*entity.OfflineRecord, error) {
	var (
		record *db.OfflineRecord
		result *entity.OfflineRecord
		err    error
	)
	err = r.coll.FindOne(ctx, filter).Decode(&record)
	if err != nil {
		return nil, err
	}

	result = &entity.OfflineRecord{
		ID:         record.ID.Hex(),
		UUID:       record.UUID,
		Name:       record.Name,
		OnlineMode: record.OnlineMode,
		Registered: record.Registered,
	}
	return result, nil
}

func (r MongoGateway) FindById(ctx context.Context, id string) (*entity.OfflineRecord, error) {
	identifier, _ := primitive.ObjectIDFromHex(id)
	return r.FindBy(ctx, bson.D{{Key: "_id", Value: identifier}})
}

func (r MongoGateway) FindByUUID(ctx context.Context, uuid string) (*entity.OfflineRecord, error) {
	return r.FindBy(ctx, bson.D{{Key: "uuid", Value: uuid}})
}

func (r MongoGateway) FindByName(ctx context.Context, name string) (*entity.OfflineRecord, error) {
	return r.FindBy(ctx, bson.D{{Key: "name", Value: name}})
}

func (r MongoGateway) FindAll(ctx context.Context, filter interface{}, page int64, perPage int64) ([]*entity.OfflineRecord, error) {
	var (
		results []*entity.OfflineRecord
		records []*db.OfflineRecord
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

	err = cursor.All(ctx, &records)
	if err != nil {
		return results, err
	}

	for _, rec := range records {
		results = append(results, &entity.OfflineRecord{
			ID:         rec.ID.Hex(),
			UUID:       rec.UUID,
			Name:       rec.Name,
			OnlineMode: rec.OnlineMode,
			Registered: rec.Registered,
		})
	}

	return results, nil
}

func (r MongoGateway) Update(ctx context.Context, rec *entity.OfflineRecord) error {
	attributes := bson.D{
		{Key: "uuid", Value: rec.UUID},
		{Key: "name", Value: rec.Name},
		{Key: "onlineMode", Value: rec.OnlineMode},
		{Key: "registered", Value: rec.Registered},
	}

	update := bson.D{{Key: "$set", Value: attributes}}
	objId, _ := primitive.ObjectIDFromHex(rec.ID)

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
