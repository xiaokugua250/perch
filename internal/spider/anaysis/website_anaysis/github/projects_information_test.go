package github

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestBaseProjectInformations(t *testing.T) {

	file, err := os.Open("github_projects.html")
	if err != nil {
		log.Error(err)
	}
	doc, err := goquery.NewDocumentFromReader(bufio.NewReader(file))
	Readme := doc.Find("div#readme")
	//fmt.Printf("readme text is %s\n",Readme.Text())
	Readme.Find("li").Has("a").Children().Each(func(i int, selection *goquery.Selection) {
		if i > 100 {
			//selection.Each()
			url, ok := selection.Attr("href")
			if ok {

			}
			fmt.Printf(" index is %d ,readme text is %s,url is %s ,descriptions is %s\n", i, selection.Text(), url, selection.Parent().Text())
		}
		//if selection.Children()
	})

}
