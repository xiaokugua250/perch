package github

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/rifflock/lfshook"
	database "perch/database/mysql"

	//"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"perch/tools"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var Log *log.Logger

func NewLogger() *log.Logger {
	if Log != nil {
		return Log
	}

	pathMap := lfshook.PathMap{
		log.InfoLevel: "E:\\WorksSpaces\\GoWorkSpaces\\perch\\internal\\spider\\anaysis\\website_anaysis\\github\\info.log",

		log.PanicLevel: "panic.log",
	}

	Log = log.New()

	Log.Hooks.Add(lfshook.NewHook(pathMap, &log.JSONFormatter{}))

	return Log
}

func init() {
	fmt.Printf("begin to init log at %s\n", time.Now().String())
	Log = NewLogger()
	database.InitMySQLDB()
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

func (BasicInfo) TableName() string {
	return "github_projects"
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
	Log.Infof(" %+v\n", basicInfo)
}

func AdvancedInformationsWithFileList(elment *colly.HTMLElement) {
	var (
		basicInfo BasicInfo
		err       error
	)

	fmt.Printf("request url is  %+v,%s\n", elment.Request.URL, strings.TrimSuffix(elment.Request.URL.String(), "/file-list/master"))
	time.Sleep(1 * time.Second)
	doc := elment.DOM
	//	basicHeadDom := doc.Find("head")
	basicBodyDom := doc

	var CommitTimeLines tools.TimeSlice

	basicBodyDom.Find("time-ago").Each(func(i int, selection *goquery.Selection) {
		commitHistorys, _ := selection.Attr("datetime")

		fmt.Printf("commit date is %s\n", commitHistorys)
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
		if err = database.MySQL_DB.Model(&basicInfo).Where("link=?", strings.TrimSuffix(elment.Request.URL.String(), "/file-list/master")).Updates(&updated_map).Error; err != nil {
			log.Error(err)
		}
	}

	//	fmt.Printf("basic info is %+v\n",basicInfo)
	//Log.Printf(" basic info is %+v\n", basicInfo)
	Log.Infof(" %+v\n", basicInfo, err)
}

func AdvancedInformationsWithTag(elment *colly.HTMLElement) {
	//queryDom := elment.DOM.Find("#readme")

	file, err := os.Open("github_projects.html")
	if err != nil {
		log.Error(err)
	}
	doc, err := goquery.NewDocumentFromReader(bufio.NewReader(file))
	Readme := doc.Find("div#readme")
	//fmt.Printf("readme text is %s\n",Readme.Text())
	targetFile, err := os.OpenFile("github_projects.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Errorln(err)
	}
	//	defer target_file.Close()
	write := bufio.NewWriter(targetFile)
	Readme.Find("li").Has("a").Children().Each(func(i int, selection *goquery.Selection) {
		if i > 100 {
			//selection.Each()
			_, ok := selection.Attr("href")
			if ok {
				url, _ := selection.Attr("href")

				str := fmt.Sprintf("project name is: %s, url is: %s, project description is: %s\n", selection.Text(), url, selection.Parent().Text())
				//fmt.Printf("str is %s\n",str)
				_, err := write.WriteString(str)
				if err != nil {
					log.Errorln(err)
				}
				fmt.Printf(" index is %d ,readme text is %s,url is %s ,descriptions is %s\n", i, selection.Text(), url, selection.Parent().Text())

				//rr=ioutil.WriteFile("github_projects.txt",[]byte(fmt.Sprintf("project name is: %s, url is: %s, project description  is %s\n",selection.Text(), url, selection.Parent().Text())),os.ModeAppend|0644)

				//fmt.Printf(" index is %d ,readme text is %s,url is %s ,descriptions is %s\n", i, )
			}
		}
		//if selection.Children()
	})
	write.Flush()

}
