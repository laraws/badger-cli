package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show badger cli version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version", "V4")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
