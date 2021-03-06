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
	"os/user"

	"github.com/spf13/cobra"
)

var file string

// kubeconfigCmd represents the kubeconfig command
var KubeconfigCmd = &cobra.Command{
	Use: "kubeconfig",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("Run `boatswain kubeconfig --help` for list of available subcommands.")

	},
}

func init() {
	RootCmd.AddCommand(KubeconfigCmd)
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	defaultConfig := usr.HomeDir + "/.kube/config"
	KubeconfigCmd.PersistentFlags().StringVarP(&file, "file", "f", defaultConfig, "Target kubeconfig file.")
}
