package gevemongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type iClient interface {
	database(string, ...*options.DatabaseOptions) iDatabase
}

type iDatabase interface {
	collection(string, ...*options.CollectionOptions) iCollection
}

type iCollection interface {
	findOne(context.Context, interface{}, ...*options.FindOneOptions) iSingleResult
}

type iSingleResult interface {
	decode(v interface{}) error
}
