package spider

import (
	"fmt"
	"testing"
)

func TestSpiderEvaluetStatus(t *testing.T) {
	applicationStatus = true
	statusChannel := make(chan int)
	textChannel := make(chan string)
	processChannel := make(chan bool)
	totalURLCount = 0
	urls = append(urls, "http://www.mastergoco.com/index1.html")
	urls = append(urls, "http://www.mastergoco.com/index2.html")
	urls = append(urls, "http://www.mastergoco.com/index3.html")
	urls = append(urls, "http://www.mastergoco.com/index4.html")
	urls = append(urls, "http://www.mastergoco.com/index5.html")
	fmt.Println("Starting spider")
	urlsProcessed = 0
	totalURLCount = len(urls)
	go SpiderEvaluetStatus(statusChannel, textChannel, processChannel)
	go SpiderURLsHander(statusChannel, textChannel)
	go SpiderContenrHander(textChannel, processChannel)
	for {
		if applicationStatus == false {
			fmt.Println(fullText)
			fmt.Println("Done!")
			break
		}
		select {
		case sC := <-statusChannel:
			fmt.Println("Message on StatusChannel", sC)
		}
	}
}
