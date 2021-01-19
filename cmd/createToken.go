package cmd

import (
	"github.com/spf13/cobra"
)

// createTokenCmd represents the createToken command
var createTokenCmd = &cobra.Command{
	Use:   "createToken",
	Short: "create OAuth 2.0 cache token",
	Long:  `create OAuth 2.0 cache token`,
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := createClient()
		return err
	},
}

func init() {
	rootCmd.AddCommand(createTokenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createTokenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createTokenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
