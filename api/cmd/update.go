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
	entities "github.com/assizkii/calendar/api/internal/domain/entities"
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
		initClient()
		id, err := cmd.Flags().GetString("id")
		title, err := cmd.Flags().GetString("title")
		description, err := cmd.Flags().GetString("description")
		start, err := cmd.Flags().GetString("start")
		ownerId, err := cmd.Flags().GetString("owner_id")
		duration, err := cmd.Flags().GetString("duration")
		notify, err := cmd.Flags().GetString("notify")


		if err != nil {
			return err
		}

		timeStart, err := time.Parse("2006-01-02", start)
		durationTime, err := time.Parse("2006-01-02", duration)
		notifyTime, err := time.Parse("2006-01-02", notify)

		if err != nil {
			return err
		}

		eventStart, err :=  ptypes.TimestampProto(timeStart)
		eventDuration, err :=  ptypes.TimestampProto(durationTime)
		eventNotify, err :=  ptypes.TimestampProto(notifyTime)

		if err != nil {
			return err
		}

		event := &entities.Event{
			Id: id,
			Title: title,
			Description: description,
			Start:  eventStart,
			OwnerId: ownerId,
			Duration: eventDuration,
			NotifyTime: eventNotify,
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
		fmt.Printf("Event created: %s\n", res.Id)
		return nil
	},
}

func init() {

	updateCmd.Flags().StringP("id", "i", "", "event id")
	updateCmd.Flags().StringP("title", "t", "", "event title")
	updateCmd.Flags().StringP("description", "d", "", "event description")
	updateCmd.Flags().StringP("owner_id", "f", "", "event owner")
	updateCmd.Flags().StringP("start", "s", "", "date to")
	updateCmd.Flags().StringP("duration", "", "", "event duration")
	updateCmd.Flags().StringP("notify", "n", "", "event duration")

	updateCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(updateCmd)
}
