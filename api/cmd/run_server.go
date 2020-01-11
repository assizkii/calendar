/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"github.com/assizkii/calendar/api/internal/adapters/servers/grpc_server"

	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runServerCmd represents the runServer command
var runServerCmd = &cobra.Command{
	Use:   "run_server",
	Short: "start grpc server",
	Long:  `start grpc server`,
	Run: func(cmd *cobra.Command, args []string) {
		grpc_server.StartServer()
	},
}

func init() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".calendar" (without extension).
		viper.AddConfigPath("api/configs")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	rootCmd.AddCommand(runServerCmd)
}
