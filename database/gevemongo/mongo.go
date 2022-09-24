package gevemongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoClient struct {
	*mongo.Client
}

func (c *mongoClient) database(name string, opts ...*options.DatabaseOptions) iDatabase {
	return nil
}

// MongoClient wraps a mongo.Client.
func MongoClient(c *mongo.Client) iClient {
	return &mongoClient{c}
}
