package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func magnit(ctx context.Context) map[string]float64 {
	var nodes_price []*cdp.Node
	var nodes_name []*cdp.Node
	var price []float64
	var name []string
	product := make(map[string]float64)

	chromedp.Run(ctx,
	  chromedp.Navigate("https://magnit.ru/catalog/"),
	  chromedp.Sleep(time.Second * 4),
	  chromedp.Click(".alcohol__success"),
	  chromedp.Click(".city-leaving-cancel"),
	  chromedp.Click(".cookies__button"),
	  chromedp.Sleep(5 * time.Second),
	)
  
	// magnitHasButton(ctx)
  
	chromedp.Run(ctx,
	  chromedp.Nodes(".new-card-product__price-regular", &nodes_price, chromedp.ByQueryAll),
	  chromedp.Nodes(".new-card-product__title", &nodes_name, chromedp.ByQueryAll),
	)
  
	for _, node := range nodes_price {
	  price = append(price, magnitConvertStringToFloat(node.Children[0].NodeValue))
	}
  
	for _, node := range nodes_name {
		name = append(name, node.Children[0].NodeValue)
	}

	for i := range name {
		product[name[i]] = price[i]
	}

	return product
  }
  
  func magnitConvertStringToFloat(price string) float64 {
	price = price[:len(price)-4]
	price = price[:len(price)-3] + "." + price[len(price)-2:]
	res, _ := strconv.ParseFloat(price, 64)
	return res
  }
  
  func magnitHasButton(ctx context.Context) {
	var button []*cdp.Node
	chromedp.Run(ctx,
	// chromedp.WaitVisible("body > footer"),
	chromedp.Sleep(5 * time.Second),
	chromedp.ActionFunc(func (ctx context.Context) error {
	  if err := chromedp.Nodes(".paginate__more", &button, chromedp.AtLeast(0)).Do(ctx); err != nil {log.Fatal(err)}
	  if len(button) > 0 {
		chromedp.Click(".paginate__more").Do(ctx)
		magnitHasButton(ctx)
	  }
	  return nil
	}),)
  }