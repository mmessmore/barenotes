package internal

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

/*
Trying headless chrome
*/
func ConvertToPDF(title string, outPath string) {

	_, _, exists := NotePathsByTitle(title)

	if !exists {
		fmt.Printf("Note does not exist: %s\n", title)
		os.Exit(2)
	}

	url := fmt.Sprintf("http://localhost:1313/notes/%s/", title)
	fmt.Printf("DEBUG: %s\n", url)
	var pdfBuffer []byte

	taskCtx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	err := chromedp.Run(taskCtx, chromePDF(url, &pdfBuffer))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(outPath, pdfBuffer, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func chromePDF(url string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		emulation.SetUserAgentOverride("Messynotes 1.0"),
		chromedp.Navigate(url),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(true).Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
