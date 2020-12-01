/**
产品信息页面
*/
package website_anaysis

import (
	"fmt"
	"github.com/gocolly/colly"
)

/**
项目描述信息
*/
func ProductsInformation(elment *colly.HTMLElement) {
	queryDom := elment.DOM
	queryDom.Find("logged-in env-production page-responsive")

	fmt.Println("begin to handle html dom...")

	fmt.Println(queryDom.Find(" span").Children().Text())
}
