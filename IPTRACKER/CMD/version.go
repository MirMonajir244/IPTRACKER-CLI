package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "To Print the version of the Application",
	Long:  `IPTRACKER version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
