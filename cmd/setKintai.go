package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

// setKintaiCmd represents the setKintai command
var setKintaiCmd = &cobra.Command{
	Use:   "setKintai <summary> [comment]",
	Short: "create / edit Kintai",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		err := cobra.MinimumNArgs(1)(cmd, args)
		if err != nil {
			return err
		}
		err = cobra.MaximumNArgs(2)(cmd, args)
		if err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := createClient()
		if err != nil {
			return err
		}

		user, err := cmd.PersistentFlags().GetString("user")
		if err != nil {
			return err
		}
		date, err := cmd.PersistentFlags().GetString("date")
		if err != nil {
			return err
		}
		now := time.Now()
		zone, _ := now.Zone()

		d, err := time.ParseInLocation(dateFormat, date, now.Location())
		if err != nil {
			return err
		}
		summary := args[0]
		description := ""
		if len(args) > 1 {
			description = args[1]
		}
		title := getGo963Summary(user, summary)
		ev := makeEvent(title, d.Format(dateFormat), zone, description)

		_, err = setEvent(client, user, d, ev)
		return err
	},
}

func init() {
	rootCmd.AddCommand(setKintaiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	setKintaiCmd.PersistentFlags().StringP("user", "u", "", "Username")
	setKintaiCmd.PersistentFlags().StringP("date", "d", "", "Date (YYYY-MM-DD)")

	setKintaiCmd.MarkPersistentFlagRequired("user")
	setKintaiCmd.MarkPersistentFlagRequired("date")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setKintaiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
