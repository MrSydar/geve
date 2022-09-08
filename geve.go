package geve

import (
	"fmt"
	"mrsydar/geve/schema"
)

type CollectionName = string

type geve struct {
	schemas map[CollectionName]schema.Schema
}

func (g *geve) RegisterSchema(collectionName string, schema schema.Schema) error {
	if err := schema.Verify(); err != nil {
		return fmt.Errorf("schema verification has failed: %w", err)
	}

	g.schemas[collectionName] = schema

	return nil
}

func (g *geve) Insert(collectionName string, document any) error {
	// TODO

	return nil
}

func New() *geve {
	return &geve{}
}
