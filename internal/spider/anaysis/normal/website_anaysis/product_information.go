/**
产品信息页面
*/
package website_anaysis

import (
	"fmt"
	"github.com/gocolly/colly"
)

/**
产品描述解析
*/
func ProductInformation(elment *colly.HTMLElement) {
	queryDom := elment.DOM
	fmt.Println("begin to handle html dom...")

	fmt.Println(queryDom.Find(" span").Children().Text())
}
