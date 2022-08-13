/*
Copyright © 2022 Mike Messmore <mike@messmore.org>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mmessmore/messynotes/internal"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Args:  cobra.MinimumNArgs(1),
	Use:   "init [new directory]",
	Short: "Initialize new site/repository",
	Long: `Bootstrap a new site

This creates the directory, sets up a git repo using the exampleSite and
adds the theme repository as a submodule.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// if cfgFile isn't set, fall to default
		// this feels gross, but Viper supports multiple sources
		// so I can't get the path from it
		if cfgFile == "" {
			home, _ := os.UserHomeDir()
			cfgFile = filepath.Join(home, ".messynotes.yaml")
		}

		repoPath, _ := filepath.Abs(args[0])
		themeUrl, _ := cmd.Flags().GetString("themeUrl")

		internal.InitRepo(repoPath, themeUrl)

		fmt.Printf("Repo successfully created in %s!\n", repoPath)
		internal.PromptToSaveConfig(repoPath, cfgFile)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringP("themeUrl",
		"T",
		"https://github.com/mmessmore/hugo-messynotes.git",
		"theme repository git URL")
	initCmd.Flags().StringP("directory", "d", "./notes",
		"directory to create repository in")
}
