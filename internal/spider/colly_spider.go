package main

import (
	"fmt"
	colly "github.com/gocolly/colly"
	"perch/internal/spider/anaysis/website_anaysis"
	"time"
)

func main() {

	/*CollySpider := &crawl.CrawlColly{}
	CollySpider.CollyInit()
	CollySpider.Collystart()*/

	targetURL := "https://www.amazon.com/DualShock-Wireless-Controller-PlayStation-Magma-4/dp/B01MD19OI2?pd_rd_w=HnCmf&pf_rd_p=842729f9-b53c-45b9-a4ab-635ab0029964&pf_rd_r=4DRFB0V12HD4WCS8TBNM&pd_rd_r=c4d4ff36-5474-4406-8133-511fabbeb4bb&pd_rd_wg=c3GZF&th=1"
	collector := colly.NewCollector(
		//colly.Debugger(&debug.LogDebugger{}),
		colly.Async(true),
	)
	collector.Limit(
		&colly.LimitRule{
			DomainRegexp: "",
			DomainGlob:   "",
			Delay:        5 * time.Second,
			RandomDelay:  5 * time.Second,
			Parallelism:  2,
		})

	// Find and visit all links
	collector.OnHTML("a", website_anaysis.ProductInformation)
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	collector.Visit(targetURL)
	collector.Wait()
}
