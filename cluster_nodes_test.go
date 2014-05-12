package stretch

import (
	"testing"
)

func TestClusterNodes(t *testing.T) {
	ts := testServer(`{
		"ok": true,
		"cluster_name": "foobar",
		"nodes": {
			"9W396AtLRBiGSCQLlVqOgA": {
				"name": "host1",
				"transport_address": "inet[/127.0.0.1:9300]",
				"hostname": "test.local",
				"version": "0.90.10",
				"http_address": "inet[/127.0.0.1:9200]",
				"attributes": {
					"data": "true",
					"master": "false"
				}
			}
		}
	}`)

	defer ts.Close()

	cluster := &Cluster{&Client{URL: ts.URL}}
	nodes := cluster.GetNodes()

	if nodes.OK != true {
		t.Fail()
	}

	if nodes.ClusterName != "foobar" {
		t.Fail()
	}

	for _, node := range nodes.Nodes {
		if node.Name != "host1" {
			t.Fail()
		}

		if node.TransportAddress != "inet[/127.0.0.1:9300]" {
			t.Fail()
		}

		if node.Hostname != "test.local" {
			t.Fail()
		}

		if node.Version != "0.90.10" {
			t.Fail()
		}

		if node.HTTPAddress != "inet[/127.0.0.1:9200]" {
			t.Fail()
		}

		if node.Attributes["data"] != "true" {
			t.Fail()
		}

		if node.Attributes["master"] != "false" {
			t.Fail()
		}
	}
}