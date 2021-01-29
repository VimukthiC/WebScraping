package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	
	c := colly.NewCollector(
		colly.AllowedDomains("ikman.lk"),
	)
	
	c.OnHTML(".details-section--2ggRy", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	
	c.Visit("https://ikman.lk/en/ad/suzuki-wagon-r-stingray-2015-for-sale-colombo-1416")
}
