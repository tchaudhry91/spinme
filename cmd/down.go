package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tchaudhry91/spinme/spin"
)

var id string

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Bring down the given container",
	Run: func(cmd *cobra.Command, args []string) {
		err := spin.SlashID(context.Background(), id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = removeConfig(*db, id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(downCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	downCmd.Flags().StringVarP(&id, "id", "i", "", "ID of the container to slash")
	err := downCmd.MarkFlagRequired("id")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
