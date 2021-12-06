package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "short",
	Long:  "display program, binary and version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO")
		//fmt.Printf("Welcome to the %s! Version `%s` (%s)\n", main.Project, Version, Binary)
	},
}
