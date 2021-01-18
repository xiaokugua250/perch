/**
html 页面元素的解析器
ref https://github.com/PuerkitoBio/goquery
*/
package anaysis

import "github.com/gocolly/colly"

type SpiderAnaysier interface {
	SpiderParser(selector interface{}) (interface{}, error)
}

func GeneralOnErrorFunc(response *colly.Response, err error) {

}

/**
html dom 元素解析
*/
func HtmlDomHandler() {

}

/**
xml 元素解析
*/
func XmlDomHandler() {

}

/**
json dom 解析
*/
func JsonDomHander() {

}
