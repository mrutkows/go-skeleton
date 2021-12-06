package cmd

import (
	"fmt"
	"os"

	"github.com/mrutkows/sbom-utility/log"
	"github.com/mrutkows/sbom-utility/utils"
	"github.com/spf13/cobra"
)

const (
	FLAG_VERBOSE               = "verbose"
	FLAG_VERBOSE_SHORT         = "v"
	FLAG_FILENAME_INPUT        = "input-file"
	FLAG_FILENAME_INPUT_SHORT  = "i"
	FLAG_FILENAME_OUTPUT       = "output-file"
	FLAG_FILENAME_OUTPUT_SHORT = "o"
)

var rootCmd = &cobra.Command{
	Use:           "spdx-parser",
	SilenceErrors: true,
	SilenceUsage:  true,
	Short:         "Root Short Desc.",
	Long:          "Root Long Desc.",
	RunE:          RootCmdImp,
}

var loggers log.MyLog

func RootCmdImp(cmd *cobra.Command, args []string) error {

	//fmt.Printf("cmd: %+v\nargs: %v\n", cmd, args)
	return nil
}

// initialize the module; primarily, initialize cobra
func init() {
	loggers.Enter()

	// Tell Cobra what our Cobra "init" call back method is
	cobra.OnInitialize(initConfig)

	// Declare top-level, persistent flags and where to place the post-parse values
	rootCmd.PersistentFlags().BoolVarP(&utils.Flags.Verbose, FLAG_VERBOSE, FLAG_VERBOSE_SHORT, false, "Verbose output (i.e., INFO")
	rootCmd.PersistentFlags().StringVarP(&utils.Flags.InputFile, FLAG_FILENAME_INPUT, FLAG_FILENAME_INPUT_SHORT, "", "Input filename")
	rootCmd.PersistentFlags().StringVarP(&utils.Flags.OutputFile, FLAG_FILENAME_OUTPUT, FLAG_FILENAME_OUTPUT_SHORT, "", "Output filename")
	loggers.Exit()
}

func initConfig() {
	loggers.Enter()
	err := log.DumpStruct("utils.Flags", utils.Flags)
	//err := log.DumpStruct("os.Args", os.Args)
	if err != nil {
		loggers.Error("structName", err.Error())
	}
	loggers.Exit()
}

func Execute(programLogger log.MyLog) {
	// instead of creating a dependency on the "main" module
	loggers = programLogger
	loggers.Enter()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	loggers.Exit()
}
