package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)
type MetaDataKey struct {
	MetaUserName string `json:"username"`
	MetaAPIKey string `json:"api_key"`
	MetaUUID string `json:"uuid"`
	MetadaProNo string `json:"pro_no"`
	MetaTimeZone int`json:"timeZone"`
	MetaLoginHref string `json:"login_href"`
	MetaUserAgent string `json:"userAgent"`
	//MetaLoginType string  `json:"meta_login_type"`

}
func main() {
	medadata := &MetaDataKey{
		MetaUserName:  "mikenig716@gmail.com",
		MetaAPIKey:    "27bfc6d6c5da40a4c054a9653490cdfa",
		MetaUUID:      "M-000260",
		MetadaProNo:   "10001",
		MetaTimeZone:  8,
		MetaLoginHref: "https://www.amazon.com/ap/signin?_encoding=UTF8&openid.assoc_handle=usflex&openid.claimed_id=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.identity=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.mode=checkid_setup&openid.ns=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0&openid.ns.pape=http%3A%2F%2Fspecs.openid.net%2Fextensions%2Fpape%2F1.0&openid.pape.max_auth_age=0&openid.return_to=https%3A%2F%2Fwww.amazon.com%2Fgp%2Fyourstore%2Fhome%3Fie%3DUTF8%26action%3Dsign-out%26path%3D%252Fgp%252Fyourstore%252Fhome%26ref_%3Dnav_AccountFlyout_signout%26signIn%3D1%26useRedirectOnSuccess%3D1",
		MetaUserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:80.0) Gecko/20100101 Firefox/80.0",
	}
	data:=[]byte(
		`{
    "username": "mikenig716@gmail.com",
    "api_key": "27bfc6d6c5da40a4c054a9653490cdfa",
    "uuid": "M-000260",
    "pro_no": "10001",
    "timeZone": 8,
    "login_href": "https://www.amazon.com/ap/signin?_encoding=UTF8&openid.assoc_handle=usflex&openid.claimed_id=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.identity=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.mode=checkid_setup&openid.ns=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0&openid.ns.pape=http%3A%2F%2Fspecs.openid.net%2Fextensions%2Fpape%2F1.0&openid.pape.max_auth_age=0&openid.return_to=https%3A%2F%2Fwww.amazon.com%2Fgp%2Fyourstore%2Fhome%3Fie%3DUTF8%26action%3Dsign-out%26path%3D%252Fgp%252Fyourstore%252Fhome%26ref_%3Dnav_AccountFlyout_signout%26signIn%3D1%26useRedirectOnSuccess%3D1",
    "userAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:80.0) Gecko/20100101 Firefox/80.0"
}`)
	_=data
	payload, err := json.Marshal(medadata)
	if err != nil {
		fmt.Println(err)
	}

	c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))
	c.OnRequest(func(request *colly.Request) {
	/*	requestBody,err:= ioutil.ReadAll(request.Body)
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Printf("request body is %s\n",string(requestBody))*/
		//request.Headers.Set("Content-Type", "application/x-www-form-urlencoded")
	//	fmt.Println(request.Method, request.URL)

		request.Headers.Set("Content-Type", "application/json;charset=UTF-8")
		//request.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:80.0) Gecko/20100101 Firefox/80.0")
		request.Headers.Set("Accept-Encoding", "gzip, deflate")
		request.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9")
	//	request.Headers.Set("Accept","*/*")
	})
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("--->", response.StatusCode)
		resultBody := string(response.Body)
		fmt.Println(resultBody)
	})
	//err =c.Request(http.MethodPost,"http://www.metadata1.com/api/metadata_api/metadata1",strings.NewReader(string(data)),nil,nil)
	c.OnError(func(response *colly.Response, e error) {
		fmt.Println("eeror is ",e)
	})
	err = c.PostRaw("http://www.metadata1.com/api/metadata_api/metadata1",payload)

	if err != nil {
		fmt.Println(err)
	}

	/*resp, err := http.Post("http://www.metadata1.com/api/metadata_api/metadata1",
		"application/x-www-form-urlencoded", strings.NewReader(string(data)))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))*/

}
