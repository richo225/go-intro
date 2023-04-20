package main

import (
	"fmt"
	"net/http"
)

var websites = []string{
	"http://google.com",
	"http://facebook.com",
	"http://stackoverflow.com",
	"http://golang.org",
	"http://amazon.com",
}

func main() {
	c := make(chan string)

	for _, website := range websites {
		go checkWebsiteStatus(website, c)
	}

	for {
		go checkWebsiteStatus(<-c, c)
	}
}

func checkWebsiteStatus(website string, c chan string) {
	resp, err := http.Get(website)

	if err != nil {
		fmt.Println(website, "might be down:", err)
	}

	fmt.Println(website, "status:", resp.Status)
	c <- website
}
