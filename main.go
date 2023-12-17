package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var URL_SPAR = "https://myspar.ru" // +catalog
var URL_SPAR_CATALOG = "https://myspar.ru/catalog/"
var SPAR_ENDPOINT = []string{"/catalog/khleb-torty-sladosti-1/",
	"/catalog/torty-pirozhnye/", "/catalog/ovoshchi-frukty-yagody-griby-1/", "/catalog/moloko-syr-yaytsa-1/",
	"/catalog/syry-1/", "/catalog/kolbasy-sosiski-myasnye-delikatesy/", "/catalog/myaso-ptitsa-kolbasy-1/",
	"/catalog/ryba-1/", "/catalog/ikra-novyy-ulov-2023/", "/catalog/ryba-moreprodukty-ikra/", "/catalog/makarony-krupy-spetsii-maslo-1/",
	"/catalog/sousy-ketchupy-mayonezy/", "/catalog/konservy-orekhi-sneki-varene/", "/catalog/okhlazhdennye-i-zamorozhennye-polufabrikaty/",
	"/catalog/voda-soki-napitki-1/", "/catalog/chay-kofe-sakhar/", "/catalog/konfety-shokolad-sladosti/", "/catalog/sneki-1/",
	"/catalog/nashi-marki/", "/catalog/zdorovoe-pitanie/", "/catalog/sportivnoe-pitanie-1/", "/catalog/krasota-gigiena-zdorove-1/",
	"/catalog/tovary-dlya-uborki/", "/catalog/tovary-dlya-doma/", "/catalog/avto-otdykh-dacha/"}
var SPAR_PATH_BUTTON_MORE = ".js-next-page"
var SPAR_PATH_BUTTON_ALERT = ".js-modal-close"
var SPAR_CUT = 7

var URL_MAGNIT = "https://magnit.ru/catalog/"
var MAGNIT_PATH_BUTTON_MORE = ".paginate__more"
var MAGNIT_CUT = 7

var URL_LENTA = "https://lenta.com/catalog/"
var LENTA_PATH_BUTTON_MORE = ".button--primary catalog-grid__pagination-button"
var LENTA_CUT = 0

var URL_AUCHAN = "https://www.auchan.ru/catalog"
var AUCHAN_PATH_BUTTON_MORE = ".showMoreButton"
var AUCHAN_CUT = 0

var URL_PYATEEROCHKA = "https://5ka.ru/rating/catalogue"
var PYATEEROCHKA_PATH_BUTTON_MORE = ""
var PYATEEROCHKA_CUT = 0

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
//   fmt.Println(spar(ctx))
  fmt.Println(magnit(ctx))
  // fmt.Println(lenta(ctx))
  // fmt.Println(auchan(ctx))
  // fmt.Println(pyaterochka(ctx))
}

func buttonMore(ctx context.Context, button_path string) {
	var button []*cdp.Node
	if err := chromedp.Run(ctx,
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
		})); err != nil {
			log.Fatalln("crash when clicked button more")
		}
}

func convertStringToInt(price string, cut int) int {
  price = price[:len(price)-cut]
	res, _ := strconv.Atoi(price)
	return res
}
