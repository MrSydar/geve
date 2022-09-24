package gevemongo

import (
	"context"
	"fmt"
	"mrsydar/geve/schema"

	"go.mongodb.org/mongo-driver/bson"
)

type geveMongo struct {
	iClient
	collections map[string]iCollection
}

type Config struct {
	Client       iClient
	Schemas      map[string]schema.Schema
	DatabaseName string
}

func (gm *geveMongo) ReadOne(collection, id string) (any, error) {
	var item any

	err := gm.collections[collection].findOne(context.TODO(), bson.M{"_id": id}).decode(&item)
	if err != nil {
		return nil, fmt.Errorf("failed to find item: %w", err)
	}

	return item, nil
}

func New(c Config) (*geveMongo, error) {
	if c.Client == nil {
		return nil, ErrMongoClientRequired
	}

	gm := &geveMongo{
		c.Client,
		make(map[string]iCollection),
	}

	if c.DatabaseName == "" {
		return nil, ErrDatabaseNameRequired
	}

	db := gm.database(c.DatabaseName)
	for name := range c.Schemas {
		gm.collections[name] = db.collection(name)
	}

	return gm, nil
}
