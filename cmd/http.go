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
	"os"
	"strconv"

	"github.com/chrisbenson/epic/pkg/epic"
	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		verifyServeHTTPFlags()
		portInt, _ := strconv.Atoi(port)
		epic.ServeHTTP(portInt, host, epicPassword, server)
		resetCreds()
	},
}

func init() {
	serveCmd.AddCommand(httpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	httpCmd.Flags().StringVarP(&port, "port", "", "", "--port 8080, -n 8080")
	httpCmd.Flags().StringVarP(&host, "host", "", "", "--host example.com, -h example.com")
	httpCmd.Flags().StringVarP(&epicPassword, "epicPassword", "", "", "--epicPassword password, -e password")
	httpCmd.Flags().StringVarP(&server, "server", "", "", "--server server, -s server")
	//httpCmd.MarkFlagRequired("port")
	//httpCmd.MarkFlagRequired("host")
	//httpCmd.MarkFlagRequired("epicPassword")
	//httpCmd.MarkFlagRequired("server")

}


func verifyServeHTTPFlags() {

	var flagFail bool
	if port == "" {
		flagFail = true
		fmt.Println("The required 'port' flag has not been set. (e.g. --port 8080, -n 8080)")
	}
	if host == "" {
		flagFail = true
		fmt.Println("The required 'host' flag has not been set. (e.g. --host example.com, -h example.com)")
	}
	//if epicUser == "" {
	//	flagFail = true
	//	fmt.Println("The required 'epicUser' flag has not been set. (e.g. --epicUser admin, -u admin)")
	//}
	if epicPassword == "" {
		flagFail = true
		fmt.Println("The required 'epicPassword' flag has not been set. (e.g. --epicPassword password, -p password)")
	}
	if server == "" {
		flagFail = true
		fmt.Println("The required 'server' flag has not been set. (e.g. --server server, -s server)")
	}
	if flagFail == true {
		fmt.Println("Correct usage is like this example: epic serve http --port 8080 --host example.com --epicUser admin --adminPassword password --server server")
		fmt.Println("Or this example: epic serve http -n 8080 -h example.com -u admin -p password -s server")
		os.Exit(1)
	}
}
