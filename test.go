package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/cdproto/browser"
	"github.com/chromedp/chromedp"
)

var URL_MAGNIT = "https://magnit.ru/catalog/"

var URL_SPAR_GREEN           = "https://myspar.ru/catalog/zelen-salaty/"
var URL_SPAR_VEGETABLES        = "https://myspar.ru/catalog/ovoshchi/"
var URL_SPAR_FRUIT           = "https://myspar.ru/catalog/frukty/"

var URL_PYATEROCHKA_CATALOG       = "https://5ka.ru/rating/catalogue"

var URL_LENTA_FRUIT_AND_VEGETABLES = "https://lenta.com/catalog/frukty-i-ovoshchi/"
var URL_LENTA_MEAT            = "https://lenta.com/catalog/myaso-ptica-kolbasa/"
var URL_LENTA_FISH           = "https://lenta.com/catalog/ryba-i-moreprodukty/"


func main() {
  opts := append(chromedp.DefaultExecAllocatorOptions[:],
    chromedp.Flag("headless", false),
    chromedp.Flag("enable-automation", false), // обходит проверки
    chromedp.Flag("disable-blink-features", "AutomationControlled"), // обходит проверки
    chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"),
  )
  
  allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
  ctx, close := chromedp.NewContext(allocCtx)
  defer close()
  // magnit(ctx) // один общий каталог
  // spar(ctx) // много подкаталогов
  // pyaterochka(ctx) // хз пока, ничего не работает
  // lenta(ctx) // не много подкаталогов
  // auchan(ctx) // не много подкаталогов относительно спара
}

func magnit(ctx context.Context) bool {
  var res string
  if err := chromedp.Run(ctx,
    chromedp.Navigate("https://magnit.ru/catalog/?categoryId=4893"),
    chromedp.Click(".city-leaving-cancel"), // choose city
    chromedp.Click(".cookies__button"), // cookies
    chromedp.Click(".paginate__more"), // more products
    chromedp.Text(".catalog-page__product-grid", &res),
  ); err != nil {
    log.Fatal(err, "упал магнит")
  }
  fmt.Println(res)
  return true
}

func spar(ctx context.Context) bool {
  var res string
  if err := chromedp.Run(ctx,
    chromedp.Navigate("https://myspar.ru/catalog/torty-i-pirozhnye-1/"),
    chromedp.Click(".js-city-confirm-close"), // choose city
    chromedp.Click(".js-popup-note-close"), // cookies
    chromedp.Click(".js-next-page"), // more products
    chromedp.Text(".catalog-list", &res),
  ); err != nil {
    log.Fatal(err, "упал спар")
  }
  fmt.Println(res)
  return true
}

func pyaterochka(ctx context.Context) bool {
  var res string

  if err := chromedp.Run(ctx,
    browser.GrantPermissions([]browser.PermissionType{browser.PermissionTypeGeolocation}).WithOrigin("https://5ka.ru/rating/catalogue"), // выключает уведомление геолокации
    chromedp.Navigate("https://5ka.ru/rating/catalogue"),
    chromedp.WaitVisible(".location-confirm__button"),
    chromedp.Click(".location-confirm__button"), // choose city
    chromedp.Click(".focus-btn__content"), // cookies
    // chromedp.Click(".js-next-page"), // more products
    // chromedp.Text(".catalog-list", &res),
  ); err != nil {
    log.Fatal(err, "упала пятерочка")
  }
  fmt.Println(res)
  return true
}

// данные о продуктах и кнопки куки / гео находятся на одном уровне
func lenta(ctx context.Context) bool {
  var res string
  if err := chromedp.Run(ctx,
    chromedp.Navigate("https://lenta.com/catalog/frukty-i-ovoshchi/"),

    chromedp.Text(".catalog-grid__grid", &res),
  ); err != nil {
    log.Fatal(err, "упала лента")
  }
  fmt.Println(res)
  return true
}

func auchan(ctx context.Context) bool {
  var res string
  if err := chromedp.Run(ctx,
    chromedp.Navigate("https://www.auchan.ru/catalog/ovoschi-frukty-zelen-griby-yagody/"),
    browser.GrantPermissions([]browser.PermissionType{browser.PermissionTypeGeolocation}).WithOrigin("https://www.auchan.ru/catalog/ovoschi-frukty-zelen-griby-yagody/"), // выключает уведомление геолокации
    chromedp.Click("css-14ci9uo css-m9m6cb"),
    chromedp.WaitVisible(".subcategoryContainer"),

    chromedp.Text(".rr-widget-5ede0ae197a52530444d9c33-popular", &res),
    // chromedp.Title(&res),
  ); err != nil {
    log.Fatal(err, "упал ашан")
  }
  fmt.Println(res)
  return true
}