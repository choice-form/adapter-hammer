package jsonschema

import (
	"fmt"
	"testing"
)

func TestValidJsonSchema(t *testing.T) {
	t.Run("valid json schema", func(t *testing.T) {
		var method = "GET /v1.0/contact/empLeaveRecords"
		var input = map[string]any{
			"startTime":  "2023-01-01T00:00:00Z",
			"endTime":    "",
			"nextToken":  "",
			"maxResults": 10,
		}
		// pass, errs, err := ValidJsonSchema(method, input)

		vi := NewValidInput(method)
		if err := vi.SetJsonSchema("../../jsonschema.json"); err != nil {
			t.Error(err)
			return
		}

		pass, err := vi.Valid(input)
		if pass {
			t.Log("pass")
		}
		if err != nil {
			if e, ok := err.(*ValidError); ok {
				t.Errorf("errs: %s", e)
			}
		}

	})
}

func TestSchemaVersion(t *testing.T) {
	t.Run("get jsonschema version", func(t *testing.T) {
		var schema = `
			{
				"$schema": "http://json-schema.org/draft-07/schema#",
				"type": "object",
			}
		`
		version := SchemaVersion(schema)
		fmt.Printf("version: %v\n", version)
	})
}
