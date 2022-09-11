package schema

import (
	"encoding/json"
	"fmt"
)

type kindName = string
type Schema map[kindName]basic

func (s *Schema) ToJSONSchema() ([]byte, error) {
	jsonSchema := make(map[string]any)

	jsonSchema["type"] = "object"
	jsonSchema["title"] = "mrsydar/jsonschema"
	jsonSchema["required"] = []string{}

	for fieldName, fieldSchema := range *s {
		if fieldSchema.Required {
			jsonSchema["required"] = append(jsonSchema["required"].([]string), fieldName)
		}
	}

	jsonSchemaBytes, err := json.Marshal(jsonSchema)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json schema: %s", err)
	}

	return jsonSchemaBytes, nil
}
