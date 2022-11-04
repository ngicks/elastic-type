package test_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"testing"
)

var ELASTICSEARCH_URL *url.URL

func init() {
	ELASTICSEARCH_URL, _ = url.Parse(os.Getenv("ELASTICSEARCH_URL"))
}

func skipIfEsNotReachable(t *testing.T, esURL url.URL, preferFail bool) {
	// We need to send a fetch request to some of Elasticsearch specific pathes
	// to ensure that there is a reachable instance.
	// "_cluster/health" is just one of those pathes. It could be replaced with one of any else es specific pathes.
	esURL.Path = ""
	esURL = *esURL.JoinPath("_cluster", "health")

	skipOrFail := func(format string, args ...interface{}) {
		if preferFail {
			t.Fatalf(format, args...)
		} else {
			t.Skipf(format, args...)
		}
	}

	res, err := http.Get(esURL.String())
	if err != nil {
		skipOrFail("request to Elasticsearch failed: %v", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		skipOrFail("request to Elasticsearch failed: %v", err)
	}

	// body must be like.
	// {
	//		"cluster_name":"docker-cluster",
	//		"status":"green",
	//		"timed_out":false,
	//		"number_of_nodes":1,
	//		"number_of_data_nodes":1,
	//		"active_primary_shards":1,
	//		"active_shards":1,
	//		"relocating_shards":0,
	//		"initializing_shards":0,
	//		"unassigned_shards":0,
	//		"delayed_unassigned_shards":0,
	//		"number_of_pending_tasks":0,
	//		"number_of_in_flight_fetch":0,
	//		"task_max_waiting_in_queue_millis":0,
	//		"active_shards_percent_as_number":100.0
	// }

	var bodyUnmarshalled map[string]interface{}

	err = json.Unmarshal(body, &bodyUnmarshalled)
	if err != nil {
		skipOrFail("request to Elasticsearch failed: %v", err)
	}

	status, ok := bodyUnmarshalled["status"]
	if !ok {
		skipOrFail("Returned body is unknown structure: expected to have string `status` field but actually is\n%+v", bodyUnmarshalled)
	}

	str, ok := status.(string)
	if !ok {
		skipOrFail("Returned body is unknown structure: expected to have string `status` field but actually is\n%+v", bodyUnmarshalled)
	}

	if str != "green" {
		skipOrFail("Status is not green: you must wait until Elasticsearch to be ready.")
	}
}
