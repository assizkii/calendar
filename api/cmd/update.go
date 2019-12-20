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
	"github.com/assizkii/calendar/api/internal/domain/entities"
	"github.com/golang/protobuf/ptypes"
	"time"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update event",
	Long:  `Update event`,
	RunE: func(cmd *cobra.Command, args []string) error {
		initClient()
		id, err := cmd.Flags().GetString("id")
		title, err := cmd.Flags().GetString("title")
		description, err := cmd.Flags().GetString("description")
		start, err := cmd.Flags().GetString("start")
		owner, err := cmd.Flags().GetInt32("owner")
		end, err := cmd.Flags().GetString("end")

		if err != nil {
			return err
		}

		timeStart, err := time.Parse("2006-01-02", start)
		timeEnd, err := time.Parse("2006-01-02", end)

		if err != nil {
			return err
		}

		eventStart, err := ptypes.TimestampProto(timeStart)
		eventEnd, err :=  ptypes.TimestampProto(timeEnd)

		if err != nil {
			return err
		}

		event := &entities.Event{
			Id:          id,
			Title:       title,
			Description: description,
			Start:       eventStart,
			OwnerId:     owner,
			EndTime: eventEnd,
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

	updateCmd.Flags().StringP("id", "i", "", "event id")
	updateCmd.Flags().StringP("title", "t", "", "event title")
	updateCmd.Flags().StringP("description", "d", "", "event description")
	updateCmd.Flags().Int32("owner", 0, "event owner")
	updateCmd.Flags().StringP("start", "s", "", "event start, example 2006-01-02")
	updateCmd.Flags().StringP("end", "e", "", "event end, example 2006-01-02")

	updateCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(updateCmd)
}
