package gevemongo

import (
	"errors"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoClientMock struct {
}

func (mcm *mongoClientMock) database(string, ...*options.DatabaseOptions) iDatabase {
	return nil
}

func TestConfigurationNoMongoClientProvided(t *testing.T) {
	config := Config{}

	gm, err := New(config)

	if gm != nil {
		t.Errorf("expected nil")
	}

	if !errors.Is(err, ErrMongoClientRequired) {
		t.Errorf("expected error to be %v, got %v", ErrMongoClientRequired, err)
	}
}

func TestConfigurationNoDatabaseNameProvided(t *testing.T) {
	config := Config{
		Client: &mongoClientMock{},
	}

	gm, err := New(config)

	if gm != nil {
		t.Errorf("expected nil")
	}

	if !errors.Is(err, ErrDatabaseNameRequired) {
		t.Errorf("expected error to be %v, got %v", ErrDatabaseNameRequired, err)
	}
}

func ExampleMongoClient() {
	mc, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	config := Config{
		Client:       MongoClient(mc),
		DatabaseName: "items",
	}

	New(config)
}
