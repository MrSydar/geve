package geve

import (
	"mrsydar/geve/database"
	"mrsydar/geve/schema"
	"mrsydar/geve/server"
)

type geve struct {
	client     database.Client
	controller server.Controller
}

type Config struct {
	Schemas    map[string]schema.Schema
	Client     database.Client
	Controller server.Controller
}

func New(c Config) *geve {
	c.Controller.GetOne(c.Client.ReadOne)
	c.Controller.GetMany(c.Client.ReadMany)

	return &geve{
		client:     c.Client,
		controller: c.Controller,
	}
}
