package crawl

import "net/http"

/**

任务类型
*/
*/
type CrawTaskType struct{
	TaskType string `json:"taskType"` //任务类型
}
type CrawlRequest struct {
	CrawHttpRequest *http.Request `json:"craw_http_request"`
	CrawDepth       int32         `json:"craw_depth"` //爬虫深度
}

type CrawResponse struct {
	CrawHttpRespone *http.Response `json:"craw_http_respone"`
	CrawDepth       int32          `json:"craw_depth"`
}

type CrawResult struct { // 爬虫获取结果
	ResultItem map[string]interface{}
	RawResult  interface{} //原始获取结果
}

type SpiderValider interface {
	SpiderValider() bool
}

const (
	SPIDER_NETWORK_ERROR       = "network error"
	SPIDER_SCHEDLER_ERROR      = "scheduler error"      //调度器错误
	SPIDER_RESULTHANDLER_ERROR = "result handler error" //爬虫结果错误
	SPIDER_ANAYSIZER_ERROR     = "anayzer error"        //爬虫分析器错误
	SPIDER_CRAWL_ERROR         = "crawer error"         //爬虫爬取服务错误

)
