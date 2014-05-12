package stretch

import (
	"net/http"

	"net/http/httptest"

	"testing"
)

func testServer(resp string) (ts *httptest.Server) {
	ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}))
	return
}

func TestClusterInfo(t *testing.T) {
	ts := testServer(`{
		"ok" : true,
		"status" : 200,
		"name" : "boxen",
		"version" : {
			"number" : "0.90.5",
			"build_hash" : "c8714e8e0620b62638f660f6144831792b9dedee",
			"build_timestamp" : "2013-09-17T12:50:20Z",
			"build_snapshot" : false,
			"lucene_version" : "4.4"
		},
		"tagline" : "You Know, for Search"
	}`)

	defer ts.Close()

	cluster := &Cluster{&Client{URL: ts.URL}}
	ci := cluster.GetInfo()

	if !ci.Ok {
		t.Error("ci.Ok was not true")
	}

	if ci.Status != 200 {
		t.Error("ci.Status was not 200")
	}

	if ci.Name != "boxen" {
		t.Error("ci.Name was not boxen")
	}

	if ci.Tagline != "You Know, for Search" {
		t.Error("ci.Tagline was not You Know, for Search")
	}

	if ci.Version.Number != "0.90.5" {
		t.Error("ci.Version.Number was not 0.90.5")
	}

	if ci.Version.BuildHash != "c8714e8e0620b62638f660f6144831792b9dedee" {
		t.Error("ci.Version.BuildHash was not c8714e8e0620b62638f660f6144831792b9dedee")
	}

	if ci.Version.BuildTimestamp != "2013-09-17T12:50:20Z" {
		t.Error("ci.Version.BuildTimestamp was not 2013-09-17T12:50:20Z")
	}

	if ci.Version.BuildSnapshot != false {
		t.Error("ci.Version.BuildSnapshot was not false")
	}

	if ci.Version.LuceneVersion != "4.4" {
		t.Error("ci.Version.LuceneVersion was not 4.4")
	}
}
