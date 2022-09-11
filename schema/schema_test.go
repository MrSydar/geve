package schema

import (
	"encoding/json"
	"testing"
)

func TestEmptyCompiledSchema(t *testing.T) {
	schema := make(Schema)

	jsonSchema, err := schema.ToJSONSchema()
	if err != nil {
		t.Errorf("failed to compile schema: %s", err)
	}

	jsonSchemaUnmarhsaled := make(map[string]any)
	if err := json.Unmarshal(jsonSchema, &jsonSchemaUnmarhsaled); err != nil {
		t.Errorf("failed to unmarshal json schema: %s", err)
	}

	if jsonSchemaUnmarhsaled["type"] != "object" {
		t.Errorf("expected type to be object, got %s", jsonSchemaUnmarhsaled["type"])
	}

	if jsonSchemaUnmarhsaled["title"] != "mrsydar/jsonschema" {
		t.Errorf("expected title to be mrsydar/jsonschema, got %s", jsonSchemaUnmarhsaled["title"])
	}

	if len(jsonSchemaUnmarhsaled["required"].([]any)) != 0 {
		t.Errorf("expected required to be empty, got %s", jsonSchemaUnmarhsaled["required"])
	}
}
