package server

type Controller interface {
	// GetOne registers a new GET endpoint which runs the given callback
	// to retrieve a single item by collection name and item id.
	GetOne(readOne func(collection, id string) (any, error))

	// PostOne registers a new POST endpoint which runs the given callback
	// to insert a single item into the database.
	PostOne(insertOne func(collection string, object any) (any, error))

	// Start starts the server.
	Start() error
}
