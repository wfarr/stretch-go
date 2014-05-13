package stretch

type FullClusterSettings struct {
	Persistent map[string]string `json:"persistent,omitempty"`
	Transient  map[string]string `json:"transient,omitempty"`
}

func (c *Cluster) GetSettings() (data FullClusterSettings) {
	c.Client.Get(&data, "/_cluster/settings")
	return
}

func (c *Cluster) SetSettings(settings interface{}) (err error) {
	err = c.Client.Put(nil, "/_cluster/settings", settings)
	return
}
