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
	"strings"
	"time"

	cfg "github.com/pmcfadden/gcommit/internal/models"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

var dryRun bool

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Make a commit",
	Long: `Create a commit using a format defined in your config file

	It will ask for your current story, pair and message`,
	Run: func(cmd *cobra.Command, args []string) {
		directory, err := os.Getwd()
		_CheckIfError(err)
		r, err := git.PlainOpen(directory)
		_CheckIfError(err)
		w, err := r.Worktree()
		_CheckIfError(err)
		gitInfo := cfg.ReadAndSetConfig(os.Stdin)
		fmt.Print("Enter text: ")
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		message = strings.TrimSuffix(message, "\n")

		fullMessage := fmt.Sprintf(gitInfo.Format, gitInfo.Story, gitInfo.Pair, message)
		if dryRun {
			fmt.Println(fullMessage)
		} else {
			_, err = w.Commit(fullMessage, &git.CommitOptions{
				Author: &object.Signature{
					Name:  gitInfo.Name,
					Email: gitInfo.Email,
					When:  time.Now(),
				},
			})
		}
		_CheckIfError(err)
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
	commitCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Don't actually make a commit, just print the full message")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
