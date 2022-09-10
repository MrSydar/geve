package schema

import "fmt"

type kind interface {
	verifyDefinition() error
	verifyValue(any) error
}

type Schema map[string]kind

func (s *Schema) Verify() error {
	for k, v := range *s {
		if err := v.verifyDefinition(); err != nil {
			return fmt.Errorf("error in definition of %v: %w", k, err)
		}
	}

	return nil
}
