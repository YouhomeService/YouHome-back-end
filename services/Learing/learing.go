package main

import (
	"net/http/httputil"
	"net/http"
	"strings"
	"fmt"
	"log"
)

/*
var userService = "http://172.22.16.5:9091"
var deviceService ="http://172.22.16.4:9092"
var sceneService = "http://172.22.16.3:9093"
 */
var userService = "localhost:9091"
var deviceService ="localhost:9092"
var sceneService = "localhost:9093"
func NewMultipleHostsReverseProxy() *httputil.ReverseProxy {
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		//req.URL.Path = target.Path
		path := req.URL.Path
		data := strings.Split(path,"/")
		service := strings.Split(data[2],"?")[0]
		switch service{
			case "users":
				req.URL.Host = userService
			case "scenes":
				req.URL.Host = sceneService
			case "devices":
				req.URL.Host = deviceService
		}

		fmt.Println(req.URL.Scheme)
		fmt.Println(req.URL.Host)
		fmt.Println(req.URL.Path)
	}
	return &httputil.ReverseProxy{Director: director}
}

func main() {
	proxy := NewMultipleHostsReverseProxy()
	log.Fatal(http.ListenAndServe(":8088", proxy))
}