/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// directory, err := os.Getwd()
		// _CheckIfError(err)
		// fmt.Println(directory)
		// r, err := git.PlainOpen(directory)
		// _CheckIfError(err)
		// w, err := r.Worktree()
		// _CheckIfError(err)
		// status, err := w.Status()
		// _CheckIfError(err)
		// fmt.Println(status)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		message, _ := reader.ReadString('\n')
		format := viper.Get("format")
		story := viper.Get("story")
		pair := viper.Get("pair")
		if format, ok := format.(string); ok {
			if story, ok := story.(string); ok {
				if pair, ok := pair.(string); ok {
					fmt.Println(fmt.Sprintf(format, story, pair, message))
				}
			}
		}
	},
}

func _CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func init() {
	rootCmd.AddCommand(commitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
