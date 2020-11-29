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
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"perch/internal/spider/anaysis"
)

func CrawSpiderWithSnapTask(collecter *colly.Collector,url string)(error){
	var (
		err error
	)

	collecter.OnError(anaysis.GeneralOnErrorFunc)

	err =collecter.Visit(url)
	if err!= nil{
		log.Error(err)
	}
	return err
}

