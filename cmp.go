// Package cmp provides a utility function to compare a given value against
package cmp

import (
	"embed"
	"encoding/json"
	"strings"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/google/go-cmp/cmp"
)

// Diff compares the provided value of type T against the contents of a file
// in the embedded filesystem fs. The file is expected to be in JSON or YAML
// format. If the file cannot be read or parsed, an error is logged to the
// testing.TB interface. The function uses the google/go-cmp package to
// perform the comparison and reports any differences.
func Diff[T any](t testing.TB, fs embed.FS, fileName string, typ T, opts ...cmp.Option) string {
	t.Helper()

	file, err := fs.ReadFile(fileName)
	if err != nil {
		t.Errorf("failed to read file %s: %v", fileName, err)
	}

	var want T

	splits := strings.Split(fileName, ".")
	ext := splits[len(splits)-1]

	switch ext {
	case "json":
		if err := json.Unmarshal(file, &want); err != nil {
			t.Errorf("failed to unmarshal JSON from file %s: %v", fileName, err)
		}
	case "yaml", "yml":
		if err := yaml.Unmarshal(file, &want); err != nil {
			t.Errorf("failed to unmarshal YAML from file %s: %v", fileName, err)
		}
	}

	return cmp.Diff(typ, want, opts...)
}
