package spider

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
