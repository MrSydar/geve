package gevemongo

import (
	"context"
	"errors"
	"mrsydar/geve/schema"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoSingleResultMock struct {
	obj any
}

func (msrm *mongoSingleResultMock) decode(obj interface{}) error {
	if msrm.obj == nil {
		return mongo.ErrNoDocuments
	}

	srcPtrValue := reflect.ValueOf(msrm.obj)

	dstPtrValue := reflect.ValueOf(obj)
	dstValue := reflect.Indirect(dstPtrValue)

	dstValue.Set(srcPtrValue)

	return nil
}

type mongoCollectionMock struct {
	storage map[string]any
}

func (mcm *mongoCollectionMock) findOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) iSingleResult {
	bsonFilter := filter.(bson.M)
	id := bsonFilter["_id"].(string)

	if mcm.storage[id] == nil {
		return &mongoSingleResultMock{nil}
	}

	return &mongoSingleResultMock{mcm.storage[id]}
}

type mongoDatabaseMock struct {
	coll iCollection
}

func (mdm *mongoDatabaseMock) collection(string, ...*options.CollectionOptions) iCollection {
	return mdm.coll
}

type mongoClientMock struct {
	db iDatabase
}

func (mcm *mongoClientMock) database(string, ...*options.DatabaseOptions) iDatabase {
	return mcm.db
}

func TestNewWithConfigurationWithoutNoMongoClient(t *testing.T) {
	config := Config{}

	gm, err := New(config)

	if gm != nil {
		t.Errorf("expected nil")
	}

	if !errors.Is(err, ErrMongoClientRequired) {
		t.Errorf("expected error to be %v, got %v", ErrMongoClientRequired, err)
	}
}

func TestNewWithConfigurationWithoutDatabaseName(t *testing.T) {
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

func TestNewWithConfigurationWithoutSchema(t *testing.T) {
	config := Config{
		Client:       &mongoClientMock{},
		DatabaseName: "items",
	}

	gm, err := New(config)

	if gm == nil {
		t.Fatalf("expected not nil")
	}

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if gm.collections == nil {
		t.Errorf("expected not nil")
	}

	if len(gm.collections) != 0 {
		t.Errorf("expected 0, got %v", len(gm.collections))
	}
}

func TestNewConfigurationWithSchema(t *testing.T) {
	config := Config{
		Client: &mongoClientMock{
			db: &mongoDatabaseMock{},
		},
		Schemas: map[string]schema.Schema{
			"items": {
				"_id":      schema.String{},
				"quantity": schema.Integer{},
			},
		},
		DatabaseName: "test-db",
	}

	gm, err := New(config)

	if gm == nil {
		t.Fatalf("expected not nil")
	}

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if gm.collections == nil {
		t.Errorf("expected not nil")
	}

	if len(gm.collections) != 1 {
		t.Errorf("expected 1")
	}

	if _, ok := gm.collections["items"]; !ok {
		t.Errorf("expected key to exist")
	}
}

func TestReadOneNoDocumentFound(t *testing.T) {
	config := Config{
		Client: &mongoClientMock{
			db: &mongoDatabaseMock{
				coll: &mongoCollectionMock{
					storage: map[string]any{},
				},
			},
		},
		Schemas: map[string]schema.Schema{
			"items": {
				"_id":      schema.String{},
				"quantity": schema.Integer{},
			},
		},
		DatabaseName: "test-db",
	}

	gm, _ := New(config)

	obj, err := gm.ReadOne("items", "123")

	if obj != nil {
		t.Errorf("expected nil")
	}

	if !errors.Is(err, mongo.ErrNoDocuments) {
		t.Errorf("expected error to be %v, got %v", mongo.ErrNoDocuments, err)
	}
}

func TestReadOneDocumentFound(t *testing.T) {
	type item struct {
		ID       string `bson:"_id"`
		Quantity int    `bson:"quantity"`
	}

	expectedItem := item{"123", 10}

	config := Config{
		Client: &mongoClientMock{
			db: &mongoDatabaseMock{
				coll: &mongoCollectionMock{
					storage: map[string]any{
						"123": expectedItem,
					},
				},
			},
		},
		Schemas: map[string]schema.Schema{
			"items": {
				"_id":      schema.String{},
				"quantity": schema.Integer{},
			},
		},
		DatabaseName: "test-db",
	}

	gm, _ := New(config)

	actualObj, err := gm.ReadOne("items", expectedItem.ID)

	if actualObj == nil {
		t.Errorf("expected not nil")
	}

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if actualObj.(item) != expectedItem {
		t.Errorf("expected %v, got %v", expectedItem, actualObj)
	}
}

func ExampleNew() {
	mc, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	config := Config{
		Client:       MongoClient(mc),
		DatabaseName: "test-db",
	}

	_, _ = New(config)
}

func ExampleMongoClient() {
	mc, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	_ = Config{
		Client:       MongoClient(mc),
		DatabaseName: "test-db",
	}
}
