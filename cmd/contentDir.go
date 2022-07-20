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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// contentDirCmd represents the contentDir command
var contentDirCmd = &cobra.Command{
	Use:   "contentDir",
	Short: "Prints the directory with markdown notes to stdout",
	Long: `Prints the directory with with markdown note files to stdout

This can be used for things like:
  cd $(messynotes contentDir)
  sed -i 's/- oldcategory/- newcategory/' $(messynotes contentDir)/*

It's just shorthand for ROOT/content/notes, where ROOT is the configured root
directory of the messynotes repo.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s/content/notes\n", viper.GetString("root"))
	},
}

func init() {
	rootCmd.AddCommand(contentDirCmd)
}
