package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// delKintaiCmd represents the delKintai command
var delKintaiCmd = &cobra.Command{
	Use:   "delKintai",
	Short: "delete Kintai",
	Long:  ``,
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
		allow, err := cmd.PersistentFlags().GetBool("allow")
		if err != nil {
			return err
		}
		now := time.Now()

		d, err := time.ParseInLocation(dateFormat, date, now.Location())
		if err != nil {
			return err
		}
		err = deleteEvent(client, user, d)
		if err != nil {
			if !allow {
				return err
			}
		} else {
			fmt.Println("finish delete")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(delKintaiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	delKintaiCmd.PersistentFlags().StringP("user", "u", "", "Username")
	delKintaiCmd.PersistentFlags().StringP("date", "d", "", "Date (YYYY-MM-DD)")
	delKintaiCmd.PersistentFlags().Bool("allow", false, "allow event not found")

	delKintaiCmd.MarkPersistentFlagRequired("user")
	delKintaiCmd.MarkPersistentFlagRequired("date")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delKintaiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
