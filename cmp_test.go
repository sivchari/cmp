package cmp_test

import (
	"embed"
	"testing"

	"github.com/sivchari/cmp"
)

//go:embed testdata/*
var testdata embed.FS

type Want struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Items   []Item `json:"items"`
}

type Item struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func TestDiff(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		fileName string
		want     Want
	}{
		"json.json": {
			fileName: "testdata/json.json",
			want: Want{
				Name:    "example",
				Version: "1.0.0",
				Items: []Item{
					{ID: 1, Title: "Item One", Description: "This is the first item."},
					{ID: 2, Title: "Item Two", Description: "This is the second item."},
				},
			},
		},
		"yaml.yaml": {
			fileName: "testdata/yaml.yaml",
			want: Want{
				Name:    "example",
				Version: "1.0.0",
				Items: []Item{
					{ID: 1, Title: "Item One", Description: "This is the first item."},
					{ID: 2, Title: "Item Two", Description: "This is the second item."},
				},
			},
		},
		"nested.test.json": {
			fileName: "testdata/nested.test.json",
			want: Want{
				Name:    "example",
				Version: "1.0.0",
				Items: []Item{
					{ID: 1, Title: "Item One", Description: "This is the first item."},
					{ID: 2, Title: "Item Two", Description: "This is the second item."},
				},
			},
		},
		"nested.test.yaml": {
			fileName: "testdata/nested.test.yaml",
			want: Want{
				Name:    "example",
				Version: "1.0.0",
				Items: []Item{
					{ID: 1, Title: "Item One", Description: "This is the first item."},
					{ID: 2, Title: "Item Two", Description: "This is the second item."},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if diff := cmp.Diff(t, testdata, tc.fileName, tc.want); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
