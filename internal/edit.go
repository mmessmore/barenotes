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
	"path/filepath"
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
	_, filename, exists := NotePathsByTitle(title)
	if !exists {
		fmt.Printf("Note does not exist: '%s'\n", title)
		/*
			I can't find a way to use fs.ErrNotExist or syscall.ENOENT
			to get back to an integer, so this is hardcoded.
		*/
		os.Exit(2)
	}
	Edit(filename)
}

func GetTitles(toComplete string) []string {
	var titles []string
	f, err := os.Open(filepath.Join(
		viper.GetString("root"),
		"content",
		"notes"),
	)
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

	filename, fullFilename, exists := NotePathsByTitle(title)
	if exists {
		fmt.Printf("File already exists, editing:\n\t%s\n", fullFilename)
		Edit(fullFilename)
	} else {
		hugo, err := GetHugo()

		run_err := Run(hugo, "new", filename)
		if err != nil {
			fmt.Println("ERROR: Failed to run hugo")
			fmt.Println(run_err.Output)
			os.Exit(run_err.ExitCode)
		}

		Edit(fullFilename)
	}

}

func Todo() {
	Edit(filepath.Join(
		viper.GetString("root"),
		"content",
		"notes",
		"TODO.md"),
	)
}

func NotePathsByTitle(title string) (string, string, bool) {
	filename := filepath.Join(
		"notes",
		fmt.Sprintf(
			"%s.md",
			strings.ReplaceAll(strings.ToLower(title), " ", "-"),
		),
	)

	fullFilename := filepath.Join(
		viper.GetString("root"),
		"content",
		filename,
	)

	if _, err := os.Stat(fullFilename); err == nil {
		return filename, fullFilename, true
	} else {
		return filename, fullFilename, false
	}
}
