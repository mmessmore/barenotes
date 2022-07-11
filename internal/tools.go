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
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func CD() {
	dir := viper.GetString("root")
	err := os.Chdir(dir)
	if err != nil {
		fmt.Printf("ERROR: Could not change directory to %s\n", dir)
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetBrowser() (string, error) {
	// Fall through config, system defaults, then look at random browsers
	browsers := []string{
		viper.GetString("browser"), //config
		"open",                     //macos default
		"xdg-open",                 //XDG default
		"x-www-browser",            // alternatives default
		"firefox",                  // try firefox?
		"google-chrome-stable",     // try chrome?
		"chromium-browser",         // try chromium?
	}
	browser := ""
	var err error
	for _, cmd := range browsers {
		browser, err = exec.LookPath(cmd)
		if err == nil && browser != "" {
			return browser, nil
		}
	}
	return browser, errors.New("Could not find a way to open a browser")
}

func GetEditor() (string, error) {
	editors := []string{
		viper.GetString("editor"), //config
		os.Getenv("VISUAL"),       // VISUAL env variable
		os.Getenv("EDITOR"),       // EDITOR env variable
		"editor",                  // alternatives default
		"vi",                      // vi's usually there
		"nano",                    // nano's usually there if not
	}
	var err error
	var editor string = ""

	for _, cmd := range editors {
		editor, err = exec.LookPath(cmd)
		if err == nil && editor != "" {
			return editor, nil
		}
	}
	return editor, errors.New("Could not find a text editor")
}

func GetHugo() (string, error) {
	hugo, err := exec.LookPath(viper.GetString("hugo"))
	if err != nil {
		return "", errors.New("Could not find hugo executable")
	}
	return hugo, nil
}

func DisplayHumanConfig() {
	fmt.Println("Paths:")
	fmt.Printf("\tHugo root: %s\n\n", viper.GetString("root"))

	fmt.Println("Programs:")
	hugo, err := GetHugo()
	if err == nil {
		fmt.Printf("\tHugo: %s\n", hugo)
	} else {
		fmt.Printf("\t%v\n", err)
	}
	editor, err := GetEditor()
	if err == nil {
		fmt.Printf("\tEditor: %s\n", editor)
	} else {
		fmt.Printf("\t%v\n", err)
	}
	browser, err := GetBrowser()
	if err == nil {
		fmt.Printf("\tBrowser (launcher): %s\n", browser)
	} else {
		fmt.Printf("\t%v\n", err)
	}
}

func DisplayYamlConfig() {
	type RunningConfig struct {
		Root    string `yaml:"root"`
		Hugo    string `yaml:"hugo"`
		Editor  string `yaml:"editor"`
		Browser string `yaml:"browser"`
	}

	root := viper.GetString("root")
	hugo, err := GetHugo()
	if err != nil {
		hugo = fmt.Sprint(err)
	}
	editor, err := GetEditor()
	if err != nil {
		editor = fmt.Sprint(err)
	}
	browser, err := GetBrowser()
	if err != nil {
		browser = fmt.Sprint(err)
	}

	rc := RunningConfig{
		Root:    root,
		Hugo:    hugo,
		Editor:  editor,
		Browser: browser,
	}

	y, _ := yaml.Marshal(&rc)

	fmt.Print(string(y))
}
