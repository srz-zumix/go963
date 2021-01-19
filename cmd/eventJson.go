/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

// eventJsonCmd represents the eventJson command
var eventJsonCmd = &cobra.Command{
	Use:   "eventJson",
	Short: "event json format",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		now := time.Now()
		zone, _ := now.Zone()
		event := makeEvent(getGo963Summary("user", "summary"), now.Format(dateFormat), zone, "detail")
		data, err := json.Marshal(event)
		if err != nil {
			return err
		}

		output, err := cmd.PersistentFlags().GetString("output")
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(output, data, 0755)
		if err != nil {
			log.Fatalf("WriteFile: %v", err)
		}
		fmt.Println(filepath.Abs(output))
		return err
	},
}

func init() {
	rootCmd.AddCommand(eventJsonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	eventJsonCmd.PersistentFlags().StringP("output", "o", "event.json", "output json path")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// eventJsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
