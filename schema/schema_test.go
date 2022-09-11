package schema

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestEmptyCompiledSchema(t *testing.T) {
	schema := make(Schema)

	jsonSchema, err := schema.Compile()
	if err != nil {
		t.Fatalf("failed to compile schema: %s", err)
	}

	if jsonSchema["type"] != "object" {
		t.Fatalf("expected type to be object")
	}

	if jsonSchema["title"] != "mrsydar/jsonschema" {
		t.Fatalf("expected title to be mrsydar/jsonschema")
	}

	if len(jsonSchema["required"].([]string)) != 0 {
		t.Fatalf("expected required to be empty")
	}
}

func TestCompiledSchemaWithString(t *testing.T) {
	schema := Schema{
		"field_1": String{
			Common{
				Required: true,
				Nullable: true,
			},
		},
	}

	jsonSchema, err := schema.Compile()
	if err != nil {
		t.Fatalf("failed to compile schema: %s", err)
	}

	bsonType := jsonSchema["properties"].(map[string]map[string]any)["field_1"]["bsonType"].([]string)
	if len(bsonType) != 2 || !slices.Contains(bsonType, "string") || !slices.Contains(bsonType, "null") {
		t.Fatalf("expected bsonType to have 2 items (string and null), got %v", bsonType)
	}
}
