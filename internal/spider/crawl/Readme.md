# 下载器
- 发送HTTP请求或获取HTTP响应
https://rapidapi.com/search/amazon

多收集器
我们前面演示的爬虫都是比较简单的，处理逻辑都很类似。如果是一个复杂的爬虫，我们可以通过创建不同的 collector 负责不同任务的处理。

如何理解这段话呢？举个例子吧。

如果大家写过一段时间爬虫，肯定遇到过父子页面抓取的问题，通常父页面的处理逻辑与子页面是不同的，并且通常父子页面间还有数据共享的需求。用过 scrapy 应该知道，scrapy 通过在 request 绑定回调函数实现不同页面的逻辑处理，而数据共享是通过在 request 上绑定数据实现将父页面数据传递给子页面。

研究之后，我们发现 scrapy 的这种方式 colly 并不支持。那该怎么做？这就是我们要解决的问题。

对于不同页面的处理逻辑，我们可以定义创建多个收集器，即 collector，不同 collector 负责处理不同的页面逻辑。

c := colly.NewCollector(
    colly.UserAgent("myUserAgent"),
    colly.AllowedDomains("foo.com", "bar.com"),
)
// Custom User-Agent and allowed domains are cloned to c2
c2 := c.Clone()
通常情况下，父子页面的 collector 是相同的。上面的示例中，子页面的 collector c2 通过 clone，将父级 collector 的配置也都复制了下来。

而父子页面之间的数据传递，可以通过 Context 实现在不同 collector 之间传递。注意这个 Context 只是 colly 实现的数据共享的结构，并非 Go 标准库中的 Context。

c.OnResponse(func(r *colly.Response) {
    r.Ctx.Put("Custom-header", r.Headers.Get("Custom-Header"))
    c2.Request("GET", "https://foo.com/", nil, r.Ctx, nil)
})
如此一来，我们在子页面中就可以通过 r.Ctx 获取到父级传入的数据了。关于这个场景，我们可以查看官方提供的案例 coursera_courses。

配置优化
