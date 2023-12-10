package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
)

// links to button more
// var URL_SPAR = "https://myspar.ru" // +catalog
var URL_SPAR = "https://myspar.ru/catalog/"
var URL_MAGNIT = "https://magnit.ru/catalog/"
var URL_LENTA = "https://lenta.com/catalog/"


func main() {
  opts := append(chromedp.DefaultExecAllocatorOptions[:],
    chromedp.Flag("headless", false),
    chromedp.Flag("start-fullscreen", true),
    chromedp.Flag("enable-automation", false), // обходит проверки
    chromedp.Flag("disable-blink-features", "AutomationControlled"), // обходит проверки
    chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"),
  )
  allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
  ctx, _ := chromedp.NewContext(allocCtx)
  // defer close()
  // fmt.Println(getLinks(ctx))
  // fmt.Print(magnit(ctx)) // один общий каталог
  // fmt.Print(spar(ctx)) // много подкаталогов
  // pyaterochka(ctx) // хз пока, ничего не работает
  fmt.Print(lenta(ctx)) // не много подкаталогов
  // auchan(ctx) // не много подкаталогов относительно спара
}

// func pyaterochka(ctx context.Context) bool {
//   var res string

//   if err := chromedp.Run(ctx,
//     browser.GrantPermissions([]browser.PermissionType{browser.PermissionTypeGeolocation}).WithOrigin("https://5ka.ru/rating/catalogue"), // выключает уведомление геолокации
//     chromedp.Navigate("https://5ka.ru/rating/catalogue"),
//     chromedp.WaitVisible(".location-confirm__button"),
//     chromedp.Click(".location-confirm__button"), // choose city
//     chromedp.Click(".focus-btn__content"), // cookies
//     // chromedp.Click(".js-next-page"), // more products
//     // chromedp.Text(".catalog-list", &res),
//   ); err != nil {
//     log.Fatal(err, "упала пятерочка")
//   }
//   fmt.Println(res)
//   return true
// }

// func auchan(ctx context.Context) bool {
//   var res string
//   if err := chromedp.Run(ctx,
//     chromedp.Navigate("https://www.auchan.ru/catalog/ovoschi-frukty-zelen-griby-yagody/"),
//     browser.GrantPermissions([]browser.PermissionType{browser.PermissionTypeGeolocation}).WithOrigin("https://www.auchan.ru/catalog/ovoschi-frukty-zelen-griby-yagody/"), // выключает уведомление геолокации
//     chromedp.Click("css-14ci9uo css-m9m6cb"),
//     chromedp.WaitVisible(".subcategoryContainer"),

//     chromedp.Text(".rr-widget-5ede0ae197a52530444d9c33-popular", &res),
//     // chromedp.Title(&res),
//   ); err != nil {
//     log.Fatal(err, "упал ашан")
//   }
//   fmt.Println(res)
//   return true
// }
