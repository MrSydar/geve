package gevemongo

import (
	"context"
	"fmt"
	"mrsydar/geve/schema"

	"go.mongodb.org/mongo-driver/bson"
)

type GeveMongo struct {
	client
	collections map[string]collection
}

type Config struct {
	Schemas      map[string]schema.Schema
	DatabaseName string
}

func (gm *GeveMongo) ReadOne(collection, id string) (any, error) {
	var item any

	err := gm.collections[collection].FindOne(context.TODO(), bson.M{"_id": id}).Decode(&item)
	if err != nil {
		return nil, fmt.Errorf("failed to find item: %w", err)
	}

	return item, nil
}

// TODO: implement limit, default limit
func (gm *GeveMongo) ReadMany(collection string) ([]any, error) {
	items := []any{}

	cursor, err := gm.collections[collection].Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find items: %w", err)
	}

	defer cursor.Close(context.TODO())

	if err := cursor.All(context.TODO(), &items); err != nil {
		return nil, fmt.Errorf("failed to read items: %w", err)
	}

	return items, nil
}

func (gm *GeveMongo) Configure(c Config) {
	gm.collections = make(map[string]collection)

	db := gm.Database(c.DatabaseName)

	for name := range c.Schemas {
		gm.collections[name] = db.Collection(name)
	}
}
