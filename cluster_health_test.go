package stretch

import (
	"testing"
)

func TestClusterHealth(t *testing.T) {
	ts := testServer(`{
		"status": "tangerine",
		"cluster_name": "foobar",
		"timed_out" : false,
		"number_of_nodes" : 1,
		"number_of_data_nodes" : 1,
		"active_primary_shards" : 10,
		"active_shards" : 20,
		"relocating_shards" : 2,
		"initializing_shards" : 0,
		"unassigned_shards" : 0
		}`)

	defer ts.Close()

	cluster := &Cluster{&Client{URL: ts.URL}}
	health := cluster.GetHealth()

	if health.Status != "tangerine" {
		t.Fail()
	}
	if health.ClusterName != "foobar" {
		t.Fail()
	}
	if health.TimedOut != false {
		t.Fail()
	}
	if health.NumberOfNodes != 1 {
		t.Fail()
	}
	if health.NumberOfDataNodes != 1 {
		t.Fail()
	}
	if health.ActivePrimaryShards != 10 {
		t.Fail()
	}
	if health.ActiveShards != 20 {
		t.Fail()
	}
	if health.RelocatingShards != 2 {
		t.Fail()
	}
	if health.InitializingShards != 0 {
		t.Fail()
	}
	if health.UnassignedShards != 0 {
		t.Fail()
	}
}

func TestClusterHealthWithIndices(t *testing.T) {
	ts := testServer(`{
		"status": "tangerine",
		"cluster_name": "foobar",
		"timed_out" : false,
		"number_of_nodes" : 1,
		"number_of_data_nodes" : 1,
		"active_primary_shards" : 10,
		"active_shards" : 20,
		"relocating_shards" : 2,
		"initializing_shards" : 0,
		"unassigned_shards" : 0,
		"indices" : {
			"test1": {
				"status" : "green",
				"number_of_shards" : 1,
				"number_of_replicas" : 1,
				"active_primary_shards" : 1,
				"active_shards" : 2,
				"relocating_shards" : 0,
				"initializing_shards" : 0,
				"unassigned_shards" : 0
			},
			"test2": {
				"status" : "green",
				"number_of_shards" : 1,
				"number_of_replicas" : 1,
				"active_primary_shards" : 1,
				"active_shards" : 2,
				"relocating_shards" : 0,
				"initializing_shards" : 0,
				"unassigned_shards" : 0
			}
		}`)

	defer ts.Close()

	cluster := &Cluster{&Client{URL: ts.URL}}
	health := cluster.GetHealth()

	for indexName, indexHealth := range health.Indices {
		if indexName == "" {
			t.Fail()
		}

		if indexHealth.Status != "green" {
			t.Fail()
		}

		if indexHealth.NumberOfShards != 1 {
			t.Fail()
		}

		if indexHealth.NumberOfReplicas != 1 {
			t.Fail()
		}

		if indexHealth.ActivePrimaryShards != 1 {
			t.Fail()
		}

		if indexHealth.ActiveShards != 2 {
			t.Fail()
		}

		if indexHealth.RelocatingShards != 0 {
			t.Fail()
		}

		if indexHealth.InitializingShards != 0 {
			t.Fail()
		}

		if indexHealth.UnassignedShards != 0 {
			t.Fail()
		}
	}
}
