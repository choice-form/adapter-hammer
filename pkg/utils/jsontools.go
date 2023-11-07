package utils

import "encoding/json"

func Struct2Map(input any, output *map[string]any) {
	b, _ := json.Marshal(input)
	json.Unmarshal(b, output)
}

func Map2Struct(input map[string]any, output any) {
	b, _ := json.Marshal(&input)
	json.Unmarshal(b, output)
}
