package internal

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/wangii/emoji"
)

/*
	Converts Hugo markdown to PDF via Pandoc
	This handles converting emoji characters or "tags" to image links
	so pdflatex doesn't lose its mind and we don't have to do crazy hacks.
*/
func ConvertToPDF(inPath string, outPath string) {
	input, err := ioutil.ReadFile(inPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inText := string(input)

	tFile, err := ioutil.TempFile("/tmp", "messynotes")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer os.Remove(tFile.Name())

	outText := emoji.UnicodeToTwemoji(inText, 12, false)
	outText = emoji.EmojiTagToTwemoji(outText, 12, false)

	outFile, err := os.Create(outPath)
	defer outFile.Close()

	outFile.WriteString(outText)
}
