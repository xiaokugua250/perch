/**
分布式爬虫，可以从几个层面考虑，分别是代理层面、执行层面和存储层面。
代理层面：
通过设置代理池，我们可以将下载任务分配给不同节点执行，有助于提供爬虫的网页下载速度。
同时，这样还能有效降低因爬取速度太快而导致IP 被禁的可能性。
执行层面
这种方式通过将任务分配给不同的节点执行，实现真正意义的分布式。
如果实现分布式执行，首先需要面对一个问题，如何将任务分配给不同的节点，实现不同任务节点之间的协同工作呢？
首先，我们选择合适的通信方案。常见的通信协议有 HTTP、TCP，一种无状态的文本协议、一个是面向连接的协议。除此之外，还可选择的有种类丰富的 RPC 协议，比如 Jsonrpc、facebook 的 thrift、google 的 grpc 等。
文档提供了一个 HTTP 服务示例代码，负责接收请求与任务执行。

存储层面
我们已经通过将任务分配到不同节点执行实现了分布式。但部分数据，比如 cookies、访问的 url 记录等，在节点之间需要共享。默认情况下，这些数据是保存内存中的，只能是每个 collector 独享一份数据。
我们可以通过将数据保存至 redis、mongo 等存储中，实现节点间的数据共享。colly 支持在任何存储间切换，只要相应存储实现 colly/storage.Storage 接口中的方法。
其实，colly 已经内置了部分 storage 的实现，查看 storage。下一节也会谈到这个话题。

useragent 信息可以参考 https://developers.whatismybrowser.com/useragents/database/
https://blog.hartleybrody.com/scrape-amazon/#code
*/
package crawl

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

type SpiderCrawer interface {
	CrawlerFetch(url string) error
}

func CrawlerFetch(url string, method string, depth int, requestData io.Reader, ctx *context.Context, requestHeader http.Header, checkRevisited bool) error {
	var err error
	httpClient := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("request url %s,response is %s\n", url, string(body))

	return nil
}

func CrawlSpiderCollyInitWithOptions(options ...CollyOptionFunc) (*colly.Collector, error) {
	var (
		collyCollector *colly.Collector
		err   error
	)
	collyCollector = colly.NewCollector()

	defaultOptions := &CollyConfigOptions{
		IgnoreRobotsTxt:true,
		AllowRevisitURL:false,
		UseRandomeUserAgent:true,
	}
	for _, opt := range options {
		//opt(&defaultOptions)
		opt(defaultOptions)
	}

	collyCollector.AllowURLRevisit = defaultOptions.AllowRevisitURL
	if defaultOptions.UseRandomeUserAgent {
		extensions.RandomUserAgent(collyCollector)
		extensions.Referer(collyCollector)
	}
	collyCollector.IgnoreRobotsTxt = defaultOptions.IgnoreRobotsTxt

	return collyCollector, err
}

func CrawlSpiderCollyStart(taskConf CrawTaskConfig) error {
	var (
		err error
		collector *colly.Collector
	)
	collector,err = CrawlSpiderCollyInitWithOptions(nil)
	collector.DisallowedDomains=taskConf.TaskDisallowDomains


	switch taskConf.TaskType {
	case CrawTaskType_Snap:
		//todo
	case CrawTaskType_Complex:
			//todo
	case CrawTaskType_Normal:
		//todo

	}


	return err
}
