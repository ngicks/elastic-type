package test_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/ngicks/gommon/pkg/randstr"
	typeparamcommon "github.com/ngicks/type-param-common"
)

var (
	ELASTICSEARCH_URL *url.URL
	client            *elasticsearch.Client
)

func init() {
	ELASTICSEARCH_URL, _ = url.Parse(os.Getenv("ELASTICSEARCH_URL"))

	var err error
	// doc says NewDefaultClient reads ELASTICSEARCH_URL env var.
	client, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{ELASTICSEARCH_URL.String()},
	})
	if err != nil {
		panic(err)
	}
}

func must[T any](v T, err error) T {
	return typeparamcommon.Must(v, err)
}

// toAnyMap marshals v into bin and then unmarshals to map[string]any.
func toAnyMap(v any) map[string]any {
	bin := must(json.Marshal(v))

	var anyMap map[string]any
	json.Unmarshal(bin, &anyMap)

	return anyMap
}

func skipIfEsNotReachable(t *testing.T, esURL url.URL, preferFail bool) {
	// We need to send a fetch request to some of Elasticsearch specific paths
	// to ensure that there is a reachable instance.
	// "_cluster/health" is just one of those paths. It could be replaced with one of any else es specific paths.

	skipOrFail := func(format string, args ...interface{}) {
		if preferFail {
			t.Fatalf(format, args...)
		} else {
			t.Skipf(format, args...)
		}
	}

	res, err := client.Cluster.Health()
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

type EsTestHelper[T any] struct {
	Client    *elasticsearch.Client
	IndexName string
}

func createRandomIndex[T any](client *elasticsearch.Client, settings []byte) (*EsTestHelper[T], error) {
	randomIndexName, err := randstr.New(randstr.Hex()).String()
	if err != nil {
		return nil, err
	}

	opt := client.Indices.Create.WithBody(bytes.NewReader(settings))
	res, err := client.Indices.Create(randomIndexName, opt)
	if err != nil {
		return nil, err
	} else if res.IsError() {
		return nil, fmt.Errorf("%s", res.String())
	}

	return &EsTestHelper[T]{
		Client:    client,
		IndexName: randomIndexName,
	}, nil
}

func (h *EsTestHelper[T]) Delete() error {
	res, err := client.Indices.Delete([]string{h.IndexName})
	if err != nil {
		return err
	} else if res.IsError() {
		return fmt.Errorf("%s", res.String())
	}
	return nil
}

func (h *EsTestHelper[T]) GetMapping() ([]byte, error) {
	res, err := h.Client.Indices.GetMapping()
	if err != nil {
		return nil, err
	} else if res.IsError() {
		return nil, fmt.Errorf("%s", res.String())
	}
	bin, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return bin, nil
}

type CommonResponse struct {
	Index_       string `json:"_index"`
	Id_          string `json:"_id"`
	Version_     int    `json:"_version"`
	SeqNo_       int    `json:"_seq_no"`
	PrimaryTerm_ int    `json:"_primary_term"`
}

type IndexResult struct {
	CommonResponse
	Result  string `json:"result"`
	Shards_ Shards `json:"_shards"`
}

type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Failed     int `json:"failed"`
}

type FetchDocResult[T any] struct {
	CommonResponse
	Found   bool `json:"found"`
	Source_ T    `json:"_source"`
}

func (h *EsTestHelper[T]) PostDoc(doc T) (docId string, err error) {
	bin, err := json.Marshal(doc)
	if err != nil {
		return "", err
	}

	res, err := h.Client.Index(h.IndexName, bytes.NewReader(bin))
	if err != nil {
		return "", err
	} else if res.IsError() {
		return "", fmt.Errorf("%s", res.String())
	}

	var result IndexResult
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.Id_, nil
}

func (h *EsTestHelper[T]) GetDoc(id string) (doc FetchDocResult[T], err error) {
	res, err := h.Client.Get(h.IndexName, id)
	if err != nil {
		return FetchDocResult[T]{}, err
	} else if res.IsError() {
		return FetchDocResult[T]{}, fmt.Errorf("%s", res.String())
	}

	err = json.NewDecoder(res.Body).Decode(&doc)
	if err != nil {
		return FetchDocResult[T]{}, err
	}
	return doc, nil
}
