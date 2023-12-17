package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func auchan(ctx context.Context) map[string]int{
	product := make(map[string]int)
	var links []string
	chromedp.Run(ctx,
		chromedp.Navigate("https://www.auchan.ru/catalog/"),
		chromedp.Sleep(3 * time.Second),

	)
	links = getLinksAuchan(ctx)
	for i := range links {
		chromedp.Navigate("https://www.auchan.ru" + links[i])
		
	}
	return product
}

func getLinksAuchan(ctx context.Context) []string{
	var links []string
	var nodes []*cdp.Node
	chromedp.Run(ctx,
		chromedp.Nodes(".youMayNeedCategoryItem", &nodes),
	)
	for _, node := range nodes {
		links = append(links, node.AttributeValue("href"))
	}
	return links
}

func getDataAuchan(ctx context.Context) map[string]int {
	products = make(map[string]int)
	chromedp.
	
	return products
}