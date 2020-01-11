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
	"context"
	"fmt"
	"calendar/entities"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete event",
	Long:  `Delete event`,
	RunE: func(cmd *cobra.Command, args []string) error {
		initClient()
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			return err
		}

		req := &entities.EventDeleteRequest{
			Id: id,
		}

		_, err = client.DeleteEvent(context.Background(), req)
		if err != nil {
			return err
		}

		fmt.Printf("Succesfully deleted the event with id %s\n", id)
		return nil
	},
}

func init() {
	deleteCmd.Flags().StringP("id", "i", "", "event id")
	deleteCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(deleteCmd)
}
