package crawl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
)

var applicationStatus bool
var urls []string
var urlsProcessed int
var foundUrls []string
var fullText string
var totalURLCount int
var wg sync.WaitGroup
var v1 int

type PerchSpider struct {
	UserAgent string // UserAgent is the User-Agent string used by HTTP requests
	MaxDepth  int    // MaxDepth limits the recursion depth of visited URLs. Set it to 0 for infinite recursion (default).
	// AllowedDomains is a domain whitelist.
	// Leave it blank to allow any domains to be visited
	AllowDomain    []string
	NotAllowDomain []string

	DisallowedDomains []string
	// DisallowedURLFilters is a list of regular expressions which restricts
	// visiting URLs. If any of the rules matches to a URL the
	// request will be stopped. DisallowedURLFilters will
	// be evaluated before URLFilters
	// Leave it blank to allow any URLs to be visited
	DisallowedURLFilters []*regexp.Regexp
	// URLFilters is a list of regular expressions which restricts
	// visiting URLs. If any of the rules matches to a URL the
	// request won't be stopped. DisallowedURLFilters will
	// be evaluated before URLFilters

	// Leave it blank to allow any URLs to be visited
	URLFilters []*regexp.Regexp

	// AllowURLRevisit allows multiple downloads of the same URL
	AllowURLRevisit bool
	// MaxBodySize is the limit of the retrieved response body in bytes.
	// 0 means unlimited.
	// The default value for MaxBodySize is 10MB (10 * 1024 * 1024 bytes).
	MaxBodySize int
	// CacheDir specifies a location where GET requests are cached as files.
	// When it's not defined, caching is disabled.
	CacheDir string
	// IgnoreRobotsTxt allows the Collector to ignore any restrictions set by
	// the target host's robots.txt file.  See http://www.robotstxt.org/ for more
	// information.
	IgnoreRobotsTxt bool
	// Async turns on asynchronous network communication. Use Collector.Wait() to
	// be sure all requests have been finished.
	Async bool
	// ParseHTTPErrorResponse allows parsing HTTP responses with non 2xx status codes.
	// By default, Colly parses only successful HTTP responses. Set ParseHTTPErrorResponse
	// to true to enable it.
	ParseHTTPErrorResponse bool
	// ID is the unique identifier of a collector
	ID uint32
	// DetectCharset can enable character encoding detection for non-utf8 response bodies
	// without explicit charset declaration. This feature uses https://github.com/saintfish/chardet
	DetectCharset bool
	// RedirectHandler allows control on how a redirect will be managed
	// use c.SetRedirectHandler to set this value
	redirectHandler func(req *http.Request, via []*http.Request) error
	// CheckHead performs a HEAD request before every GET to pre-validate the response
	CheckHead bool
	// TraceHTTP enables capturing and reporting request performance for crawler tuning.
	// When set to true, the Response.Trace will be filled in with an HTTPTrace object.
	// TraceHTTP enables capturing and reporting request performance for crawler tuning.
	// When set to true, the Response.Trace will be filled in with an HTTPTrace object.
	TraceHTTP bool

	//robotsMap                map[string]*robotstxt.RobotsData
	//htmlCallbacks            []*htmlCallbackContainer
	//xmlCallbacks             []*xmlCallbackContainer
	//requestCallbacks         []RequestCallback
	//responseCallbacks        []ResponseCallback
	//responseHeadersCallbacks []ResponseHeadersCallback
	////errorCallbacks           []ErrorCallback
	//	scrapedCallbacks         []ScrapedCallback
	requestCount  uint32
	responseCount uint32
	//backend                  *httpBackend
	wg   *sync.WaitGroup
	lock *sync.RWMutex
}

/**
读取ulr
*/
func SpiderURLsHander(statusChannel chan int, textChannel chan string) {

	for i := 0; i < totalURLCount; i++ {
		resp, _ := http.Get(urls[i])
		text, err := ioutil.ReadAll(resp.Body)
		textChannel <- string(text)
		if err != nil {
			//handler err
		}
		statusChannel <- 0
	}

}

func SpiderContenrHander(textchannel chan string, processChannel chan bool) {
	for {
		select {
		case pC := <-processChannel:
			if pC == true {
				//hang on
			}
			if pC == false {
				close(textchannel)
				close(processChannel)
			}
		case tC := <-textchannel:
			fullText += tC

		}
	}
}

func SpiderEvaluetStatus(statusChannel chan int, textChannel chan string, processChannel chan bool) {
	for {
		select {
		case status := <-statusChannel:
			fmt.Println(urlsProcessed, totalURLCount)
			urlsProcessed++
			if status == 0 {
				fmt.Println("Got url")
			}
			if status == 1 {
				close(statusChannel)
			}
			if urlsProcessed == totalURLCount {
				processChannel <- false
				applicationStatus = false
			}

		}
	}

}
