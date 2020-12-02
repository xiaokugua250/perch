package github

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"perch/tools"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type BasicInfo struct {
	Name            string      `json:"project_name"`
	Title           string      `json:"title"`
	Description     string      `json:"description"`
	Stars           int64       `json:"stars"`
	CommitTimeLines []time.Time `json:"commit_time_lines"`
}

func BasictInformations(elment *colly.HTMLElement) {
	var (
		basicInfo BasicInfo
		err       error
	)
	doc := elment.DOM
	stars, exists := doc.Find("a.social-count.js-social-count").Attr("aria-label")
	if !exists {
		log.Error("find project stars failed....")
	}
	re := regexp.MustCompile("[0-9]+")
	starNums := re.FindAllString(stars, -1)

	titleAndDescription := doc.Find("title").Text()
	fmt.Printf("===%s\n", titleAndDescription)
	basicInfo.Name = strings.Split(titleAndDescription, ":")[0]
	basicInfo.Description = doc.Find(".f4.mb-3").Text()
	basicInfo.Stars, err = strconv.ParseInt(starNums[0], 10, 64)
	if err != nil {
		log.Error("convert stars failed ....")
	}
	var CommitTimeLines tools.TimeSlice
	doc.Find("time-ago").Each(func(i int, selection *goquery.Selection) {
		commitHistorys, _ := selection.Attr("datetime")
		formatedTime, err := time.Parse("2006-01-02T15:04:05Z", commitHistorys)
		if err != nil {
			log.Error(err)
		}
		CommitTimeLines = append(CommitTimeLines, formatedTime)

	})

	sort.Sort(CommitTimeLines)
	basicInfo.CommitTimeLines = CommitTimeLines
	log.Printf(" basic info is %+v\n", basicInfo)
}

func AdvancedInformationsWithTag(elment *colly.HTMLElement) {
	queryDom := elment.DOM.Find("#readme")

	queryDom.Each(
		func(i int, selection *goquery.Selection) {
			projectsInfo := selection.Find("ul")
			projectsInfo.Children()

			fmt.Printf("html is %s\n", projectsInfo.Children().Has("li").Has("a").Text())

			//	fmt.Println("projects info is %",projectsInfo.Text(),"---\n")

			//fmt.Println("index is %d,text is %s\n",i,selection.Find("ul>li").Text())
		})
	fmt.Println("begin to handle html dom...")

}
