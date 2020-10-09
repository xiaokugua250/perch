package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func sysUserProxy(w http.ResponseWriter,r *http.Request){
	path,err:= url.Parse("http://127.0.0.1:8081")
	if err!= nil{
		fmt.Printf("error is %s",err.Error())
	}
	//r.URL.Path = strings.Replace(r.RequestURI, "/forum", "", 1)
	fmt.Println("===>",r.URL.Path)
	r.URL.Path=r.URL.Path[9:]
	fmt.Println("===>",r.URL.Path)
	httputil.NewSingleHostReverseProxy(path).ServeHTTP(w,r)
}

func main()  {
	fmt.Println("====beigin to run local proxy in 0.0.0.0:80 =====")
	router := mux.NewRouter()

	router.PathPrefix("/sys-user/").HandlerFunc(sysUserProxy)
	server := &http.Server{
		Addr:"0.0.0.0:80",
		Handler:router,
	}
	log.Fatal(server.ListenAndServe())

}
