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

	"github.com/mmessmore/messynotes/internal"
	"github.com/spf13/cobra"
)

// pdfCmd represents the pdf command
var pdfCmd = &cobra.Command{
	Args:  cobra.MinimumNArgs(1),
	Use:   "pdf [TITLE]",
	Short: "Create a PDF from an existing note",
	Long: `Create a pdf of an existing note by title.

Titles must be the lowercased filename without the extension as it is in the
URL when running hugo.

Shell completion provides the title in the correct format.
`,
	Run: func(cmd *cobra.Command, args []string) {
		outPath, _ := cmd.Flags().GetString("output")
		if outPath == "" {
			outPath = fmt.Sprintf("./%s.pdf", args[0])
		}

		if !internal.IsRunning() {
			fmt.Println("Hugo server does not appear to be running.")
			fmt.Println("Exiting")
			os.Exit(1)
		}

		internal.ConvertToPDF(args[0], outPath)
		fmt.Printf("PDF written to %s\n", outPath)
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
	rootCmd.AddCommand(pdfCmd)
	pdfCmd.Flags().StringP(
		"output",
		"o",
		"",
		"Output file, ./[title].pdf if unspecified",
	)
}
