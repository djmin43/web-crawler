package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"time"
)

func main() {

	domain := "pokemongo.inven.co.kr"
	c := colly.NewCollector(
		colly.AllowedDomains(domain),
		colly.MaxDepth(1),
	)

	c.Limit(&colly.LimitRule{
		// Filter domains affected by this rule
		DomainGlob: domain + "/*",
		// Set a delay between requests to these domains
		Delay: 1 * time.Second,
		// Add an additional random delay
		RandomDelay: 1 * time.Second,
	})

	// Find and visit all links
	c.OnHTML("span[class]", func(e *colly.HTMLElement) {
		name := e.Attr("class")
		if name == "pokemonname" {
			fmt.Println(e.Text)
		}
	})

	c.Visit("https://pokemongo.inven.co.kr/dataninfo/pokemon/")

}
