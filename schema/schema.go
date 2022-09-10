package schema

import "fmt"

type kind interface {
	verifyDefinition() error
	verifyValue(any) error
}

type Schema map[string]kind

func (s Schema) Verify() error {
	for k, v := range s {
		if err := v.verifyDefinition(); err != nil {
			return fmt.Errorf("error in definition of %v: %w", k, err)
		}
	}

	return nil
}

func (s Schema) Fit(document map[string]any) error {
	for k, v := range s {
		if err := v.verifyValue(document[k]); err != nil {
			return fmt.Errorf("error in value of %v: %w", k, err)
		}
	}

	return nil
}
