package schema

import "fmt"

type String struct {
	Basic
	MinLength, MaxLength uint64
}

func (ts String) verifyDefinition() error {
	if ts.MinLength > ts.MaxLength {
		return fmt.Errorf("min length (%v) is greater than max length (%v)", ts.MinLength, ts.MaxLength)
	}

	return nil
}

func (ts String) verifyValue(v any) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("value is not a string")
	}

	if len(s) < int(ts.MinLength) {
		return fmt.Errorf("value is shorter than min length (%v)", ts.MinLength)
	}

	if len(s) > int(ts.MaxLength) {
		return fmt.Errorf("value is longer than max length (%v)", ts.MaxLength)
	}

	return nil
}
