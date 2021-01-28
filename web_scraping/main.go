package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("ikman.lk", "https://ikman.lk/en/ad/suzuki-wagon-r-stingray-2015-for-sale-colombo-1416"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML(".details-section--2ggRy", func(e *colly.HTMLElement) {
		link := e.Attr("class")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL.String())
	// })

	// Start scraping on https://hackerspaces.org
	c.Visit("https://ikman.lk/en/ad/suzuki-wagon-r-stingray-2015-for-sale-colombo-1416")

}
