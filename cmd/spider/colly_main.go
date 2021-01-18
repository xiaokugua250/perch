package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"time"
)

func main() {
	startTime := time.Date(2020, 11, 18, 10, 16, 30, 55, time.Local)
	for {
		if time.Now().After(startTime) && time.Now().Before(startTime.Add(time.Minute*1)) {
			for {
				c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))

				c.OnRequest(func(r *colly.Request) {
					c.OnRequest(func(r *colly.Request) {
						r.Headers.Set("Cookie", "sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%22175cebcc73513e-03c3167705e63e8-4c3f2678-2359296-175cebcc736308%22%2C%22%24device_id%22%3A%22175cebcc73513e-03c3167705e63e8-4c3f2678-2359296-175cebcc736308%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_referrer%22%3A%22%22%2C%22%24latest_referrer_host%22%3A%22%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%7D%7D")
						r.Headers.Add("referer", "https://www.youcash.com/wechat-web/html/psbcTribleGift/index.html?channelCode=1000000499999999&activityId=3fc213ae83a74ea299148073796cdeb7&usage=1&override=0")
						r.Headers.Add("accept", "text/javascript, application/javascript, application/ecmascript, application/json, text/plain, */*")
						r.Headers.Add("accept-language", "zh,en-US;q=0.7,en;q=0.3")

					})
					fmt.Println("Visiting", r.URL)
				})
				c.WithTransport(&http.Transport{
					//Proxy: http.ProxyFromEnvironment,
					DialContext: (&net.Dialer{
						Timeout:   30 * time.Second, // 超时时间
						KeepAlive: 30 * time.Second, // keepAlive 超时时间
					}).DialContext,
					MaxIdleConns:          100,              // 最大空闲连接数
					IdleConnTimeout:       90 * time.Second, // 空闲连接超时
					TLSHandshakeTimeout:   10 * time.Second, // TLS 握手超时
					ExpectContinueTimeout: 1 * time.Second,
				})
				c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:82.0) Gecko/20100101 Firefox/82.0"
				//extensions.RandomUserAgent(c)
				//extensions.Referer(c)

				c.OnError(func(_ *colly.Response, err error) {
					log.Println("Something went wrong:", err)
				})

				c.OnResponse(func(r *colly.Response) {
					fmt.Println("Visited", r.Request.URL)
				})
				c.OnResponse(func(response *colly.Response) {
					//fmt.Println("-=----", response.StatusCode)
					fmt.Printf("response from %s is %s\n", response.Request.URL, string(response.Body))
				})
				err := c.Post("https://www.youcash.com/wechat-web/limitTimeCoupon/currentTime", nil)
				//err :=c.Visit("https://www.youcash.com/wechat-web/html/psbcTribleGift/index.html?channelCode=1000000499999999&activityId=3fc213ae83a74ea299148073796cdeb7&usage=1&override=0#/home")
				if err != nil {
					fmt.Println("error is", err)
				}
				time.Sleep(3 * time.Nanosecond)
			}
		}
	}

}
