/**
todo  要解决的问题有:
robot.txt 文件的解析和反扒设置 参考	https://github.com/samclarke/robotstxt

proxy 的设置	ref https://smartproxy.com/what-is-web-scraping/scraping-amazon-with-parsehub

cookie处理

UserAgent的处理 需要随机处理user-agent collector 的配置可以在爬虫执行到任何阶段改变。一个经典的例子，通过随机改变 user-agent，可以帮助我们实现简单的反爬

html 元素的处理 refer https://github.com/PuerkitoBio/goquery

* Rate limiting * Parallel crawling * Respecting robots.txt * HTML/Link parsing
启用异步加快任务执行
colly 默认会阻塞等待请求执行完成，这将会导致等待执行任务数越来越大。我们可以通过设置 collector 的 Async 选项为 true 实现异步处理，从而避免这个问题。如果采用这种方式，记住增加 c.Wait()，否则程序会立刻退出。
 */

package crawl

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/mottet-dev/medium-go-colly-basics/utils"
	"github.com/gocolly/colly/extensions"
	log "github.com/sirupsen/logrus"
	_ "perch/pkg/log"
)


type CrawlColly struct {
	Coller *colly.Collector
	TargetUrls []string
}

/**
colly init方法
 */
func ( crawColly *CrawlColly)CollyInit(){

	crawColly.Coller = colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36"), //todo 需要添加浏览器AGENT信息
		colly.IgnoreRobotsTxt(), //// IgnoreRobotsTxt allows the Collector to ignore any restrictions set by // the target host's robots.txt file.  See http://www.robotstxt.org/ for more // information.
		colly.AllowURLRevisit(),
		)
	extensions.RandomUserAgent(crawColly.Coller)	//todo
	extensions.Referer(crawColly.Coller)
	crawColly.Coller.OnRequest(func(request *colly.Request) {
		/**
		//todo 对请求头进行处理
		r.Headers.Set("Host", "query.sse.com.cn")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Accept", "*//*")
		r.Headers.Set("Origin", "http://www.sse.com.cn")
		r.Headers.Set("Referer", "http://www.sse.com.cn/assortment/stock/list/share/") //关键头 如果没有 则返回 错误
		r.Headers.Set("Accept-Encoding", "gzip, deflate")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
		 */
		fmt.Println("Visiting", request.URL)
	})
	//crawColly.Coller.Visit("http://go-colly.org/")
}
/**
爬虫规则进行设置
   // Filter domains affected by this rule
    DomainGlob:  "godoc.org/*",
    // Set a delay between requests to these domains
    Delay: 1 * time.Second
    // Add an additional random delay
    RandomDelay: 1 * time.Second,
Parallelism:5
 */
func ( crawColly *CrawlColly)CollyLimitConfig(rules []*colly.LimitRule)(error){

	return  crawColly.Coller.Limits(rules)

}



/**
todo  在发起请求前被调用
 */
func ( crawColly *CrawlColly)CollyOnRequestHandler(crawRequestHandler colly.RequestCallback ){
	crawColly.Coller.OnRequest(crawRequestHandler)
}

/**
todo  请求过程中如果发生错误被调用
*/
func ( crawColly *CrawlColly)CollyOnErrorHandler(crawErrorHandler colly.ErrorCallback ){
	crawColly.Coller.OnError(crawErrorHandler)
}

/**
todo  收到回复后被调用
*/
func ( crawColly *CrawlColly)CollyOnResponseHandler(crawResponseHandler colly.ResponseCallback){
	crawColly.Coller.OnResponse(crawResponseHandler )
}

/**
todo   在 OnResponse 之后被调用，如果收到的内容是 HTML
*/
func ( crawColly *CrawlColly)CollyOnHTMLHandler(htmlSelector string,crawHtmlHandler colly.HTMLCallback ){
	crawColly.Coller.OnHTML(htmlSelector,crawHtmlHandler)
}

/**
todo   在 OnHTML 之后被调用
*/
func ( crawColly *CrawlColly)CollyOnScrapedtHandler(scrapedHandler colly.ScrapedCallback ){
	crawColly.Coller.OnScraped(scrapedHandler)
}

func ( crawColly *CrawlColly)CollyElmentHTMLHandler(){

	crawColly.Coller.OnHTML("div.s-result-list.s-search-results.sg-row", func(e *colly.HTMLElement) {
		e.ForEach("div.a-section.a-spacing-medium", func(_ int, e *colly.HTMLElement) {
			var productName, stars, price string

			productName = e.ChildText("span.a-size-medium.a-color-base.a-text-normal")

			if productName == "" {
				// If we can't get any name, we return and go directly to the next element
				return
			}

			stars = e.ChildText("span.a-icon-alt")
			utils.FormatStars(&stars)

			price = e.ChildText("span.a-price > span.a-offscreen")
			utils.FormatPrice(&price)

			fmt.Printf("Product Name: %s \nStars: %s \nPrice: %s \n", productName, stars, price)
		})
	})
}




func ( crawColly *CrawlColly)CollyVisit(urls ... string)(error){
	var (
		err error
	)
	if len(urls)==0{
		return errors.New(fmt.Sprintf("urls %v is empty",urls))
	}
	for _,url:= range urls{

		go func(url string) {
			err =crawColly.Coller.Visit(url)
			log.Errorf(fmt.Sprintf("error %s happend at visiting url %s",err.Error(),url))
		}(url)
	}
	crawColly.Coller.Wait()	 //todo 但是多用了Wait方法，这是因为在Async为true时需要等待协程都完成再结束
	return nil

}
