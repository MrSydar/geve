package gevemongo

import (
	"context"
	"mrsydar/geve/schema"

	"go.mongodb.org/mongo-driver/bson"
)

// geveMongo is a MongoDB implementation of the geve database client interface.
type geveMongo struct {
	iClient
	collections map[string]iCollection
}

func (gm *geveMongo) ReadOne(collection, id string) (any, error) {
	var item any

	err := gm.collections[collection].findOne(context.TODO(), bson.M{"_id": id}).decode(&item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// Config is the configuration for GeveMongo.
type Config struct {
	Client       iClient
	Schemas      map[string]schema.Schema
	DatabaseName string
}

// New creates a new GeveMongo instance.
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
