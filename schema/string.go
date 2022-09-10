package schema

type String struct {
	basic

	MinLength, MaxLength uint64
}
