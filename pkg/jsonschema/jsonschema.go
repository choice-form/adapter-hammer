package jsonschema

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/xeipuuv/gojsonschema"
)

type ValidError struct {
	Pass bool     `json:"pass,omitempty"`
	Errs []string `json:"errs,omitempty"`
	Err  error    `json:"err,omitempty"`
}

func (ve *ValidError) Error() string {
	if ve.Err != nil {
		return ve.Err.Error()
	}
	return strings.Join(ve.Errs, ",")
}

type InterfaceCellProperty string

var (
	INPUT          InterfaceCellProperty = "input"
	OUTPUT         InterfaceCellProperty = "output"
	jsonSchemaFile                       = "jsonschema.json"
	callPath                             = "properties.InterfaceCall.properties.%s.properties.input"
	jsonSchemaMap                        = make(map[string]string) // 缓存jsonSchema
)

type ValidInput struct {
	jsonSchema *string
	schemaPath string
	method     string
	input      map[string]any
}

func NewValidInput(method string) *ValidInput {
	return &ValidInput{
		method: method,
	}
}

func (v *ValidInput) SetJsonSchema(jsonSchemaPath string) error {
	if val, ok := jsonSchemaMap[jsonSchemaPath]; ok {
		v.jsonSchema = &val
		return nil
	}
	str, err := ReadJsonSchema(jsonSchemaPath)
	if err != nil {
		return err
	}
	v.jsonSchema = &str
	jsonSchemaMap[jsonSchemaPath] = str
	return nil
}

// 验证 input 是否符合 jsonschema
func (v *ValidInput) Valid(input map[string]any) (pass bool, err error) {
	version := SchemaVersion(*v.jsonSchema)

	propertyValue := gjson.Get(*v.jsonSchema, fmt.Sprintf(callPath, handleMethod(v.method)))
	fmt.Printf("method: %s; value: %s\n", v.method, propertyValue)

	var jsonMap map[string]any
	json.Unmarshal([]byte(propertyValue.String()), &jsonMap)
	jsonMap["$schema"] = version

	loader := gojsonschema.NewGoLoader(&jsonMap)
	inputLoader := gojsonschema.NewGoLoader(&input)
	result, err := gojsonschema.Validate(loader, inputLoader)
	if err != nil {
		return false, &ValidError{
			Pass: false,
			Err:  err,
		}
	}

	if result.Valid() {
		return true, nil
	}

	ve := &ValidError{
		Pass: false,
		Errs: []string{},
	}

	if len(result.Errors()) > 0 {
		for _, e := range result.Errors() {
			ve.Errs = append(ve.Errs, e.String())
		}
	}

	return false, ve
}

func handleMethod(method string) string {
	method = strings.ReplaceAll(method, ".", "\\.")
	method = strings.ReplaceAll(method, "*", "\\*")
	return strings.ReplaceAll(method, "?", "\\?")
}

func SchemaVersion(schema string) string {
	_schemaVersion := gjson.Get(schema, "$schema")
	return _schemaVersion.String()
}

func ReadJsonSchema(jsonSchemaPath string) (string, error) {
	// wd, _ := os.Getwd()
	// fmt.Printf("wd: %v\n", wd)
	// content, err := os.ReadFile(filepath.Join(wd, jsonSchemaFile))
	content, err := os.ReadFile(jsonSchemaPath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
