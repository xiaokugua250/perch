package main

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	colly "github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"os"
	database "perch/database/mysql"
	"perch/internal/spider/anaysis/website_anaysis/github"
	"perch/tools"

	"sort"

	"strings"
	"time"
)

func FileLinesScanner(fileName string) ([]string, error) {
	var results []string
	file, err := os.Open(fileName)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		results = append(results, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		//log.Fatal(err)
		return nil, err
	}
	return results, nil
}

func main() {
var (
	basicInfos     []github.BasicInfo
	err error
)
	/*CollySpider := &crawl.CrawlColly{}
	CollySpider.CollyInit()
	CollySpider.Collystart()*/
	database.InitMySQLDB()


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
	collector.OnHTML("div", BasictInformations)
	//collector.OnHTML("div",github.AdvancedInformationsWithFileList)
	collector.OnResponse(func(response *colly.Response) {
		collector.OnError(func(response *colly.Response, err error) {
			fmt.Printf("error in visit url %s is %s\n", response.Request.URL, err)
		})

	})
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})


	err = database.MysqlDb.Limit(-1).Where("oldest_commit_at=0").Find(&basicInfos).Error
	if err != nil {
		log.Fatalln(err)
	}
	for _, item := range basicInfos {
		err = collector.Visit(item.Link+"/file-list/master")

		if err != nil {
			fmt.Printf("error in visist url %s is %s\n", item.Link, err)
		}
		collector.Wait()

	}

}

func BasictInformations(elment *colly.HTMLElement) {
	var (
		basicInfo github.BasicInfo
		err       error
	)
	time.Sleep(1 * time.Second)
	doc := elment.DOM

	basicBodyDom := doc




	var CommitTimeLines tools.TimeSlice

	basicBodyDom.Find("time-ago").Each(func(i int, selection *goquery.Selection) {
		commitHistorys, _ := selection.Attr("datetime")

		//fmt.Printf("commit date is %s,%b\n",commitHistorys,ok)
		formatedTime, err := time.Parse("2006-01-02T15:04:05Z", commitHistorys)
		if err != nil {
			log.Error(err)
		}
		CommitTimeLines = append(CommitTimeLines, formatedTime)

	})

	sort.Sort(CommitTimeLines)
	//fmt.Printf("commit line is %s\n",CommitTimeLines)
	//basicInfo.CommitTimeLines = CommitTimeLines
	if CommitTimeLines.Len() != 0 {
		basicInfo.OldestCommitAt = CommitTimeLines[0].Unix()
		basicInfo.NewestCommitAt = CommitTimeLines[CommitTimeLines.Len()-1].Unix()
	}



	fmt.Printf("%+v\n",CommitTimeLines)
	/*if err =database.MySQL_DB.Create(&basicInfo).Error;err!= nil{
		log.Printf("error in create basic info is %s\n",err)
	}*/

	updated_map := make(map[string]interface{})
	updated_map["oldest_commit_at"] = basicInfo.OldestCommitAt
	updated_map["newest_commit_at"] = basicInfo.NewestCommitAt
	if basicInfo.NewestCommitAt != 0 {
		link := strings.TrimSuffix(elment.Request.URL.String(),"/file-list/master")
		if err = database.MysqlDb.Model(&basicInfo).Where("link=?", link).Updates(&updated_map).Error; err != nil {
			log.Error(err)
		}

	}

	//	fmt.Printf("basic info is %+v\n",basicInfo)
	//Log.Printf(" basic info is %+v\n", basicInfo)
	//log.Infof(" %+v\n", basicInfo)
}

