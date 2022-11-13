package test_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/ngicks/gommon/pkg/randstr"
	typeparamcommon "github.com/ngicks/type-param-common"
	"github.com/ngicks/type-param-common/slice"
)

var ELASTICSEARCH_URL *url.URL

func init() {
	ELASTICSEARCH_URL, _ = url.Parse(os.Getenv("ELASTICSEARCH_URL"))
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

func getOne(v map[string]map[string]any) map[string]any {
	for _, v := range v {
		return v
	}
	return nil
}

func skipIfEsNotReachable(t *testing.T, esURL url.URL, preferFail bool) {
	// We need to send a fetch request to some of Elasticsearch specific paths
	// to ensure that there is a reachable instance.
	// "_cluster/health" is just one of those paths. It could be replaced with one of any else es specific paths.
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

func jsonRequest(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if slice.Has([]string{http.MethodPatch, http.MethodPost, http.MethodPut}, method) {
		req.Header.Set("Content-Type", "application/json")
	}
	return http.DefaultClient.Do(req)
}

func createRandomIndex(esURL url.URL, settings []byte) (indexName string, err error) {
	randomIndexName, err := randstr.New(randstr.Hex()).String()
	if err != nil {
		panic(err)
	}

	esURL.Path = ""
	esURL = *esURL.JoinPath(randomIndexName)

	res, err := jsonRequest(http.MethodPut, esURL.String(), bytes.NewReader(settings))
	if err != nil {
		return "", err
	}

	bodyBin, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}

	if res.StatusCode > 300 {
		return "", fmt.Errorf(
			"error response: status = %d, body = %s",
			res.StatusCode,
			string(bodyBin),
		)
	}

	body := map[string]any{}
	err = json.Unmarshal(bodyBin, &body)
	if err != nil {
		return "", err
	}

	ack, ok := body["acknowledged"]
	if !ok {
		return "", fmt.Errorf(
			"error response: status = %d, body = %s",
			res.StatusCode,
			string(bodyBin),
		)
	}
	ackBool, ok := ack.(bool)
	if !ok || !ackBool {
		return "", fmt.Errorf(
			"error response: status = %d, body = %s",
			res.StatusCode,
			string(bodyBin),
		)
	}

	return randomIndexName, nil
}

func deleteIndex(esURL url.URL, indexName string) error {
	esURL.Path = ""
	esURL = *esURL.JoinPath(indexName)
	_, err := jsonRequest(http.MethodDelete, esURL.String(), nil)
	return err
}

func getMapping(esURL url.URL, indexName string) ([]byte, error) {
	esURL.Path = ""
	mappingsURL := *esURL.JoinPath(indexName, "_mappings")

	res, err := jsonRequest(http.MethodGet, mappingsURL.String(), nil)
	if err != nil {
		return nil, err
	}

	bodyBin, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return nil, err
	}
	return bodyBin, nil
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

type FetchDocResult struct {
	CommonResponse
	Found   bool `json:"found"`
	Source_ any  `json:"_source"`
}

func postDoc(esURL url.URL, indexName string, doc any) (IndexResult, error) {
	esURL.Path = ""
	docURL := *esURL.JoinPath(indexName, "_doc")

	bin, err := json.Marshal(doc)
	if err != nil {
		return IndexResult{}, err
	}
	res, err := jsonRequest(http.MethodPost, docURL.String(), bytes.NewReader(bin))
	if err != nil {
		return IndexResult{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return IndexResult{}, err
	}
	defer res.Body.Close()

	var result IndexResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return IndexResult{}, err
	}

	if result.Result != "created" {
		return IndexResult{}, fmt.Errorf("%s", string(body))
	}

	return result, nil
}

func getDoc(esURL url.URL, indexName, docId string) (doc FetchDocResult, err error) {
	esURL.Path = ""
	docURL := *esURL.JoinPath(indexName, "_doc", docId)

	res, err := jsonRequest(http.MethodGet, docURL.String(), nil)
	if err != nil {
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	res.Body.Close()

	var result FetchDocResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return
	}
	return result, nil
}
