package gevemongo

import (
	"github.com/pkg/errors"
)

var (
	ErrDatabaseNameRequired = errors.New("database name is required")
	ErrMongoClientRequired  = errors.New("mongo client is required")
	ErrFailedToFindItem     = errors.New("failed to find item")
)
