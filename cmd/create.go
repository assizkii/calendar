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
	"github.com/spf13/cobra"
	"time"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A create event",
	Long: `Create event`,
	RunE: func(cmd *cobra.Command, args []string) error{

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
			Title: title,
			Description: description,
			From:  eventFrom,
			To: eventTo,
		}
		fmt.Printf("Event created: %s\n", "START")
		res, err := client.CreateEvent(
			context.TODO(),
			&entities.EventCreateRequest{
				Event: event,
			},

		)
		fmt.Printf("Event created: %s\n", "FINISHED")
		if err != nil {
			return err
		}
		fmt.Printf("Event created: %s\n", res.Event.Id)
		return nil
	},
}

func init() {
	createCmd.Flags().StringP("title", "n", "", "event title")
	createCmd.Flags().StringP("description", "d", "", "event description")
	createCmd.Flags().StringP("from", "f", "", "date from, example 2006-01-02")
	createCmd.Flags().StringP("to", "t", "", "date to, example 2006-01-02")

	createCmd.MarkFlagRequired("title")
	createCmd.MarkFlagRequired("to")
	createCmd.MarkFlagRequired("from")

	rootCmd.AddCommand(createCmd)

}
