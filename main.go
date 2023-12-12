package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	// "fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var URL_SPAR = "https://myspar.ru" // +catalog
var URL_SPAR_CATALOG = "https://myspar.ru/catalog/"
var SPAR_PATH_BUTTON_MORE = ".js-next-page"
var SPAR_PATH_BUTTON_ALERT = ".js-modal-close"
var SPAR_CUT = 6
var SPAR_START = 1
var SPAR_END = 1

var URL_MAGNIT = "https://magnit.ru/catalog/"
var MAGNIT_PATH_BUTTON_MORE = ".paginate__more"
var MAGNIT_CUT = 4
var MAGNIT_START = 3
var MAGNIT_END = 2

var URL_LENTA = "https://lenta.com/catalog/"
var LENTA_PATH_BUTTON_MORE = ".button--primary catalog-grid__pagination-button"
var LENTA_CUT = 0
var LENTA_START = 0
var LENTA_END = 0

var URL_AUCHAN = "https://www.auchan.ru/catalog"
var AUCHAN_PATH_BUTTON_MORE = ".showMoreButton"
var AUCHAN_CUT = 0
var AUCHAN_START = 0
var AUCHAN_END = 0

var URL_PYATEEROCHKA = "https://5ka.ru/rating/catalogue"
var PYATEEROCHKA_PATH_BUTTON_MORE = ""
var PYATEEROCHKA_CUT = 0
var PYATEEROCHKA_START = 0
var PYATEEROCHKA_END = 0

func main() {
  opts := append(chromedp.DefaultExecAllocatorOptions[:],
    chromedp.Flag("headless", false),
    chromedp.Flag("start-fullscreen", true),
    chromedp.Flag("enable-automation", false), // обходит проверки
    chromedp.Flag("disable-blink-features", "AutomationControlled"), // обходит проверки
    chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"),
  )
  allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
  ctx, close := chromedp.NewContext(allocCtx)
  defer close()
  fmt.Println(spar(ctx))
  // fmt.Println(magnit(ctx))
  // fmt.Println(lenta(ctx))
  // fmt.Println(auchan(ctx))
  // fmt.Println(pyaterochka(ctx))
}

func buttonMore(ctx context.Context, button_path string) {
	var button []*cdp.Node
	chromedp.Run(ctx,
		chromedp.WaitVisible("body > footer"),
		chromedp.Sleep(4*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			if err := chromedp.Nodes(button_path, &button, chromedp.AtLeast(0)).Do(ctx); err != nil {
				log.Fatal(err)
			}
			if len(button) > 0 {
				chromedp.Click(button_path).Do(ctx)
				buttonMore(ctx, button_path)
			}
			return nil
		}))
}

func convertStringToFloat(price string, cut int, start int, end int) float64 {
  price = price[:len(price)-cut]
	price = price[:len(price)-start] + "." + price[len(price)-end:]
	res, _ := strconv.ParseFloat(price, 64)
	return res
}
