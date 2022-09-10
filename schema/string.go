package schema

import "fmt"

type String struct {
	Basic
	MinLength, MaxLength uint64
}

func (ks String) verifyDefinition() error {
	if ks.MinLength > ks.MaxLength {
		return fmt.Errorf("min length (%v) is greater than max length (%v)", ks.MinLength, ks.MaxLength)
	}

	return nil
}

func (ks String) verifyValue(v any) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("value is not a string")
	}

	if len(s) < int(ks.MinLength) {
		return fmt.Errorf("value is shorter than min length (%v)", ks.MinLength)
	}

	if len(s) > int(ks.MaxLength) {
		return fmt.Errorf("value is longer than max length (%v)", ks.MaxLength)
	}

	return nil
}
