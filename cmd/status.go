package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Status shows the list of all running services spun via spinme",
	Long:  `Status shows the list of all running services spun via spinme`,
	Run: func(cmd *cobra.Command, args []string) {
		oo, err := getConfigs(*db)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("ID\tIP\tService\tEndpoints\n")
		for _, o := range oo {
			fmt.Printf("%s\t%s\t%s\t%v\n", o.ID, o.IP, o.Service, o.Endpoints)
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
