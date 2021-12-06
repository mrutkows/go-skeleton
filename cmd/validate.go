package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xeipuuv/gojsonschema"
)

const (
	SCHEMA_SPDX_2_2_LOCAL = "file://schema/spdx/2.2/spdx-schema.json"
)

var (
	supportedSchemas = []string{SCHEMA_SPDX_2_2_LOCAL}
)

func init() {
	rootCmd.AddCommand(validateCmd)
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "",
	Long:  "validate input file's schema",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Validating...")
	},
}

func isSchemaSupported(schema string) bool {
	for _, s := range supportedSchemas {

		if schema == s {
			return true
		}
	}
	return false
}

func Validate(schema string, document string) (bool, error) {

	if !isSchemaSupported(schema) {
		return false, fmt.Errorf("schema [%s] not supported", schema)
	}

	loader := gojsonschema.NewReferenceLoader(SCHEMA_SPDX_2_2_LOCAL)

	// create a reusable schema object (to validate multiple documents)
	_, err := gojsonschema.NewSchema(loader)

	if err != nil {
		return false, err
	}

	return true, nil
}
