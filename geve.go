package geve

import (
	"mrsydar/geve/database"
	"mrsydar/geve/server"
)

type geve struct {
	client     database.Client
	controller server.Controller
}

func (g *geve) Start() error {
	return g.controller.Start()
}

type Config struct {
	Client     database.Client
	Controller server.Controller
}

func New(c Config) *geve {
	c.Controller.GetOne(c.Client.ReadOne)
	c.Controller.PostOne(c.Client.InsertOne)

	return &geve{
		client:     c.Client,
		controller: c.Controller,
	}
}
