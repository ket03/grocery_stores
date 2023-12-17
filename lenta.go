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

func lenta(ctx context.Context) map[string]int {
	var nodes []*cdp.Node
	product := make(map[string]int)
	// var links []string
	chromedp.Run(ctx,
		chromedp.Navigate("https://lenta.com/catalog/"),
		chromedp.Sleep(3 * time.Second),
		chromedp.Nodes(".group-card__image", &nodes),
	)

	fmt.Println(len(nodes))
	return product
}

