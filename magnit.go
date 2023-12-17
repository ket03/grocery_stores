package main

import (
	"context"
	"maps"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func magnit(ctx context.Context) map[string]int {
	var last_page int
	product := make(map[string]int)

	chromedp.Run(ctx,
		chromedp.Navigate(URL_MAGNIT),
		chromedp.Sleep(time.Second*4),
		chromedp.Click(".alcohol__success"),
		chromedp.Click(".city-leaving-cancel"),
		chromedp.Click(".cookies__button"),
		chromedp.Sleep(3*time.Second),
	)

	last_page = getCounterPages(ctx)
	clickButtonMore(ctx, last_page)
	maps.Copy(product, getDataMagnit(ctx))

	return product
}

func getCounterPages(ctx context.Context) int {
	var nodes_page []*cdp.Node
	var last_page_str string
	var last_page int
	chromedp.Run(ctx,
		chromedp.Nodes(".num", &nodes_page),
	)
	for _, node := range nodes_page {
		last_page_str = node.Children[0].NodeValue
	}
	last_page, _ = strconv.Atoi(last_page_str)
	return last_page
}

func clickButtonMore(ctx context.Context, last_page int) {
	for i := 0; i < last_page-1; i++ {
		chromedp.Run(ctx,
			chromedp.Sleep(3*time.Second),
			chromedp.Click(MAGNIT_PATH_BUTTON_MORE),
		)
	}
}

func getDataMagnit(ctx context.Context) map[string]int {
	var nodes_price []*cdp.Node
	var nodes_name []*cdp.Node
	var price []int
	var name []string
	product := make(map[string]int)

	chromedp.Run(ctx,
		chromedp.Nodes(".new-card-product__price-regular", &nodes_price),
		chromedp.Nodes(".new-card-product__title", &nodes_name),
	)
	for _, node := range nodes_price {
		price = append(price, convertStringToInt(node.Children[0].NodeValue, MAGNIT_CUT))
	}
	for _, node := range nodes_name {
		name = append(name, node.Children[0].NodeValue)
	}
	for i := range price {
		product[name[i]] = price[i]
	}
	return product
}
