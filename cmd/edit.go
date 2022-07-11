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
	"strings"

	"github.com/mmessmore/barenotes/internal"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Args:  cobra.MinimumNArgs(1),
	Use:   "edit [TITLE]",
	Short: "Edit an existing note",
	Long: `Run editor on existing note by title.

Titles are the filename without the '.md' extension, case insensitive, and
spaces replaced by '-'s if given.

Shell completion provides the lowercased filenames without the extension.
`,
	Run: func(cmd *cobra.Command, args []string) {
		title := strings.Join(args, " ")
		internal.EditByTitle(title)
	},
	ValidArgsFunction: func(
		cmd *cobra.Command,
		args []string,
		toComplete string,
	) ([]string, cobra.ShellCompDirective) {
		return internal.GetTitles(toComplete),
			cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
