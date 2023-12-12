package main

import (
	"context"
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
	)

	buttonMore(ctx, MAGNIT_PATH_BUTTON_MORE)

	chromedp.Run(ctx,
	  chromedp.Nodes(".new-card-product__price-regular", &nodes_price, chromedp.ByQueryAll),
	  chromedp.Nodes(".new-card-product__title", &nodes_name, chromedp.ByQueryAll),
	)
  
	for _, node := range nodes_price {
	  price = append(price, convertStringToFloat(node.Children[0].NodeValue, MAGNIT_CUT, MAGNIT_START, MAGNIT_END))
	}
  
	for _, node := range nodes_name {
		name = append(name, node.Children[0].NodeValue)
	}

	for i := range name {
		product[name[i]] = price[i]
	}

	return product
  }

  // get data
