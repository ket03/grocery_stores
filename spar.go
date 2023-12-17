package main

import (
	"context"
	"log"
	"maps"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func spar(ctx context.Context) map[string]int {
	product := make(map[string]int)

	for i := range SPAR_ENDPOINT {
		if err := chromedp.Run(ctx,
			chromedp.Navigate(URL_SPAR+SPAR_ENDPOINT[i]),
			chromedp.Sleep(5*time.Second),
		); err != nil {
			log.Fatalln("crash when navigating to endpoint(outside)")
		}
		maps.Copy(product, getData(ctx))
		log.Println(i+1, " catalog completed")
	}
	return product
}

func getData(ctx context.Context) map[string]int {
	product := make(map[string]int)
	var nodes []*cdp.Node
	var links []string
	isSubcatalog := false

	if err := chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			if err := chromedp.Nodes(".catalog-tile__area", &nodes, chromedp.AtLeast(0)).Do(ctx); err != nil {
				log.Fatal(err)
			}
			if len(nodes) == 0 {
				isSubcatalog = true
			}
			return nil
		}),
	); err != nil {
		log.Fatalln("crash when test-search product")
	}

	if !isSubcatalog {
		buttonMore(ctx, SPAR_PATH_BUTTON_MORE)
		maps.Copy(product, getDataHelper(ctx))

	} else if isSubcatalog {
		links = getLinks(ctx)
		for i, link := range links {
			log.Println("link =", i, link)
			if err := chromedp.Run(ctx,
				chromedp.Navigate(URL_SPAR+links[i]),
				chromedp.Sleep(5*time.Second),
			); err != nil {
				log.Fatalln("crash when navigating to endpoint(inside)")
			}
			buttonMore(ctx, SPAR_PATH_BUTTON_MORE)
			maps.Copy(product, getDataHelper(ctx))
			log.Println(i+1, " subcatalog completed")
		}
	}
	return product
}

func getDataHelper(ctx context.Context) map[string]int {
	product := make(map[string]int)
	var price_str string
	var name string
	var nodes []*cdp.Node
	var nodes_price []*cdp.Node
	var nodes_name []*cdp.Node
	if err := chromedp.Run(ctx,
		chromedp.Nodes(".catalog-tile__area", &nodes),
		chromedp.Nodes(".prices__cur", &nodes_price),
		chromedp.Nodes(".catalog-tile__name", &nodes_name),
	); err != nil {
		log.Fatalln("crash when count products")
	}
	if len(nodes_name) > len(nodes_price) {
		return product
	}
	for _, node := range nodes {
		if err := chromedp.Run(ctx,
			chromedp.Text(".prices__cur", &price_str, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(".catalog-tile__name", &name, chromedp.ByQuery, chromedp.FromNode(node)),
		); err != nil {
			log.Fatalln("crash when get product")
		}
		product[name] = convertStringToInt(price_str, SPAR_CUT)
	}
	return product
}

func getLinks(ctx context.Context) []string {
	var nodes []*cdp.Node
	var links []string
	if err := chromedp.Run(ctx,
		chromedp.WaitVisible("body > footer"),
		chromedp.Nodes(`a.section-tile__head`, &nodes, chromedp.NodeVisible),
	); err != nil {
		log.Fatalln("crash when parsing links")
	}

	for _, node := range nodes {
		links = append(links, node.AttributeValue("href"))
	}

	return links
}

// non-price-oblect
// button time-alert
