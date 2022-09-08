package schema

import "testing"

func TestSchemaWithStringDefinition(t *testing.T) {
	s := Schema{
		"name": String{},
	}

	if err := s.Verify(); err != nil {
		t.Fatalf("error was not expected: %v", err)
	}
}
