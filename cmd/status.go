/*
Copyright © 2019 Tanmay Chaudhry <tanmay.chaudhry@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
		fmt.Printf("ID\tIP\tService\n")
		for _, o := range oo {
			fmt.Printf("%s\t%s\t%s\n", o.ID, o.IP, o.Service)
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
