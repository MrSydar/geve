package gevemongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IClient interface {
	Database(string, ...*options.DatabaseOptions) *mongo.Database
}

type IDatabase interface {
	Collection(string, ...*options.CollectionOptions) *mongo.Collection
}

type ICollection interface {
	FindOne(context.Context, interface{}, ...*options.FindOneOptions) *mongo.SingleResult
}

type ISingleResult interface {
	Decode(v interface{}) error
}
