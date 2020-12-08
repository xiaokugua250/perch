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
	"perch/tools"
	"regexp"
	"sort"
	strconv "strconv"
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

	/*CollySpider := &crawl.CrawlColly{}
	CollySpider.CollyInit()
	CollySpider.Collystart()*/

	var (
		err error
	)
	dir, err := os.Getwd()
	fmt.Printf("%s\n", dir)
	urls, err := FileLinesScanner("url.txt")

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
			Parallelism: 5,
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
	collector.OnHTML("html", BasictInformations)
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
	/*err = database.MySQL_DB.Limit(-1).Where("oldest_commit_at=0").Find(&basicInfos).Error
	if err!= nil{
		log.Fatalln(err)
	}
	fmt.Printf("%+v",basicInfos)
	*/

	for _, url := range urls {
		if strings.Contains(url, "https://github.com") {
			fmt.Printf("begin to visit url %s\n", url)
			//q.AddURL(url)
			err = collector.Visit(url)
			project_github += 1
			if err != nil {
				fmt.Printf("error in visist url %s is %s\n", url, err)
			}
			collector.Wait()
			/*	if url =="https://github.com/Comcast/gaad"{
				err = collector.Visit(url)
				if err!= nil{
					fmt.Printf("error in visist url %s is %s\n",url,err)
				}


			}*/

		}
	}
}

func BasictInformations(elment *colly.HTMLElement) {
	var (
		basicInfo BasicInfo
		err       error
	)
	time.Sleep(1 * time.Second)
	doc := elment.DOM
	basicHeadDom := doc.Find("head")
	basicBodyDom := doc.Find("body")
	stars, exists := basicBodyDom.Find("a.social-count.js-social-count").Attr("aria-label")
	if !exists {
		log.Error("find project stars failed....")
	}
	re := regexp.MustCompile("[0-9]+")
	starNums := re.FindAllString(stars, -1)

	basicHeadDom.Find("meta").Each(func(i int, selection *goquery.Selection) {
		og_property, ok := selection.Attr("property")
		if ok {
			if og_property == "og:title" {
				title, ok := selection.Attr("content")
				if ok {
					basicInfo.Title = title
				}

			}
			if og_property == "og:url" {
				url, ok := selection.Attr("content")
				if ok {
					basicInfo.Link = url
				}
			}

		}
	})
	basicInfo.Description = basicBodyDom.Find("p.f4.mt-3").Text()
	basicInfo.Description = strings.Replace(basicInfo.Description, "\n", "", -1)
	if len(starNums) != 0 {
		basicInfo.Stars, err = strconv.ParseInt(starNums[0], 10, 64)
	} else {
		fmt.Println("error link is ", elment.Request.URL)
		//return
		//	//basicInfo.Stars = 0
	}

	if err != nil {
		log.Error("convert stars failed ....")
	}
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

	/*if err =database.MySQL_DB.Create(&basicInfo).Error;err!= nil{
		log.Printf("error in create basic info is %s\n",err)
	}*/

	updated_map := make(map[string]interface{})
	updated_map["oldest_commit_at"] = basicInfo.OldestCommitAt
	updated_map["newest_commit_at"] = basicInfo.NewestCommitAt
	if basicInfo.NewestCommitAt != 0 {
		if err = database.MySQL_DB.Model(&basicInfo).Where("link=?", basicInfo.Link).Updates(&updated_map).Error; err != nil {
			log.Error(err)
		}

	}

	//	fmt.Printf("basic info is %+v\n",basicInfo)
	//Log.Printf(" basic info is %+v\n", basicInfo)
	log.Infof(" %+v\n", basicInfo)
}

type BasicInfo struct {
	ID              int         `json:"id"  gorm:"column:id;type:int(11);primary_key;not null"`
	Name            string      `json:"project_name" gorm:"column:name;type:varchar(128);primary_key;not null"`
	Link            string      `json:"link" gorm:"column:link;type:varchar(128)"`
	Title           string      `json:"title" gorm:"column:title;type:varchar(128)"`
	Description     string      `json:"description" gorm:"column:description;type:varchar(255)"`
	Stars           int64       `json:"stars" gorm:"column:stars;type:varchar(128)"`
	License         string      `json:"license" gorm:"-"`
	Contributors    []string    `json:"contributors" gorm:"-"`
	Languages       []string    `json:"languages" gorm:"-"`
	CommitTimeLines []time.Time `json:"commit_time_lines" gorm:"-"`
	NewestCommitAt  int64       `json:"last_commit_at" gorm:"column:newest_commit_at;type:int(11);"` //
	OldestCommitAt  int64       `json:"oldest_commit_at" gorm:"column:oldest_commit_at;type:int(11);"`
}
