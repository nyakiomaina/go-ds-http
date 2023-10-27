package httpdatastore

type HttpConfig struct {
	ServerURL string
}

func (cfg *HttpConfig) DiskSpec() map[string]interface{} {
	return map[string]interface{}{
		"type":      "httpDatastore",
		"serverURL": cfg.ServerURL,
	}
}

func (cfg *HttpConfig) ConfigFromMap(m map[string]interface{}) error {
	// "serverURL"??
	if url, ok := m["serverURL"]; ok {
		cfg.ServerURL = url.(string)
	}
	return nil
}
