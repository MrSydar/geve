// Package schema implements a JSON Schema compiler.
package schema

// kind represents the type of a JSON Schema property.
type kind interface {
	isRequired() bool
	compileProperties() map[string]any
}

// Schema is used to configure a mongodb collection.
type Schema map[string]kind

// compile converts a Schema to a JSON Schema draft-04.
func (s *Schema) Compile() (map[string]any, error) {
	jsonSchema := make(map[string]any)

	jsonSchema["type"] = "object"
	jsonSchema["title"] = "mrsydar/jsonschema"
	jsonSchema["required"] = []string{}

	properties := make(map[string]map[string]any)
	for fieldName, fieldSchema := range *s {
		properties[fieldName] = fieldSchema.compileProperties()

		if fieldSchema.isRequired() {
			jsonSchema["required"] = append(jsonSchema["required"].([]string), fieldName)
		}
	}
	jsonSchema["properties"] = properties

	return jsonSchema, nil
}
