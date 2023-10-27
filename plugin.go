package httpdatastore

import (
	"github.com/ipfs/kubo/plugin"
	"github.com/ipfs/kubo/repo/fsrepo"
)

type HttpPlugin struct{}

func (p *HttpPlugin) Name() string {
	return "ds-http"
}

func (p *HttpPlugin) Version() string {
	return "0.1.0"
}

func (p *HttpPlugin) Init(env *plugin.Environment) error {
	return nil
}

func (p *HttpPlugin) DatastoreTypeName() string {
	return "ds-http"
}

func (p *HttpPlugin) DatastoreConfigParser() fsrepo.ConfigFromMap {
	return func(m map[string]interface{}) (interface{}, error) {
		cfg := &HttpConfig{}
		err := cfg.ConfigFromMap(m)
		return cfg, err
	}
}

var Plugins = []plugin.PluginDatastore{
	&HttpPlugin{},
}
