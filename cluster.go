package stretch

type Cluster struct {
	Client *Client
}

type ClusterVersion struct {
	Number         string `json:"number"`
	BuildHash      string `json:"build_hash"`
	BuildTimestamp string `json:"build_timestamp"`
	BuildSnapshot  bool   `json:"build_snapshot"`
	LuceneVersion  string `json:"lucene_version"`
}

type ClusterInfo struct {
	Ok      bool           `json:"ok"`
	Status  int            `json:"status"`
	Name    string         `json:"name"`
	Version ClusterVersion `json:"version"`
	Tagline string         `json:"tagline"`
}

type ClusterState struct {
	ClusterName string `json:"cluster_name"`
	MasterNode  string `json:"master_node"`
}

func (c *Cluster) GetInfo() (ClusterInfo, error) {
	var data ClusterInfo
	err := c.Client.Get(&data, "/")
	return data, err
}

func (c *Cluster) GetState() (ClusterState, error) {
	var data ClusterState
	err := c.Client.Get(&data, "/_cluster/state")
	return data, err
}
