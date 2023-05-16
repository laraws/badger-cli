package cmd

import (
	"fmt"
	"github.com/laraws/badger-cli/v4/pkg/badger"
	"github.com/spf13/cobra"
	"log"
)

var getCmd = &cobra.Command{
	Use:   "get [key]...",
	Short: "Get content of a specific key",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := badger.Open(cmd.Flag("dir").Value.String())
		if err != nil {
			log.Fatalln(err)
		}
		defer db.Close()

		values, err := db.Get(args...)
		if err != nil {
			log.Fatalln(err)
		}

		for _, v := range values {
			fmt.Println(v)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
