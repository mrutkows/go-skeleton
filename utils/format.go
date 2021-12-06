package utils

import (
	"encoding/json"
	"fmt"

	"github.com/hokaccha/go-prettyjson"
)

func PrettyJSON(rawData interface{}) (string, error) {
	formatter := prettyjson.NewFormatter()
	bytes, err := formatter.Marshal(rawData)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func FormatStructAsJsonString(a interface{}) string {
	out, err := json.Marshal(a)
	if err == nil {
		return string(out)
	}
	return ""
}

func ConvertMapToJSONString(name string, mapIn interface{}) string {
	strMapOut, _ := json.MarshalIndent(mapIn, "", "  ")
	return fmt.Sprintf("%s: %s", name, string(strMapOut))
}
