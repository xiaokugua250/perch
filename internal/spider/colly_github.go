package main

import (
	"fmt"
	colly "github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"perch/internal/spider/anaysis/website_anaysis/github"
	"time"
)

func main() {

	/*CollySpider := &crawl.CrawlColly{}
	CollySpider.CollyInit()
	CollySpider.Collystart()*/

	targetURL := "https://github.com/avelino/awesome-go"
	collector := colly.NewCollector(
		//colly.Debugger(&debug.LogDebugger{}),

		colly.Async(true),
	)
	extensions.RandomUserAgent(collector)
	collector.Limit(
		&colly.LimitRule{
			DomainRegexp: "",
			DomainGlob:   "",
			Delay:        5 * time.Second,
			RandomDelay:  5 * time.Second,
			Parallelism:  2,
		})

	// Find and visit all links
	collector.OnHTML("body", github.BaseProjectInformations)
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	collector.Visit(targetURL)
	collector.Wait()
}
