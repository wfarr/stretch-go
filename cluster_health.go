package stretch

type ClusterHealth struct {
	Status              string `json:"status,omitempty"`
	ClusterName         string `json:"cluster_name,omitempty"`
	TimedOut            bool   `json:"timed_out,omitempty"`
	NumberOfNodes       int    `json:"number_of_nodes,omitempty"`
	NumberOfDataNodes   int    `json:"number_of_data_nodes,omitempty"`
	ActivePrimaryShards int    `json:"active_primary_shards,omitempty"`
	ActiveShards        int    `json:"active_shards,omitempty"`
	RelocatingShards    int    `json:"relocating_shards,omitempty"`
	InitializingShards  int    `json:"initializing_shards,omitempty"`
	UnassignedShards    int    `json:"unassigned_shards,omitempty"`
	// Map of "index-name-as-string": IndexHealth
	Indices map[string]*IndexHealth `json:"indices"`
}

type IndexHealth struct {
	Status              string `json:"status"`
	NumberOfShards      int    `json:"number_of_shards"`
	NumberOfReplicas    int    `json:"number_of_replicas"`
	ActivePrimaryShards int    `json:"active_primary_shards"`
	ActiveShards        int    `json:"active_shards"`
	RelocatingShards    int    `json:"relocating_shards"`
	InitializingShards  int    `json:"initializing_shards"`
	UnassignedShards    int    `json:"unassigned_shards"`
}

func (c *Cluster) GetHealth(indices ...string) (ClusterHealth, error) {
	var data ClusterHealth
	err := c.Client.Get(&data, "/_cluster/health?level=indices")

	return data, err
}
