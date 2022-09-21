package gevemongo

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoCursorMock struct {
	allFuncCalled   bool
	closeFuncCalled bool
}

func (m *mongoCursorMock) All(ctx context.Context, v interface{}) error {
	m.allFuncCalled = true
	return nil
}

func (m *mongoCursorMock) Close(ctx context.Context) error {
	m.closeFuncCalled = true
	return nil
}

type mongoSingleResultMock struct {
	decodeFuncCalled bool
}

func (m *mongoSingleResultMock) Decode(v interface{}) error {
	m.decodeFuncCalled = true
	return nil
}

type mongoCollectionMock struct {
	findFuncCalled    bool
	findOneFuncCalled bool
}

func (m *mongoCollectionMock) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) singleResult {
	m.findOneFuncCalled = true
	return &mongoSingleResultMock{}
}

func (m *mongoCollectionMock) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cursor, error) {
	m.findFuncCalled = true
	return &mongoCursorMock{}, nil
}

type mongoDatabaseMock struct {
	collectionFuncCalled bool
}

func (m *mongoDatabaseMock) Collection(name string, opts ...*options.CollectionOptions) collection {
	m.collectionFuncCalled = true
	return &mongoCollectionMock{}
}

type mongoClientMock struct {
	databaseFuncCalled bool
}

func (m *mongoClientMock) Database(name string, opts ...*options.DatabaseOptions) database {
	if name != "test" {
		panic("wrong name")
	}

	m.databaseFuncCalled = true

	return &mongoDatabaseMock{}
}

func TestConfiguration(t *testing.T) {
	mcm := &mongoClientMock{}

	gm := GeveMongo{
		client: mcm,
	}

	gm.Configure(Config{
		Schemas:      nil,
		DatabaseName: "test",
	})

	if mcm.databaseFuncCalled != true {
		t.Error("database function not called")
	}
}
