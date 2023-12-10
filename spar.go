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

func getLinks(ctx context.Context) []string {
	var nodes []*cdp.Node
	var links []string
	chromedp.Run(ctx,
	  chromedp.Navigate("https://myspar.ru/catalog/"),
	  chromedp.WaitVisible("body > footer"),
	  chromedp.Nodes(`a.section-tile__head`, &nodes, chromedp.ByQueryAll, chromedp.NodeVisible),
	)
  
	for _, node := range nodes {
	  links = append(links, node.AttributeValue("href"))
	}

	return links
  }
  
  func spar(ctx context.Context) map[string]float64 {
	product := make(map[string]float64)
	var nodes []*cdp.Node
	var price_str string
	var name string
  
	chromedp.Run(ctx,
	  chromedp.Navigate("https://myspar.ru/catalog/pitstsa-4/"),
	)
  
	sparHasButton(ctx)
  
	chromedp.Run(ctx, 
	  chromedp.Nodes(".catalog-tile__area", &nodes, chromedp.ByQueryAll),
	)
  
	for _, node := range nodes {
	  chromedp.Run(ctx,
		chromedp.Text(".prices__cur", &price_str, chromedp.ByQuery, chromedp.FromNode(node)),
		chromedp.Text(".catalog-tile__name", &name, chromedp.ByQuery, chromedp.FromNode(node)),
	  )
	  price := sparConvertStringToFloat(price_str)
	  product[name] = price
	}
  
	return product
  }
  
  func sparHasButton(ctx context.Context) {
	var button []*cdp.Node
	chromedp.Run(ctx,
	chromedp.WaitVisible("body > footer"),
	chromedp.Sleep(4 * time.Second),
	chromedp.ActionFunc(func (ctx context.Context) error {
	  if err := chromedp.Nodes(".js-next-page", &button, chromedp.AtLeast(0)).Do(ctx); err != nil {log.Fatal(err)}
	  if len(button) > 0 {
		chromedp.Click(".js-next-page").Do(ctx)
		sparHasButton(ctx)
	  }
	  return nil
	}),)
  }
  
  func sparConvertStringToFloat(price string) float64 {
	price = price[:len(price)-6]
	price = price[:len(price)-1] + "." + price[len(price)-1:]
	res, _ := strconv.ParseFloat(price, 64)
	return res
  }

  func chooseCity() string {
	var city string
	fmt.Scan(&city)
	return city
  }