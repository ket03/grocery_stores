package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func getLinks(ctx context.Context) []string {
	var nodes []*cdp.Node
	var links []string
	chromedp.Run(ctx,
		chromedp.WaitVisible("body > footer"),
		chromedp.Nodes(`a.section-tile__head`, &nodes, chromedp.NodeVisible),
	)

	for _, node := range nodes {
		links = append(links, node.AttributeValue("href"))
	}

	return links
}

func spar(ctx context.Context) map[string]float64 {
	product := make(map[string]float64)
	chromedp.Run(ctx,
		chromedp.Navigate("https://myspar.ru/catalog/"))
	endPoint := getLinks(ctx)

	for i := range endPoint {
		chromedp.Run(ctx,
			chromedp.Navigate(URL_SPAR+endPoint[i]),
			chromedp.Sleep(5*time.Second),
		)
		product = getData(ctx)
	}
	return product
}

func getData(ctx context.Context) map[string]float64 {
	product := make(map[string]float64)
	var nodes []*cdp.Node
	var links []string
	isSubcatalog := false

	chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			if err := chromedp.Nodes(".catalog-tile__area", &nodes, chromedp.AtLeast(0)).Do(ctx); err != nil {
				log.Fatal(err)
			}
			if len(nodes) == 0 {
				isSubcatalog = true
			}
			return nil
		}),
	)

	if !isSubcatalog {
		buttonMore(ctx, SPAR_PATH_BUTTON_MORE)
		product = getDataHelper(ctx)

	} else if isSubcatalog {
		links = getLinks(ctx)
		for i, link := range links {
			fmt.Println("link =", i, link)
			chromedp.Run(ctx,
				chromedp.Navigate(URL_SPAR+links[i]),
				chromedp.Sleep(5*time.Second),
			)
			buttonMore(ctx, SPAR_PATH_BUTTON_MORE)
			product = getDataHelper(ctx)
		}
	}
	fmt.Println(product, "\n\n\n\n\n\n")
	return product
}

func getDataHelper(ctx context.Context) map[string]float64 {
	product := make(map[string]float64)
	var price_str string
	var name string
	var nodes []*cdp.Node
	var nodes_price []*cdp.Node
	var nodes_name []*cdp.Node
	chromedp.Run(ctx,
		chromedp.Nodes(".catalog-tile__area", &nodes),
		chromedp.Nodes(".prices__cur", &nodes_price),
		chromedp.Nodes(".catalog-tile__name", &nodes_name),
	)
	if len(nodes_name) > len(nodes_price) {
		return product
	}
	for _, node := range nodes {
		chromedp.Run(ctx,
			chromedp.Text(".prices__cur", &price_str, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text(".catalog-tile__name", &name, chromedp.ByQuery, chromedp.FromNode(node)),
		)
		fmt.Println(name)

		price := convertStringToFloat(price_str, SPAR_CUT, SPAR_START, SPAR_END)
		product[name] = price
	}
	fmt.Println(product)
	return product
}

// func test(ctx context.Context) {
// 	chromedp.Run(ctx,
// 		chromedp.Navigate("https://myspar.ru/catalog/gorodetskaya-rospis-1/"),
// 	)
// 	buttonMore(ctx, SPAR_PATH_BUTTON_MORE)
// 	fmt.Println(getDataHelper(ctx))
// }

// non-price-oblect
// button time-alert
