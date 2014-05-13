package stretch

import (
	"testing"
)

func TestClusterSettingsOhNinety(t *testing.T) {
	ts := testServer(`{
		"persistent": {
			"cluster.routing.allocation.disable_allocation": true
		},
		"transient": {
			"cluster.routing.allocation.disable_replica_allocation": true,
			"indices.recovery.max_bytes_per_sec" : "2gb",
			"indices.recovery.concurrent_streams" : "24",
			"cluster.routing.allocation.node_concurrent_recoveries" : "6"
		}
	}`)

	defer ts.Close()

	cluster := &Cluster{&Client{URL: ts.URL}}
	settings := cluster.GetSettings()

	if settings.Persistent.ClusterRoutingAllocationDisableAllocation != true {
		t.Fail()
	}

	if settings.Persistent.ClusterRoutingAllocationDisableReplicaAllocation != false {
		t.Fail()
	}

	if settings.Transient.ClusterRoutingAllocationDisableAllocation != false {
		t.Fail()
	}

	if settings.Transient.ClusterRoutingAllocationDisableReplicaAllocation != true {
		t.Fail()
	}
}

func TestClusterSettingsOneOh(t *testing.T) {
	ts := testServer(`{
		"persistent": {
			"cluster.routing.allocation.enable": "all"
		},
		"transient": {
			"cluster.routing.allocation.enable": "new_primaries"
		}
	}`)

	defer ts.Close()

	cluster := &Cluster{&Client{URL: ts.URL}}
	settings := cluster.GetSettings()

	if settings.Persistent.ClusterRoutingAllocationEnable != "all" {
		t.Fail()
	}

	if settings.Transient.ClusterRoutingAllocationEnable != "new_primaries" {
		t.Fail()
	}
}
