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

func TestBadBoundaries(t *testing.T) {
	ts := String{
		MinLength: 10,
		MaxLength: 5,
	}

	if err := ts.verifyDefinition(); err == nil {
		t.Fatalf("error was expected")
	}
}

func TestVerifyValueStringFitsLength(t *testing.T) {
	ts := String{
		MinLength: 5,
		MaxLength: 10,
	}

	if err := ts.verifyValue("hello"); err != nil {
		t.Fatalf("error was not expected: %v", err)
	}
}

func TestVerifyValueStringDoesNotFitLength(t *testing.T) {
	ts := String{
		MinLength: 5,
		MaxLength: 10,
	}

	if err := ts.verifyValue("hello world"); err == nil {
		t.Fatalf("error was expected")
	}
}

func TestVerifyValueInteger(t *testing.T) {
	ts := String{
		MinLength: 10,
		MaxLength: 20,
	}

	if err := ts.verifyValue(10); err == nil {
		t.Fatalf("error was expected")
	}
}
