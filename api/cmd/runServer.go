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
	"calendar/api/internal/adapters/servers/grpc"
	"calendar/api/internal/adapters/storages/inmemory"
	"fmt"
	"github.com/spf13/cobra"
)

// runServerCmd represents the runServer command
var runServerCmd = &cobra.Command{
	Use:   "runServer",
	Short: "start grpc server",
	Long: `start grpc server`,
	Run: func(cmd *cobra.Command, args []string) {
		//db := database.GetDBConnection()
		//database.RunMigrations(db)
		//defer db.Close()
		fmt.Println("Starting Calendar Service Client")
		storage :=  inmemory.New()

		grpc.StartServer(storage)
	},
}

func init() {
	rootCmd.AddCommand(runServerCmd)
}