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
	"context"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/tchaudhry91/spinme/spin"
)

var (
	service string
	image   string
	name    string
	env     []string
	tag     string
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Start a particular service",
	Run: func(cmd *cobra.Command, args []string) {
		conf := spin.SpinConfig{
			Image: image,
			Tag:   tag,
			Name:  name,
			Env:   env,
		}
		spinner, err := spin.SpinnerFrom(service)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		out, err := spinner(context.Background(), &conf)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = storeConfig(*db, out)
		if err != nil {
			fmt.Println(err)
			spew.Dump(out)
			os.Exit(1)
		}
		spew.Dump(out)
	},
}

func init() {
	rootCmd.AddCommand(upCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	upCmd.Flags().StringVarP(&service, "service", "s", "", "Service to spin up. Inbuilts: mongo/postgres/mysql/redis")
	upCmd.Flags().StringVarP(&name, "name", "n", "", "Override docker container name to use")
	upCmd.Flags().StringVarP(&image, "image", "i", "", "Override docker image to use")
	upCmd.Flags().StringVarP(&tag, "tag", "t", "", "Override docker image tag to use")
	upCmd.Flags().StringArrayVarP(&env, "env", "a", []string{}, "Environment variables to pass in, e.g --env PG_PASSWORD=1231")

	err := upCmd.MarkFlagRequired("service")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
