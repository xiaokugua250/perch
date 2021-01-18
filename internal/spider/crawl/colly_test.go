/**
ref https://medium.com/swlh/increase-your-scraping-speed-with-go-and-colly-the-basics-41038bc3647e
*/

package crawl

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"testing"
)

func TestCrawlColly_CollyElmentHTMLHandler(t *testing.T) {
	/*collerSpider := CrawlColly{}
	collerSpider.CollyInit()
	collerSpider.CollyElmentHTMLHandler()
	collerSpider.CollyVisit("https://www.amazon.com/s?k=nintendo+switch&ref=nb_sb_noss_1")*/

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("div.s-result-list.s-search-results.sg-row", func(e *colly.HTMLElement) {
		e.ForEach("div.a-section.a-spacing-medium", func(_ int, e *colly.HTMLElement) {
			var productName string

			productName = e.ChildText("span.a-size-medium.a-color-base.a-text-normal")

			if productName == "" {
				// If we can't get any name, we return and go directly to the next element
				fmt.Println("====")
				return
			}

			fmt.Printf("Product Name: %s \n", productName)
		})
	})

	err := c.Visit("https://www.amazon.com/s?k=nintendo+switch&ref=nb_sb_noss_1")
	if err != nil {
		log.Print(err)
	}
}
