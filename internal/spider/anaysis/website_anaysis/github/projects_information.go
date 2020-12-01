package github

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func BaseProjectInformations(elment *colly.HTMLElement) {
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

func BaseReadMeInformations(elment *colly.HTMLElement) {

}
