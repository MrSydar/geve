package schema

import "fmt"

type String struct {
	Basic
	MinLength, MaxLength uint64
}

func (ts String) verifyDefinition() error {
	if ts.MinLength > ts.MaxLength {
		return fmt.Errorf("value of field MaxLength cannot be greater than value of field MaxLength")
	}

	return nil
}

// TODO how to make it verifyValue(string) ?
func (ts String) verifyValue(v any) error {
	strValue, ok := v.(string)

	if !ok {
		return fmt.Errorf("passed value cannot be asserted to a string type")
	}

	if length := uint64(len(strValue)); length < ts.MinLength || length > ts.MaxLength {
		return fmt.Errorf("length doesn't fit the defined boundaries")
	}

	return nil
}
