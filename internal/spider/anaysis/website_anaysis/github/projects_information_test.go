package github

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"os"
	"perch/tools"
	"regexp"
	"sort"
	"strings"
	"testing"
	"time"
)

func TestBaseProjectInformations(t *testing.T) {

	file, err := os.Open("github_projects.html")
	if err != nil {
		log.Error(err)
	}
	doc, err := goquery.NewDocumentFromReader(bufio.NewReader(file))

	//获取star
	stars, _ := doc.Find("a.social-count.js-social-count").Attr("aria-label")
	re := regexp.MustCompile("[0-9]+")
	starNums := re.FindAllString(stars, -1)

	titleAndContents := doc.Find("title").Text()
	title := strings.Split(titleAndContents, ":")[0]
	//content:= strings.Split(titleAndContents,":")[1]
	content := doc.Find(".f4.mb-3").Text()
	var CommitHistoryTimeArray tools.TimeSlice
	fmt.Printf("project star is %s, nums is %s,project title  is %s and content is %s\n", stars, starNums[0], title, content)
	doc.Find("time-ago").Each(func(i int, selection *goquery.Selection) {
		commitHistorys, _ := selection.Attr("datetime")
		//	fmt.Printf("commit_history is %s\n", commitHistorys)
		//loc, _ := time.LoadLocation("Local")                            //重要：获取时区
		//formatedTime,err:= time.ParseInLocation(commitHistorys,"2006-01-02T15:04:05-0700", loc)
		formatedTime, err := time.Parse("2006-01-02T15:04:05Z", commitHistorys)
		if err != nil {
			log.Error(err)
		}
		CommitHistoryTimeArray = append(CommitHistoryTimeArray, formatedTime)

		//	fmt.Printf("time is %+v",formatedTime)
	})

	sort.Sort(CommitHistoryTimeArray)
	fmt.Printf("commit history is %+v\n", CommitHistoryTimeArray)

}
