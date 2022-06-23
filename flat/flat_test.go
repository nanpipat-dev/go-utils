package flat

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	tests := []struct {
		given   string
		options *Options
		want    map[string]interface{}
	}{
		// test with different primitives
		// String: 'world',
		// Number: 1234.99,
		// Boolean: true,
		// null: null,
		{
			`{"hello": "world"}`,
			nil,
			map[string]interface{}{"hello": "world"},
		},
		{
			`{"hello": 1234.99}`,
			nil,
			map[string]interface{}{"hello": 1234.99},
		},
		{
			`{"hello": true}`,
			nil,
			map[string]interface{}{"hello": true},
		},
		{
			`{"hello": null}`,
			nil,
			map[string]interface{}{"hello": nil},
		},
		// nested once
		{
			`{"hello":{}}`,
			nil,
			map[string]interface{}{"hello": map[string]interface{}{}},
		},
		{
			`{"hello":{"world":"good morning"}}`,
			nil,
			map[string]interface{}{"hello.world": "good morning"},
		},
		{
			`{"hello":{"world":1234.99}}`,
			nil,
			map[string]interface{}{"hello.world": 1234.99},
		},
		{
			`{"hello":{"world":true}}`,
			nil,
			map[string]interface{}{"hello.world": true},
		},
		{
			`{"hello":{"world":null}}`,
			nil,
			map[string]interface{}{"hello.world": nil},
		},
		// empty slice
		{
			`{"hello":{"world":[]}}`,
			nil,
			map[string]interface{}{"hello.world": []interface{}{}},
		},
		// slice
		{
			`{"hello":{"world":["one","two"]}}`,
			nil,
			map[string]interface{}{
				"hello.world.0": "one",
				"hello.world.1": "two",
			},
		},
		// nested twice
		{
			`{"hello":{"world":{"again":"good morning"}}}`,
			nil,
			map[string]interface{}{"hello.world.again": "good morning"},
		},
		// multiple keys
		{
			`{
				"hello": {
					"lorem": {
						"ipsum":"again",
						"dolor":"sit"
					}
				},
				"world": {
					"lorem": {
						"ipsum":"again",
						"dolor":"sit"
					}
				}
			}`,
			nil,
			map[string]interface{}{
				"hello.lorem.ipsum": "again",
				"hello.lorem.dolor": "sit",
				"world.lorem.ipsum": "again",
				"world.lorem.dolor": "sit"},
		},
		// empty object
		{
			`{"hello":{"empty":{"nested":{}}}}`,
			nil,
			map[string]interface{}{"hello.empty.nested": map[string]interface{}{}},
		},
		// custom delimiter
		{
			`{"hello":{"world":{"again":"good morning"}}}`,
			&Options{
				Delimiter: ":",
				MaxDeep:   20,
			},
			map[string]interface{}{"hello:world:again": "good morning"},
		},
		// custom depth
		{
			`{
				"hello": {
					"world": {
						"again": "good morning"
					}
				},
				"lorem": {
					"ipsum": {
						"dolor": "good evening"
					}
				}
			}
			`,
			&Options{
				MaxDeep:   2,
				Delimiter: ".",
			},
			map[string]interface{}{
				"hello.world": map[string]interface{}{"again": "good morning"},
				"lorem.ipsum": map[string]interface{}{"dolor": "good evening"},
			},
		},
		// custom safe = true
		{
			`{"hello":{"world":["one","two"]}}`,
			&Options{
				Safe:      true,
				Delimiter: ".",
			},
			map[string]interface{}{
				"hello.world": []interface{}{"one", "two"},
			},
		},
	}
	for i, test := range tests {
		var given interface{}
		err := json.Unmarshal([]byte(test.given), &given)
		if err != nil {
			t.Errorf("%d: failed to unmarshal test: %v", i+1, err)
		}
		got, err := Flatten(given.(map[string]interface{}), test.options)
		if err != nil {
			t.Errorf("%d: failed to flatten: %v", i+1, err)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: mismatch, got: %v want: %v", i+1, got, test.want)
		}
	}
}
