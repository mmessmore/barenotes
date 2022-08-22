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

	"gopkg.in/yaml.v3"
	"github.com/spf13/viper"
)

type BookmarkCategory struct {
	Name string `yaml:"Name"`
	Items []Bookmark `yaml:"Items"`
}

type Bookmark struct {
	Name string `yaml:"Name"`
	Url string `yaml:"URL"`
}

func GetBookmarks() []BookmarkCategory {
	var categories []BookmarkCategory

	// TODO
	bookmarkPath := ""

	f, err := os.ReadFile(bookmarkPath)

	return categories
}

func SetBookmakrs(categories []BookmarkCategory) {
}

func RmBookmark(category string, title string) {
}

func AddBookmark(category string, title, string, url string) {
}

func GetBookmarkCategories(toComplete string) []string {
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

