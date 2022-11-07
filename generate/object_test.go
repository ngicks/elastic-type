package generate_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/ngicks/elastic-type/generate"
	"github.com/ngicks/elastic-type/mapping"
	"github.com/stretchr/testify/require"
)

var nestedObjectMapping = []byte(`{
  "mappings": {
    "properties": { 
      "manager": { 
        "properties": {
          "age":  { "type": "integer" },
          "name": { 
            "properties": {
              "first": { "type": "text" },
              "last":  { "type": "text" }
            }
          }
        }
      }
    }
  }
}`)

var TEST_GENERATE_TARGET string

func init() {
	TEST_GENERATE_TARGET = os.Getenv("TEST_GENERATE_TARGET")
}

func applyGoimports(srcPath string) (stdout []byte, err error) {
	return exec.Command("goimports", "-w", srcPath).Output()
}

func TestObjectGenerate(t *testing.T) {
	// if TEST_GENERATE_TARGET == "" {
	// 	t.Skip("TEST_GENERATE_TARGET is not set")
	// }

	require := require.New(t)

	settings := mapping.IndexSettings{}

	err := json.Unmarshal(nestedObjectMapping, &settings)
	require.NoError(err)

	highLevelTypes, rawTypes, err := generate.Generate(
		settings.Mappings.Properties,
		"sample",
		generate.Options{},
	)

	t.Logf("high level types: %s", highLevelTypes)
	t.Logf("raw types: %s", rawTypes)
	t.Logf("err: %v", err)
}
