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
	targetFile, err := os.OpenFile("github_projects.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Errorln(err)
	}
	//	defer target_file.Close()
	write := bufio.NewWriter(targetFile)
	Readme.Find("li").Has("a").Children().Each(func(i int, selection *goquery.Selection) {
		if i > 100 {
			//selection.Each()
			url, _ := selection.Attr("href")

			str := fmt.Sprintf("project name is: %s, url is: %s, project description is: %s\n", selection.Text(), url, selection.Parent().Text())
			//fmt.Printf("str is %s\n",str)
			_, err := write.WriteString(str)
			if err != nil {
				log.Errorln(err)
			}

			//rr=ioutil.WriteFile("github_projects.txt",[]byte(fmt.Sprintf("project name is: %s, url is: %s, project description  is %s\n",selection.Text(), url, selection.Parent().Text())),os.ModeAppend|0644)

			//fmt.Printf(" index is %d ,readme text is %s,url is %s ,descriptions is %s\n", i, )
		}

		//if selection.Children()
	})
	write.Flush()

}
