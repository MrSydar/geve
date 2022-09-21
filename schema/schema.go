package schema

type (
	// kind represents a field type.
	Kind interface{}

	// FieldName is	 a field name.
	FieldName = string

	// Schema represents a schema definition.
	Schema map[FieldName]Kind
)
