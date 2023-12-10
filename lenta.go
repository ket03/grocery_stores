package main

import (
	"context"
	"fmt"
	"time"

	// "log"
	// "strconv"
	// "time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func lenta(ctx context.Context) []string {
	var nodes []*cdp.Node
	var links []string
	chromedp.Run(ctx,
		chromedp.Navigate("https://lenta.com/catalog/"),
		// chromedp.WaitVisible("body > footer"),
		chromedp.Sleep(5 * time.Second),
		chromedp.Nodes(".group-card__title", &nodes, chromedp.ByQueryAll),
	)

	fmt.Print(len(nodes))
	// for _, node := range nodes {
	// 	links = append(links, node.AttributeValue("href"))
	//   }

	return links
}

