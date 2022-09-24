package gevemongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoClient struct {
	*mongo.Client
}

func (c *mongoClient) database(name string, opts ...*options.DatabaseOptions) iDatabase {
	return nil
type singleResult struct {
	*mongo.SingleResult
}

func (s *singleResult) decode(v interface{}) error {
	return s.SingleResult.Decode(v)
}

type collection struct {
	*mongo.Collection
}

func (c *collection) findOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) iSingleResult {
	return &singleResult{c.Collection.FindOne(ctx, filter, opts...)}
}

type database struct {
	*mongo.Database
}

func (d *database) collection(name string, opts ...*options.CollectionOptions) iCollection {
	return &collection{d.Database.Collection(name, opts...)}
}

type client struct {
	*mongo.Client
}

func (c *client) database(name string, opts ...*options.DatabaseOptions) iDatabase {
	return &database{c.Client.Database(name, opts...)}
}

// MongoClient wraps a mongo.Client.
func MongoClient(c *mongo.Client) iClient {
<<<<<<< HEAD
	return &mongoClient{c}
=======
	return &client{c}
>>>>>>> 87bb281 (Selected optimal direction for gevemongo configuration)
}
