package main

import (
	"context"
	"fmt"
	"maps"
	"time"

	// "log"
	// "strconv"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func lenta(ctx context.Context) map[string]int {
	product := make(map[string]int)
	var links []string

	chromedp.Run(ctx,
		chromedp.Navigate("https://lenta.com/catalog/"),
		chromedp.Sleep(3*time.Second),
	)

	links = getLinksLenta(ctx)

	for i := range links {
		chromedp.Run(ctx,
			chromedp.Navigate(links[i]),
			chromedp.Sleep(3*time.Second),
		)
		buttonMore(ctx, LENTA_PATH_BUTTON_MORE)
		maps.Copy(product, getDataLenta(ctx))
	}

	return product
}

func getLinksLenta(ctx context.Context) []string {
	var nodes []*cdp.Node
	var links []string
	chromedp.Run(ctx,
		chromedp.Nodes("a.group-card", &nodes),
	)
	for _, node := range nodes {
		links = append(links, node.AttributeValue("href"))
	}
	return links
}

func getDataLenta(ctx context.Context) map[string]int {
	product := make(map[string]int)
	var nodes []*cdp.Node
	var name string
	var price_str string

	chromedp.Run(ctx,
		chromedp.Nodes(".catalog-grid-sku-product-card", &nodes),
	)

	for _, node := range nodes {
		chromedp.Run(ctx,
			chromedp.Text(".lui-sku-product-card-text", &name, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(".lui-priceText", &price_str, chromedp.ByQuery, chromedp.FromNode(node)),
		)
		product[name] = convertStringToInt(price_str, MAGNIT_CUT)
	}
	fmt.Println(product)
	fmt.Println(price_str)
	return product
}
