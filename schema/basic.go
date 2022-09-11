package schema

// Common is a set of common configuration options for all types of fields.
type Common struct {
	// Required is a flag that indicates whether the field is mandatory on insertion.
	Required bool

	// Nullable is a flag that indicates whether the field can be nil.
	Nullable bool
}
