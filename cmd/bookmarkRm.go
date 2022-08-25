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

	"github.com/mmessmore/messynotes/internal"
	"github.com/spf13/cobra"
)

// bookmarkRmCmd represents the rm command
var bookmarkRmCmd = &cobra.Command{
	Use:   "rm -t TITLE",
	Short: "Remove a bookmark",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		title, _ := cmd.Flags().GetString("title")
		success := internal.RmBookmark(category, title)
		if success {
			fmt.Println("Bookmark removed")
		} else {
			fmt.Println("Bookmark does not exist")
		}
	},
	ValidArgsFunction: func(
		cmd *cobra.Command,
		args []string,
		toComplete string,
	) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	bookmarkCmd.AddCommand(bookmarkRmCmd)
	bookmarkRmCmd.Flags().StringP("category", "c", "main", "Category of bookmark")
	bookmarkRmCmd.RegisterFlagCompletionFunc("category", func(
		cmd *cobra.Command,
		args []string,
		toComplete string,
	) ([]string, cobra.ShellCompDirective) {
		return internal.CategoryComplete(toComplete),
			cobra.ShellCompDirectiveNoFileComp | cobra.ShellCompDirectiveNoSpace
	},
	)
	bookmarkRmCmd.Flags().StringP("title", "t", "", "Title of bookmark")
	bookmarkRmCmd.MarkFlagRequired("title")
	bookmarkRmCmd.RegisterFlagCompletionFunc("title", func(
		cmd *cobra.Command,
		args []string,
		toComplete string,
	) ([]string, cobra.ShellCompDirective) {
		category, err := cmd.Flags().GetString("category")
		if err != nil || category == "" {
			category = "main"
		}
		return internal.BookmarkComplete(category, toComplete),
			cobra.ShellCompDirectiveNoFileComp | cobra.ShellCompDirectiveNoSpace
	},
	)
}
