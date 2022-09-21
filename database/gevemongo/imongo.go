package gevemongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type client interface {
	Database(string, ...*options.DatabaseOptions) database
}

type database interface {
	Collection(string, ...*options.CollectionOptions) collection
}

type collection interface {
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) singleResult
	Find(context.Context, interface{}, ...*options.FindOptions) (cursor, error)
}

type singleResult interface {
	Decode(v interface{}) error
}

type cursor interface {
	All(context.Context, interface{}) error
	Close(context.Context) error
}
