package main

import (
	"fmt"
	colly "github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	database "perch/database/mysql"
	"perch/internal/spider/anaysis/website_anaysis/github"
	"time"
)

func main() {

	/*CollySpider := &crawl.CrawlColly{}
	CollySpider.CollyInit()
	CollySpider.Collystart()*/
handler:

	var (
		basicInfos     []github.BasicInfo
		unHandledCount int64
		err            error
	)
	//	urls, err:= tools.FileLinesScanner("E:\\WorksSpaces\\GoWorkSpaces\\perch\\internal\\spider\\github_projects_url.txt")

	if err != nil {
		fmt.Println("error is ", err)
		log.Error(err)
	}

	/*q, _ := queue.New(
		2, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)*/
	//targetURL := "https://github.com/avelino/awesome-go"
	collector := colly.NewCollector(
		//colly.Debugger(&debug.LogDebugger{}),
		//	colly.AllowedDomains("github.com"),
		colly.Async(true),
		colly.IgnoreRobotsTxt(),
	)
	collector.AllowURLRevisit = false
	extensions.RandomUserAgent(collector)
	extensions.Referer(collector)
	collector.Limit(
		&colly.LimitRule{
			//DomainGlob:  "github.com/*",
			//DomainRegexp: " DomainGlob:  \"godoc.org/*\",",
			//	DomainGlob:   "",

			Delay:       5 * time.Second,
			RandomDelay: 3 * time.Second,
			Parallelism: 2,
		})
	collector.WithTransport(&http.Transport{
		//	Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 10 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 3 * time.Second,
	})

	// Find and visit all links
	// ref https://github.com/gocolly/colly/blob/master/_examples/instagram/instagram.go
	//在爬取过程中，可能出现html元素还没加载完的情况就开始处理
	collector.OnHTML("html", github.BasictInformations)
	//collector.OnHTML("div",github.AdvancedInformationsWithFileList)
	collector.OnResponse(func(response *colly.Response) {
		collector.OnError(func(response *colly.Response, err error) {
			fmt.Printf("error in visit url %s is %s\n", response.Request.URL, err)
		})
		//= goquery.NewDocumentFromReader(response.Body)
		/*	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(response.Body))
			if err!= nil{
				log.Printf("error in parase url %s ,error is %s\n",response.Request.URL,err)
			}*/

	})
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	var project_github int64
	err = database.MySQL_DB.Limit(-1).Where("oldest_commit_at=0").Find(&basicInfos).Error
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v", basicInfos)
	for _, item := range basicInfos {
		project_github += 1
		//fmt.Printf("item is %+v\n",item)
		err = collector.Visit(item.Link)
		//	err = collector.Visit(item.Link+"/file-list/master")

		collector.Wait()
		/*for _,url:= range urls{
		if strings.Contains(url,"https://github.com"){
			fmt.Printf("begin to visit url %s\n",url)
			//q.AddURL(url)
			err = collector.Visit(url)
			project_github+=1
			if err!= nil{
				fmt.Printf("error in visist url %s is %s\n",url,err)
			}
			collector.Wait()
		/*	if url =="https://github.com/Comcast/gaad"{
				err = collector.Visit(url)
				if err!= nil{
					fmt.Printf("error in visist url %s is %s\n",url,err)
				}


			}*/

	}
	/*collector_clone := collector.Clone()
	for _,item := range basicInfos{

		//fmt.Printf("item is %+v\n",item)
		err = collector_clone.Visit(item.Link+"/file-list/master")
		//	err = collector.Visit(item.Link+"/file-list/master")

		collector_clone.Wait()
		/*for _,url:= range urls{
		if strings.Contains(url,"https://github.com"){
			fmt.Printf("begin to visit url %s\n",url)
			//q.AddURL(url)
			err = collector.Visit(url)
			project_github+=1
			if err!= nil{
				fmt.Printf("error in visist url %s is %s\n",url,err)
			}
			collector.Wait()
		/*	if url =="https://github.com/Comcast/gaad"{
				err = collector.Visit(url)
				if err!= nil{
					fmt.Printf("error in visist url %s is %s\n",url,err)
				}


			}


	}

	collector_clone.OnHTML("",github.AdvancedInformationsWithFileList)*/
	fmt.Print("gityhub project is ", project_github)
	if err = database.MySQL_DB.Model(&github.BasicInfo{}).Where("oldest_commit_at=0").Count(&unHandledCount).Error; err != nil {
		log.Error(err)
	}
	if unHandledCount >= 5 {
		goto handler
	}

	//collector.Visit("https://github.com/avelino/awesome-go")

	//collector.Visit(targetURL)

}
