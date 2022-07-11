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

	"github.com/otiai10/copy"
)

func InitRepo(path string, url string) {

	// this changes the CWD as a side-effect
	err := createBaseRepo(path)
	if err != nil {
		fmt.Printf("Could not create or use directory:\n%s\n", path)
		fmt.Println(err)
		os.Exit(1)
	}

	// make themes directory
	err = os.Mkdir("./themes", 0750)
	if err != nil && !os.IsExist(err) {
		fmt.Printf("Could not create or use directory:\n%s/themes\n", path)
		fmt.Println(err)
		os.Exit(1)
	}

	// derive the theme name and path from the URL
	pieces := strings.Split(url, "/")
	themeName := pieces[len(pieces)-1]
	themeName = strings.TrimSuffix(themeName, ".git")
	themeName = strings.TrimPrefix(themeName, "hugo-")
	themePath := fmt.Sprintf("./themes/%s", themeName)
	if err != nil {
		fmt.Printf("Could not cd to %s\n", path)
		os.Exit(1)
	}

	runErr := addSubmodule(url, path, themePath)
	if runErr != nil {
		fmt.Printf("Adding theme git submodule failed in %s\n", path)
		fmt.Println(runErr.Output)
		os.Exit(runErr.ExitCode)
	}
	err = copyExampleSite(themeName)
}

func createBaseRepo(path string) error {
	// create directory or just use it if it exists
	err := os.Mkdir(path, 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}
	// cd to directory
	os.Chdir(path)

	// init git repo
	runErr := initGitRepo()
	if runErr != nil {
		fmt.Printf("'git init' failed in %s\n", path)
		fmt.Println(runErr.Output)
		os.Exit(runErr.ExitCode)
	}
	return nil
}

func copyExampleSite(themeName string) error {
	err := copy.Copy(
		fmt.Sprintf("themes/%s/exampleSite/", themeName),
		"./",
	)
	if err != nil {
		return err
	}
	return nil
}
