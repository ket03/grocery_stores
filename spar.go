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
		chromedp.Navigate("https://myspar.ru/catalog/"),)
	endPoint := getLinks(ctx)

	for i := range endPoint {
		fmt.Println("catalog =", i)
		chromedp.Run(ctx,
			chromedp.Navigate(URL_SPAR + endPoint[i]),
			chromedp.Sleep(5 * time.Second),
		)
		fmt.Println("exit catalog")
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
		chromedp.ActionFunc(func (ctx context.Context) error {
			if err := chromedp.Nodes(".catalog-tile__area", &nodes, chromedp.AtLeast(0)).Do(ctx); err != nil {log.Fatal(err)}
			if len(nodes) == 0 {
				isSubcatalog = true
			}
			return nil
		}),
	  )
	    
	if !isSubcatalog {
		sparHasButton(ctx)
		product = getDataHelper(ctx)

	} else if isSubcatalog {
		links = getLinks(ctx)
		for i, link := range links {
			fmt.Println("link =", i, link)
			chromedp.Run(ctx,
				chromedp.Navigate(URL_SPAR + links[i]),
				chromedp.Sleep(5 * time.Second),
			)
			sparHasButton(ctx)
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
	var reserve []*cdp.Node
	hasPrice := true
	fmt.Println("before search node")
	chromedp.Run(ctx,
		chromedp.Nodes(".catalog-tile__area", &nodes),
	)
	fmt.Println("after search node")
		for _, node := range nodes {
			chromedp.Run(ctx,
				chromedp.ActionFunc(func (ctx context.Context) error {
					if err := chromedp.Text(".prices__cur", &price_str, chromedp.FromNode(node)).Do(ctx); err != nil {log.Fatal(err)}
					if len(price_str) == 0 {
						hasPrice = false
					}
					fmt.Println(len(reserve))
					return nil
				}),
			  )
			  fmt.Println("after search price")
			  if hasPrice {
				chromedp.Run(ctx,
					chromedp.Text(".prices__cur", &price_str, chromedp.ByQuery, chromedp.FromNode(node)),
					chromedp.Text(".catalog-tile__name", &name, chromedp.ByQuery, chromedp.FromNode(node)),
				  )
			  } else if !hasPrice {
				price_str = "0"
				chromedp.Run(ctx,
					chromedp.Text(".catalog-tile__name", &name, chromedp.ByQuery, chromedp.FromNode(node)),
				)
			  }

			price := sparConvertStringToFloat(price_str)
			product[name] = price
		}
		fmt.Println("exit ")
		fmt.Println(product)
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
	res, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println("цена указана не в корректном формате")
		return 0
	}
	return res
  }
