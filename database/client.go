package database

type Client interface {
	// GetOne returns a single object from the database.
	ReadOne(collection, id string) (any, error)

	// GetMany returns multiple objects from the database.
	ReadMany(collection string) ([]any, error)
}
