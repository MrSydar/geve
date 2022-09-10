package schema

import (
	"testing"
)

func TestSchemaVerificationWithStringDefinitionWithGoodBoundaries(t *testing.T) {
	s := Schema{
		"name": &String{
			MinLength: 5,
			MaxLength: 10,
		},
	}

	if err := s.Verify(); err != nil {
		t.Fatalf("error was not expected: %v", err)
	}
}

func TestSchemaVerificationWithStringDefinitionWithBadBoundaries(t *testing.T) {
	s := Schema{
		"name": &String{
			MinLength: 10,
			MaxLength: 5,
		},
	}

	if err := s.Verify(); err == nil {
		t.Fatalf("error was expected")
	}
}

// func TestSchemaVerifyDocument(t *testing.T) {
// 	s := Schema{
// 		"name": &String{
// 			MinLength: 5,
// 			MaxLength: 10,
// 		},
// 	}

// 	if err := s.VerifyDocument(map[string]interface{}{
// 		"name": "hello",
// 	}); err != nil {
// 		t.Fatalf("error was not expected: %v", err)
// 	}
// }
