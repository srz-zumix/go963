package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// listKintaiCmd represents the listKintai command
var listKintaiCmd = &cobra.Command{
	Use:   "listKintai",
	Short: "list Kintai",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := createClient()
		if err != nil {
			return err
		}

		user, _ := cmd.PersistentFlags().GetString("user")
		date, err := cmd.PersistentFlags().GetString("date")
		if err != nil {
			return err
		}
		now := time.Now()

		d, err := time.ParseInLocation(dateFormat, date, now.Location())
		if err != nil {
			return err
		}
		events, err := listEvent(client, user, d)
		if err != nil {
			return err
		}

		for _, event := range events.Items {
			fmt.Println(getKintaiSummary(event.Summary))
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(listKintaiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	listKintaiCmd.PersistentFlags().StringP("user", "u", "", "list date")
	listKintaiCmd.PersistentFlags().StringP("date", "d", "", "Username")

	listKintaiCmd.MarkPersistentFlagRequired("date")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listKintaiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
