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
	"gopkg.in/yaml.v3"
)

type BookmarkCategory struct {
	Name  string     `yaml:"Name"`
	Items []Bookmark `yaml:"Items"`
}

type Bookmark struct {
	Name string `yaml:"Name"`
	Url  string `yaml:"URL"`
}

func GetBookmarks() []BookmarkCategory {
	var categories []BookmarkCategory

	d, err := os.ReadFile(filepath.Join(
		viper.GetString("root"),
		"data",
		"bookmarks.yaml",
	))
	if err != nil {
		return categories
	}

	yerr := yaml.Unmarshal(d, &categories)
	if yerr != nil {
		fmt.Println("ERROR: Could not parse bookmarks")
		fmt.Println(yerr)
	}

	return categories
}

func SetBookmarks(categories []BookmarkCategory) {
	d, _ := yaml.Marshal(&categories)

	f, err := os.Create(filepath.Join(
		viper.GetString("root"),
		"data",
		"bookmarks.yaml",
	))
	defer f.Close()

	if err != nil {
		fmt.Println("ERROR: Could not save bookmark changes")
		fmt.Println(err)
	}

	_, err = f.Write(d)
	if err != nil {
		fmt.Println("ERROR: Could not save bookmark changes")
		fmt.Println(err)
	}
}

func RmBookmark(category string, title string) bool {
	bookmarks := GetBookmarks()
	success := false

outer:
	for i := range bookmarks {
		if bookmarks[i].Name == category {
			for j := range bookmarks[i].Items {
				if bookmarks[i].Items[j].Name == title {
					bookmarks[i].Items = append(
						bookmarks[i].Items[:j],
						bookmarks[i].Items[j+1:]...,
					)
					success = true
					break outer
				}
			}
		}
	}
	if success {
		SetBookmarks(bookmarks)
	}
	return success
}

func AddBookmark(category string, title string, url string) {
	bookmarks := GetBookmarks()
	added := false

outer:
	for i := range bookmarks {
		if bookmarks[i].Name == category {
			for j := range bookmarks[i].Items {
				// bookmark exists, set the url
				if bookmarks[i].Items[j].Name == title {
					bookmarks[i].Items[j].Url = url
					// our work is done, leave the loop
					added = true
					break outer
				}
			}

			// We have the category, but no existing bookmark, so append
			bookmarks[i].Items = append(
				bookmarks[i].Items,
				Bookmark{
					Name: title,
					Url:  url,
				},
			)
			// our work is done, leave the loop
			// "outer" is not needed, just being clear
			added = true
			break outer
		}

	}

	// We never found the category, so create a new one with 1 bookmark
	if !added {
		bookmarks = append(
			bookmarks,
			BookmarkCategory{
				Name: category,
				Items: []Bookmark{{
					Name: title,
					Url:  url,
				}},
			},
		)
	}

	SetBookmarks(bookmarks)
}

func GetBookmarkCategories(toComplete string) []string {

	var categories []string
	bookmarks := GetBookmarks()

	for _, category := range bookmarks {
		if strings.Contains(
			strings.ToLower(category.Name),
			strings.ToLower(toComplete),
		) {
			categories = append(
				categories,
				category.Name,
			)
		}
	}
	return categories
}
