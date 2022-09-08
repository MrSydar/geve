package schema

import "testing"

func TestGoodBoundaries(t *testing.T) {
	ts := String{
		MinLength: 10,
		MaxLength: 20,
	}

	if err := ts.verifyDefinition(); err != nil {
		t.Fatalf("error was not expected: %v", err)
	}
}
