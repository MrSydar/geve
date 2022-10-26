package database

// Client is a database client.
type Client interface {
	// GetOne returns a single object from the database.
	ReadOne(collection, id string) (any, error)

	// InsertOne inserts a single object into the database and returns the inserted object filled with generated values.
	InsertOne(collection string, object any) (any, error)
}
