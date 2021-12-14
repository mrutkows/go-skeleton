/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"fmt"
	"os"

	"github.com/mrutkows/go-skeleton/log"
	"github.com/mrutkows/go-skeleton/utils"
	"github.com/spf13/cobra"
)

var ProjectLogger *log.MiniLogger

const (
	FLAG_TRACE                 = "trace"
	FLAG_TRACE_SHORT           = "t"
	FLAG_DEBUG                 = "debug"
	FLAG_DEBUG_SHORT           = "d"
	FLAG_FILENAME_INPUT        = "input-file"
	FLAG_FILENAME_INPUT_SHORT  = "i"
	FLAG_FILENAME_OUTPUT       = "output-file"
	FLAG_FILENAME_OUTPUT_SHORT = "o"
)

var rootCmd = &cobra.Command{
	Use:           utils.Flags.Project,
	SilenceErrors: false, // TODO: investigate if we should use
	SilenceUsage:  false, // TODO: investigate if we should use
	Short:         "Software Bill-of-Materials (SBOM) base utility.",
	Long:          "This utility serves as centralized command line interface into various Software Bill-of-Materials (SBOM) helper utilities.",
	RunE:          RootCmdImpl,
}

// initialize the module; primarily, initialize cobra
func init() {
	ProjectLogger = log.NewLogger(log.TRACE)
	ProjectLogger.Enter()

	// Tell Cobra what our Cobra "init" call back method is
	cobra.OnInitialize(initConfig)

	// Declare top-level, persistent flags and where to place the post-parse values
	// TODO: move command help strings to (centralized) constants for better editing/translation across all files
	//rootCmd.PersistentFlags().BoolVarP(nil, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&utils.Flags.Trace, FLAG_TRACE, FLAG_TRACE_SHORT, false, "enable trace logging")
	rootCmd.PersistentFlags().BoolVarP(&utils.Flags.Debug, FLAG_DEBUG, FLAG_DEBUG_SHORT, false, "enable debug logging")
	rootCmd.PersistentFlags().StringVarP(&utils.Flags.InputFile, FLAG_FILENAME_INPUT, FLAG_FILENAME_INPUT_SHORT, "", "input filename")
	rootCmd.PersistentFlags().StringVarP(&utils.Flags.OutputFile, FLAG_FILENAME_OUTPUT, FLAG_FILENAME_OUTPUT_SHORT, "", "output filename")
	ProjectLogger.Exit()
}

func initConfig() {
	ProjectLogger.Enter()

	// Update log level
	if utils.Flags.Debug {
		ProjectLogger.SetLevel(log.DEBUG)
	} else if utils.Flags.Trace {
		// debug level implies trace
		ProjectLogger.SetLevel(log.TRACE)
	}

	// Print global flags in debug mode
	flagInfo, err := log.FormatStruct("Flags", utils.Flags)
	if err != nil {
		ProjectLogger.Error(err.Error())
	} else {
		ProjectLogger.Debug(flagInfo)
	}

	// Print logger settings in debug mode
	logInfo, err2 := log.FormatStruct("Flags", utils.Flags)
	if err2 != nil {
		ProjectLogger.Error(err2.Error())
	} else {
		ProjectLogger.Debug(logInfo)
	}
	ProjectLogger.Exit()
}

func RootCmdImpl(cmd *cobra.Command, args []string) error {
	ProjectLogger.Enter()
	//fmt.Printf("cmd: %+v\nargs: %v\n", cmd, args)
	ProjectLogger.Exit()
	return nil
}

func Execute() {
	// instead of creating a dependency on the "main" module
	ProjectLogger.Enter()
	if err := rootCmd.Execute(); err != nil {
		// TODO: use log errors
		// TODO: invalid command (empty); display help
		rootCmd.Help()
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	ProjectLogger.Exit()
}
