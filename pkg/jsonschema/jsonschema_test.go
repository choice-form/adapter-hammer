package jsonschema

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/tidwall/gjson"
	"github.com/xeipuuv/gojsonschema"
)

func TestValidJsonSchema1(t *testing.T) {
	t.Run("valid jsonSchema.json", func(t *testing.T) {
		var method = "POST/open-apis/im/v1/messages"
		var input = map[string]any{
			"id":          1111,
			"name":        "true",
			"description": "desc",
			"type":        "number",
		}

		vi := NewValidInput(method)
		if err := vi.SetJsonSchema("./test.json"); err != nil {
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

func TestValidJsonSchema(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		var input = map[string]any{
			"id":          "11",
			"name":        "true",
			"description": "desc",
			"type":        "number",
		}

		b, err := os.ReadFile("test.json")
		if err != nil {
			t.Error(err)
		}

		def := gjson.Get(string(b), "definitions").Value()
		inp := gjson.Get(string(b), "properties.abcdefg.properties.input").Value()

		if v, ok := inp.(map[string]any); ok {
			fmt.Printf("v: %v\n", v)
			v["definitions"] = def
		}

		rootLoader := gojsonschema.NewGoLoader(inp)
		printJsonLoader("rootLoader", rootLoader)

		mainSchema := gojsonschema.NewSchemaLoader()
		mainSchema.Draft = gojsonschema.Draft7
		mainSchema.AutoDetect = true
		mainLoader, err := mainSchema.Compile(rootLoader)
		if err != nil {
			t.Error(err)
		}

		res, err := mainLoader.Validate(gojsonschema.NewGoLoader(input))
		if err != nil {
			t.Error(err)
		}
		if e := res.Errors(); len(e) > 0 {
			fmt.Printf("errs: %v\n", e)
		}

		fmt.Printf("res===>: %v\n", res.Valid())
	})
}

func generateJson(src []byte) *bytes.Buffer {
	var b bytes.Buffer
	json.Indent(&b, src, "", "\t")
	return &b
}

func printJsonLoader(name string, src gojsonschema.JSONLoader) {
	value, err := src.LoadJSON()
	if err != nil {
		fmt.Printf("err1: %v\n", err)
	}
	b, err := json.Marshal(value)
	if err != nil {
		fmt.Printf("err2: %v\n", err)
	}
	fmt.Printf("%v: %v\n", name, generateJson(b).String())
}
