package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// createEventsCmd represents the createEvents command
var createEventsCmd = &cobra.Command{
	Use:   "createEvents <event json path>",
	Short: "create event from events.json",
	Long:  `create event from events.json`,
	Args: func(cmd *cobra.Command, args []string) error {
		err := cobra.MinimumNArgs(1)(cmd, args)
		if err != nil {
			return err
		}
		_, err = os.Stat(args[0])
		return err
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := createClient()
		if err != nil {
			return err
		}
		_, err = createEventFromJson(client, args[0])
		return err
	},
}

func init() {
	rootCmd.AddCommand(createEventsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createEventsCmd.PersistentFlags().St("json", "", "events.json path")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createEventsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
