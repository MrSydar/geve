package gevemongo

import (
	"errors"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClientMock struct {
}

func (mcm *MongoClientMock) Database(string, ...*options.DatabaseOptions) *mongo.Database {
	return nil
}

func TestConfigurationNoMongoClientProvided(t *testing.T) {
	gm, err := NewClient(Config{})

	if gm != nil {
		t.Errorf("expected nil")
	}

	if !errors.Is(err, ErrMongoClientRequired) {
		t.Errorf("expected error to be %v, got %v", ErrMongoClientRequired, err)
	}
}

func TestConfigurationNoDatabaseNameProvided(t *testing.T) {
	config := Config{
		MongoClient:  &MongoClientMock{},
		DatabaseName: "test",
	}

	gm, err := NewClient(config)

	if gm != nil {
		t.Errorf("expected nil")
	}

	if !errors.Is(err, ErrDatabaseNameRequired) {
		t.Errorf("expected error to be %v, got %v", ErrDatabaseNameRequired, err)
	}
}
