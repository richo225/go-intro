package main

import (
	"fmt"
	"net/http"
	"os"
)

var websites = []string{
	"http://google.com",
	"http://facebook.com",
	"http://stackoverflow.com",
	"http://golang.org",
	"http://amazon.com",
}

func main() {
	for _, website := range websites {
		checkWebsiteStatus(website)
	}
}

func checkWebsiteStatus(website string) {
	resp, err := http.Get(website)

	if err != nil {
		fmt.Println(website, "might be down: ", err)
		os.Exit(1)
	}

	fmt.Println(website, "status:", resp.Status)
}
