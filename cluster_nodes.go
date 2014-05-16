package stretch

type ClusterNodes struct {
	OK          bool                   `json:"ok"`
	ClusterName string                 `json:"cluster_name"`
	Nodes       map[string]ClusterNode `json:"nodes"`
}

type ClusterNode struct {
	Name             string            `json:"name"`
	TransportAddress string            `json:"transport_address"`
	Hostname         string            `json:"hostname"`
	Version          string            `json:"version"`
	HTTPAddress      string            `json:"http_address"`
	Attributes       map[string]string `json:"attributes"`
}

func (c *Cluster) GetNodes() (ClusterNodes, error) {
	var data ClusterNodes
	err := c.Client.Get(&data, "/_cluster/nodes")
	return data, err
}
