package schema

type String struct {
	Common
}

func (s String) isRequired() bool {
	return s.Required
}

func (s String) compileProperties() map[string]any {
	bsonType := []string{"string"}

	if s.Nullable {
		bsonType = append(bsonType, "null")
	}

	return map[string]any{
		"bsonType": bsonType,
	}
}
