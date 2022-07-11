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
)

func addSubmodule(url string, path string, theme_path string) *RunError {
	// set up the submodule
	err := RunGit("submodule", "add", url, theme_path)
	if err != nil {
		return err
	}
	err = RunGit("submodule", "update", "--init")
	if err != nil {
		return err
	}
	return nil
}

func initialCommit() *RunError {
	err := RunGit("add", ".")
	if err != nil {
		return err
	}
	err = RunGit("commit", "-am", "init new notes repo")
	if err != nil {
		return err
	}
	return nil
}

func initGitRepo() *RunError {
	err := RunGit("init")
	if err != nil {
		return err
	}
	return nil
}

func UpdateSubmodule() {
	ExecGit("submodule", "update", "--remote", "--merge")
}

func ExecGit(args ...string) {
	git, err := GetGit()
	if err != nil {
		fmt.Println("Failed find git")
		fmt.Println(err)
		os.Exit(1)
	}
	realCommand := append([]string{git}, args...)
	err = Exec(realCommand...)
	if err != nil {
		fmt.Println("Failed to run git")
		fmt.Println(err)
		os.Exit(1)
	}
}

func RunGit(args ...string) *RunError {
	git, err := GetGit()
	if err != nil {
		fmt.Println("Failed find git")
		fmt.Println(err)
		os.Exit(1)
	}
	realCommand := append([]string{git}, args...)
	runErr := Run(realCommand...)
	if runErr != nil {
		return runErr
	}
	return nil
}
