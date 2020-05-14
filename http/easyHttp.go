package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main1() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	response, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", string(response))



}
