package stretch

type FullClusterSettings struct {
	Persistent map[string]string `json:"persistent,omitempty"`
	Transient  map[string]string `json:"transient,omitempty"`
}

func (c *Cluster) GetSettings() (FullClusterSettings, error) {
	var data FullClusterSettings
	err := c.Client.Get(&data, "/_cluster/settings")
	return data, err
}

func (c *Cluster) SetSettings(settings interface{}) error {
	err := c.Client.Put(nil, "/_cluster/settings", settings)
	return err
}
