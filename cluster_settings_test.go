package stretch

import (
	"net/http"
	"testing"
)

func TestClusterSettingsOhNinety(t *testing.T) {
	ts := testServer(http.StatusOK, `{
		"persistent": {
			"cluster.routing.allocation.disable_allocation": "true"
		},
		"transient": {
			"cluster.routing.allocation.disable_replica_allocation": "true",
			"indices.recovery.max_bytes_per_sec" : "2gb",
			"indices.recovery.concurrent_streams" : "24",
			"cluster.routing.allocation.node_concurrent_recoveries" : "6"
		}
	}`)

	defer ts.Close()

	cluster := &Cluster{&Client{URL: ts.URL}}
	settings := cluster.GetSettings()

	if settings.Persistent["cluster.routing.allocation.disable_allocation"] != "true" {
		t.Fail()
	}

	if settings.Transient["cluster.routing.allocation.disable_replica_allocation"] != "true" {
		t.Fail()
	}

	if settings.Transient["indices.recovery.max_bytes_per_sec"] != "2gb" {
		t.Fail()
	}
}

func TestClusterSettingsOneOh(t *testing.T) {
	ts := testServer(http.StatusOK, `{
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

	if settings.Persistent["cluster.routing.allocation.enable"] != "all" {
		t.Fail()
	}

	if settings.Transient["cluster.routing.allocation.enable"] != "new_primaries" {
		t.Fail()
	}
}
