package log

import (
	"fmt"
	"reflect"
	"strings"
)

func FormatStruct(structName string, field interface{}) (string, error) {

	if reflect.ValueOf(field).Kind() != reflect.Struct {
		return "", fmt.Errorf("invalid `Struct`; actual Type: (%v)", reflect.TypeOf(field))
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("  %s (%s) = {\n", structName, reflect.TypeOf(field)))

	flagNames := reflect.TypeOf(field)
	numNames := flagNames.NumField()

	if numNames > 0 {
		flagValues := reflect.ValueOf(field)
		var name string
		var value interface{}
		var fieldType string

		for i := 0; i < numNames; i++ {
			name = flagNames.Field(i).Name
			value = flagValues.Field(i)
			fieldType = fmt.Sprintf("(%+v)", flagValues.Field(i).Type())
			line := fmt.Sprintf("\t%12s %-10s %s %v\n", name, fieldType, ":", value)
			sb.WriteString(line)
		}
	} else {
		sb.WriteString("\t<empty>\n")
	}
	sb.WriteString("  }\n")
	return sb.String(), nil
}
