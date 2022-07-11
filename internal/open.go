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

// TODO: this shouldn't be hardcoded
var URL = "http://localhost:1313"

func Open() {

	path, err := GetBrowser()
	if err != nil {
		fmt.Println("ERROR: could not find command to open browser")
		os.Exit(0)
	}

	args := []string{path, URL}

	proc, err := os.StartProcess(path, args, &os.ProcAttr{})
	if err != nil {
		fmt.Println("ERROR: could not open browser")
		fmt.Println(err)
		os.Exit(1)
	}
	proc.Release()
}
