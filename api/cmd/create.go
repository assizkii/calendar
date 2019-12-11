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
	"calendar/api/internal/domain/entities"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"time"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A create event",
	Long: `Create event`,
	RunE: func(cmd *cobra.Command, args []string) error{
		initClient()
		title, err := cmd.Flags().GetString("title")
		description, err := cmd.Flags().GetString("description")
		start, err := cmd.Flags().GetString("start")
		ownerId, err := cmd.Flags().GetString("owner_id")
		//duration, err := cmd.Flags().GetString("duration")
		//notify, err := cmd.Flags().GetString("notify")


		if err != nil {
			return err
		}

		timeStart, err := time.Parse("2006-01-02", start)
		//durationTime, err := time.Parse("2006-01-02", duration)
		//notifyTime, err := time.Parse("2006-01-02", notify)

		if err != nil {
			return err
		}

		eventStart, err :=  ptypes.TimestampProto(timeStart)
		//eventDuration, err :=  ptypes.TimestampProto(durationTime)
		//eventNotify, err :=  ptypes.TimestampProto(notifyTime)

		if err != nil {
			return err
		}

		event := &entities.Event{
			Id: uuid.New().String(),
			Title: title,
			Description: description,
			Start:  eventStart,
			OwnerId: ownerId,
			//Duration: eventDuration,
			//NotifyTime: eventNotify,
		}

		res, err := client.CreateEvent(
			context.TODO(),
			&entities.EventCreateRequest{
				Event: event,
			},

		)

		if err != nil {
			return err
		}
		fmt.Printf("Event created: %s\n", res.Event.Id)
		return nil
	},
}

func init() {

	createCmd.Flags().StringP("title", "t", "", "event title")
	createCmd.Flags().StringP("description", "d", "", "event description")
	createCmd.Flags().StringP("owner_id", "f", "", "event owner")
	createCmd.Flags().StringP("start", "s", "", "event start, example 2006-01-02")
	createCmd.Flags().StringP("duration", "", "", "event duration")
	createCmd.Flags().StringP("notify", "n", "", "event duration")

	createCmd.MarkFlagRequired("title")
	createCmd.MarkFlagRequired("start")

	rootCmd.AddCommand(createCmd)

}
