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

	"github.com/mrutkows/go-skeleton/utils"
	"github.com/spf13/cobra"
)

func init() {
	ProjectLogger.Enter()
	rootCmd.AddCommand(validateCmd)
	ProjectLogger.Exit()
}

var validateCmd = &cobra.Command{
	Use:   "validate -i <input-sbom.json>",
	Short: "validate input file against its declared SBOM schema.",
	Long:  "validate input file against its declared SBOM schema, if detectable and supported.",
	Run: func(cmd *cobra.Command, args []string) {
		ProjectLogger.Enter()
		// TODO: remove when execution call order satisfactory
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		ProjectLogger.Exit()
	},
	RunE: validateCmdImpl,
}

func validateCmdImpl(cmd *cobra.Command, args []string) error {
	ProjectLogger.Enter()
	isValid, err := Validate()
	if err != nil {
		ProjectLogger.Error(err)
		os.Exit(-3)
	}
	ProjectLogger.Info(fmt.Sprintf("Document %s: valid=[%t]", utils.Flags.InputFile, isValid))
	ProjectLogger.Exit()
	return nil
}

func Validate() (bool, error) {
	ProjectLogger.Enter()


	ProjectLogger.Exit(true)
	return true, nil
}
