package utils

import (
	"fmt"

	"github.com/mrutkows/sbom-utility/log"
)

type MyFlags struct {
	Verbose      bool // Verbose logging
	InputFile    string
	InputFormat  string
	OutputFile   string
	OutputFormat string
}

// format and output the MyFlags struct as a string using Go's Stringer interface
func (flags *MyFlags) String() string {
	value, err := log.FormatStruct("utils.Flags", flags)

	if err != nil {
		return fmt.Sprintf("%s\n", err.Error())
	}
	return value
}

var Flags MyFlags
