package gevemongo

import (
	"context"
	"fmt"
	"mrsydar/geve/schema"

	"go.mongodb.org/mongo-driver/bson"
)

type geveMongo struct {
	IClient
	collections map[string]ICollection
}

type Config struct {
	MongoClient  IClient
	Schemas      map[string]schema.Schema
	DatabaseName string
}

func (gm *geveMongo) ReadOne(collection, id string) (any, error) {
	var item any

	err := gm.collections[collection].FindOne(context.TODO(), bson.M{"_id": id}).Decode(&item)
	if err != nil {
		return nil, fmt.Errorf("failed to find item: %w", err)
	}

	return item, nil
}

func NewClient(c Config) (*geveMongo, error) {
	if c.MongoClient == nil {
		return nil, ErrMongoClientRequired
	}

	gm := &geveMongo{
		IClient:     c.MongoClient,
		collections: make(map[string]ICollection),
	}

	if c.DatabaseName == "" {
		return nil, ErrDatabaseNameRequired
	}

	db := gm.Database(c.DatabaseName)
	for name := range c.Schemas {
		gm.collections[name] = db.Collection(name)
	}

	return gm, nil
}
