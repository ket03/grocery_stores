package main

import (
  "context"
  "fmt"
  "log"

  "github.com/chromedp/chromedp"
)

// можно добавить еще много вкладок
var URL_SPAR_GREEN           = "https://myspar.ru/catalog/zelen-salaty/"
var URL_SPAR_VEGETABLES        = "https://myspar.ru/catalog/ovoshchi/"
var URL_SPAR_FRUIT           = "https://myspar.ru/catalog/frukty/"

// можно добавить отдельно по айдишкам
var URL_MAGNIT_CATALOG          = "https://magnit.ru/catalog/"

// можно добавить отдельно по айдишкам
var URL_PYATEROCHKA_CATALOG       = "https://5ka.ru/rating/catalogue"

// можно добавить еще пару вкладок
var URL_LENTA_FRUIT_AND_VEGETABLES = "https://lenta.com/catalog/frukty-i-ovoshchi/"
var URL_LENTA_MEAT            = "https://lenta.com/catalog/myaso-ptica-kolbasa/"
var URL_LENTA_FISH           = "https://lenta.com/catalog/ryba-i-moreprodukty/"


func main() {
  opts := append(chromedp.DefaultExecAllocatorOptions[:],
    chromedp.Flag("headless", false),
    chromedp.Flag("enable-automation", false), // обходит проверки
    chromedp.Flag("disable-blink-features", "AutomationControlled"), // обходит проверки
  )
  
  allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
  ctx, _ := chromedp.NewContext(allocCtx)
  // defer close()
  // magnit(ctx)
  // spar(ctx)
  // pyaterochka(ctx)
  lenta(ctx)
  // auchan(ctx)
}

func magnit(ctx context.Context) bool {
  var res string
  if err := chromedp.Run(ctx,
    chromedp.Navigate(https://magnit.ru/catalog/?categoryId=4893),
    chromedp.Click(".city-leaving-cancel"), // choose city
    chromedp.Click(".cookies__button"), // cookies
    chromedp.Click(".paginate__more"), // more products
    chromedp.Text(".catalog-page__product-grid", &res),
  ); err != nil {
    log.Fatal(err)
  }
  fmt.Println(res)
  return true
}

func spar(ctx context.Context) bool {
  var res string
  if err := chromedp.Run(ctx,
    chromedp.Navigate(https://myspar.ru/catalog/torty-i-pirozhnye-1/),
    chromedp.Click(".js-city-confirm-close"), // choose city
    chromedp.Click(".js-popup-note-close"), // cookies
    chromedp.Click(".js-next-page"), // more products
    chromedp.Text(".catalog-list", &res),
  ); err != nil {
    log.Fatal(err)
  }
  fmt.Println(res)
  return true
}

func pyaterochka(ctx context.Context) bool {
  var res string
  if err := chromedp.Run(ctx,
    chromedp.Navigate(https://5ka.ru/rating/catalogue),
    chromedp.WaitVisible(".location-confirm__button"),
    chromedp.Click(".location-confirm__button"), // choose city
    chromedp.Click(".focus-btn__content"), // cookies
    // chromedp.Click(".js-next-page"), // more products
    // chromedp.Text(".catalog-list", &res),
  ); err != nil {
    log.Fatal(err)
  }
  fmt.Println(res)
  return true
}

func lenta(ctx context.Context) bool {
  var res string
  if err := chromedp.Run(ctx,
    chromedp.Navigate(https://lenta.com/catalog/frukty-i-ovoshchi/),
    // chromedp.Click(".js-city-confirm-close"), // choose city
    // chromedp.Click(".js-popup-note-close"), // cookies
    // chromedp.Click(".js-next-page"), // more products
    // chromedp.Text(".catalog-list", &res),
  ); err != nil {
    log.Fatal(err)
  }
  fmt.Println(res)
  return true
}

func auchan(ctx context.Context) bool {
  var res string
  if err := chromedp.Run(ctx,
    chromedp.Navigate(https://myspar.ru/catalog/torty-i-pirozhnye-1/),
    // chromedp.Click(".js-city-confirm-close"), // choose city
    // chromedp.Click(".js-popup-note-close"), // cookies
    // chromedp.Click(".js-next-page"), // more products
    // chromedp.Text(".catalog-list", &res),
  ); err != nil {
    log.Fatal(err)
  }
  fmt.Println(res)
  return true
}