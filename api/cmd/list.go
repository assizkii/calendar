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
	"github.com/assizkii/calendar/api/internal/domain/entities"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "event list",
	Long: `Event list for any time: day, week or month`,
	RunE: func(cmd *cobra.Command, args []string) error {

		period, err := cmd.Flags().GetString("period")
		var eventPeriod entities.Period

		if err != nil {
			return err
		}

		switch period {
			case "day":
				eventPeriod = entities.Period_DAY
			case "week":
				eventPeriod = entities.Period_WEEK
			case "month":
				eventPeriod = entities.Period_MONTH
			default:
				err = errors.New("you must set correct period")

		}

		if err != nil {
			return err
		}

		req := &entities.EventListRequest{
			Period:eventPeriod,
		}
		stream, err := client.EventList(context.Background(), req)

		for {
			// stream.Recv returns a pointer to a EventList at the current iteration
			res, err := stream.Recv()
			// If end of stream, break the loop
			if err == io.EOF {
				break
			}
			// if err, return an error
			if err != nil {
				return err
			}
			// If everything went well use the generated getter to print the event message
			fmt.Println(res.Event)
		}


		return nil

	},
}

func init() {

	listCmd.Flags().StringP("period", "p", "", "day, week or month")
	updateCmd.MarkFlagRequired("period")

	rootCmd.AddCommand(listCmd)
}
