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

	"github.com/mrutkows/go-skeleton/utils"
	"github.com/spf13/cobra"
)

func init() {
	ProjectLogger.Enter()
	rootCmd.AddCommand(versionCmd)
	ProjectLogger.Exit()
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "display program, binary and version information",
	Long:  "display program, binary and version information in SemVer format (e.g., `<project> version <x.y.z>`)",
	Run: func(cmd *cobra.Command, args []string) {
		ProjectLogger.Enter()
		// TODO: print cpu architecture of binary (e.g., go version go1.16.3 darwin/amd64)
		fmt.Printf("%s version %s\n", utils.Flags.Project, utils.Flags.Version)
		ProjectLogger.Enter()
	},
}
