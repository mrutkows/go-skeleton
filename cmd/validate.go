package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
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

func Validate(schema string, document string) (bool, error) {

	return true, nil
}
