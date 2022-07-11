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
package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func Edit(path string) {
	editor, err := GetEditor()
	if err != nil {
		fmt.Println("ERROR: Could not find an editor to execute")
		os.Exit(1)
	}

	err = Exec(editor, path)
	if err != nil {
		fmt.Printf("ERROR: failed to open editor on %s\n", path)
		fmt.Println(err)
		os.Exit(1)
	}
}

func EditByTitle(title string) {
	filename := fmt.Sprintf("%s/content/notes/%s.md",
		viper.GetString("root"),
		strings.ReplaceAll(strings.ToLower(title), " ", "-"))
	Edit(filename)
}

func GetTitles(toComplete string) []string {
	var titles []string
	f, err := os.Open(viper.GetString("root") + "/content/notes")
	if err != nil {
		return nil
	}
	files, err := f.Readdir(0)
	if err != nil {
		return nil
	}

	for _, file := range files {
		if strings.Contains(
			strings.ToLower(file.Name()),
			strings.ToLower(toComplete),
		) {
			titles = append(
				titles,
				strings.ToLower(
					strings.TrimSuffix(file.Name(), ".md"),
				),
			)
		}
	}
	return titles
}

func Create(title string) {
	filename := fmt.Sprintf("notes/%s.md",
		strings.ReplaceAll(strings.ToLower(title), " ", "-"))

	hugo, err := GetHugo()

	run_err := Run(hugo, "new", filename)
	if err != nil {
		fmt.Println("ERROR: Failed to run hugo")
		fmt.Println(run_err.Output)
		os.Exit(run_err.ExitCode)
	}
	filename = fmt.Sprintf("content/%s", filename)

	Edit(filename)
}

func Todo() {
	Edit("content/notes/TODO.md")
}
