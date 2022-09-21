package server

type Controller interface {
	// GetOne registers a new GET endpoint which runs the given callback
	// to retrieve a single item by collection name and item id.
	GetOne(readOne func(collection, id string) (any, error))

	// GetMany registers a new GET endpoint which runs the given callback
	// to retrieve multiple items by collection name.
	GetMany(readMany func(collection string) ([]any, error))
}
