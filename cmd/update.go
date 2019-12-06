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
	"calendar/internal/domain/entities"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"time"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update event",
	Long: `Update event`,
	RunE: func(cmd *cobra.Command, args []string) error {

		id, err := cmd.Flags().GetInt32("id")
		title, err := cmd.Flags().GetString("title")
		description, err := cmd.Flags().GetString("description")
		from, err := cmd.Flags().GetString("from")
		to, err := cmd.Flags().GetString("to")

		if err != nil {
			return err
		}

		timeStart, err := time.Parse("2006-01-02", from)
		timeEnd, err := time.Parse("2006-01-02", to)
		if err != nil {
			return err
		}

		eventFrom, err :=  ptypes.TimestampProto(timeStart)
		eventTo, err :=  ptypes.TimestampProto(timeEnd)
		if err != nil {
			return err
		}


		event := &entities.Event{
			Id: id,
			Title: title,
			Description:    description,
			From:  eventFrom,
			To: eventTo,
		}
		res, err := client.UpdateEvent(
			context.TODO(),
			&entities.EventUpdateRequest{
				Event: event,
			},
		)
		if err != nil {
			return err
		}
		fmt.Printf("Event updated: %s\n", res.Id)
		return nil
	},
}

func init() {

	updateCmd.Flags().Int32("id", 0, "event id")
	updateCmd.Flags().StringP("title", "n", "", "event title")
	updateCmd.Flags().StringP("description", "d", "", "event description")
	updateCmd.Flags().StringP("from", "f", "", "date from")
	updateCmd.Flags().StringP("to", "t", "", "date to")

	updateCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(updateCmd)
}
