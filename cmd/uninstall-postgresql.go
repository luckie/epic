// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/chrisbenson/epic/pkg/epic"
	"github.com/spf13/cobra"
	"os"
)

// uninstallPostreSQLCmd represents the uninstallPostreSQL command
var uninstallPostgreSQLCmd = &cobra.Command{
	Use:   "postgresql",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		verifyUninstallFlags()
		fmt.Println("Removing Epic database.")
		err := epic.UninstallPostgreSQLDatabase(adminUser, adminPassword, server)
		resetCreds()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Epic database successfully removed.")
		}
	},
}

func init() {
	uninstallCmd.AddCommand(uninstallPostgreSQLCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uninstallPostreSQLCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uninstallPostreSQLCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	uninstallPostgreSQLCmd.Flags().StringVarP(&adminUser, "adminUser", "u", "", "--adminUser admin, -u admin")
	uninstallPostgreSQLCmd.Flags().StringVarP(&adminPassword, "adminPassword", "p", "", "--adminPassword password1, -p password1")
	uninstallPostgreSQLCmd.Flags().StringVarP(&server, "server", "s", "", "--server server, -s server")
	uninstallPostgreSQLCmd.MarkFlagRequired("adminUser")
	uninstallPostgreSQLCmd.MarkFlagRequired("adminPassword")
	uninstallPostgreSQLCmd.MarkFlagRequired("server")

}

func verifyUninstallFlags() {

	var flagFail bool
	if adminUser == "" {
		flagFail = true
		fmt.Println("The required 'adminUser' flag has not been set. (e.g. --adminUser admin, -u admin)")
	}
	if adminPassword == "" {
		flagFail = true
		fmt.Println("The required 'adminPassword' flag has not been set. (e.g. --adminPassword password1, -p password1)")
	}
	if server == "" {
		flagFail = true
		fmt.Println("The required 'server' flag has not been set. (e.g. --server server, -s server)")
	}
	if flagFail == true {
		fmt.Println("Correct usage is like this example: epic uninstall postgresql --adminUser admin --adminPassword password1 --server server")
		fmt.Println("Or this example: epic uninstall postgresql --u admin --p password1 --s server")
		os.Exit(1)
	}
}

