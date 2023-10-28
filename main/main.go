package main

import (
	"plugin"

	_ "github.com/nyakiomaina/go-ds-http/httpdatastore"
)

var Plugins = []plugin.Plugin{
	&httpdatastore.HttpPlugin{},
}
