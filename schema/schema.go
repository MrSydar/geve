package schema

import "fmt"

type geveType interface {
	verifyDefinition() error
	verifyValue(v any) error
}

type FieldName = string

type Schema map[FieldName]geveType

func (s Schema) Verify() error {
	for k, v := range s {
		if err := v.verifyDefinition(); err != nil {
			return fmt.Errorf("verification of field <%v> has failed: %w", k, err)
		}
	}
	return nil
}
