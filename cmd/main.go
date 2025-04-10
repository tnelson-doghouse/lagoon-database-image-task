package cmd

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/uselagoon/database-image-task/internal/builder"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "database-image-task",
	Short: "A tool to help with generating database images via a task in Lagoon.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// version/build information (populated at build time by make file)
var (
	dbitName    = "database-image-task"
	dbitVersion = "0.x.x"
	dbitBuild   = ""
	goVersion   = ""
)

// version/build information command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version information",
	Run: func(cmd *cobra.Command, args []string) {
		displayVersionInfo()
	},
}

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Return the build and mtk values",
	RunE: func(cmd *cobra.Command, args []string) error {
		return builder.Run()
	},
}

func displayVersionInfo() {
	fmt.Printf("%s %s (built: %s / go %s)\n", dbitName, dbitVersion, dbitBuild, goVersion)
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(dumpCmd)
}
