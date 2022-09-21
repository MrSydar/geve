package geve

import "testing"

type MockObject struct {
	MockField string `json:"mock_field"`
}

type MockClient struct{}

func (m *MockClient) ReadOne(collection, id string) (any, error) {
	expectedCollection := "test_collection"
	if collection != expectedCollection {
		panic("expected collection to be " + expectedCollection + " but got " + collection)
	}

	expectedId := "test_id"
	if id != expectedId {
		panic("expected id to be " + expectedId + " but got " + id)
	}

	return MockObject{"test_field"}, nil
}

func (m *MockClient) ReadMany(collection string) ([]any, error) {
	expectedCollection := "test_collection"
	if collection != expectedCollection {
		panic("expected collection to be " + expectedCollection + " but got " + collection)
	}

	mo := MockObject{"test_field"}
	mos := []any{mo}

	return mos, nil
}

type MockController struct{}

func (m *MockController) GetOne(readOne func(collection, id string) (any, error)) {
	obj, err := readOne("test_collection", "test_id")
	if err != nil {
		panic(err)
	}

	objMock, ok := obj.(MockObject)
	if !ok {
		panic("expected obj to be of type MockObject")
	}

	expectedMockField := "test_field"
	if objMock.MockField != expectedMockField {
		panic("expected mock field to be " + expectedMockField + " but got " + objMock.MockField)
	}
}

func (m *MockController) GetMany(readMany func(collection string) ([]any, error)) {
	readMany("test_collection")
}

func Test(t *testing.T) {
	c := Config{
		Schemas:    nil,
		Client:     &MockClient{},
		Controller: &MockController{},
	}

	New(c)
}
