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
	c := make(chan string)

	for _, website := range websites {
		go checkWebsiteStatus(website, c)
	}

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func checkWebsiteStatus(website string, c chan string) {
	resp, err := http.Get(website)

	if err != nil {
		c <- fmt.Sprintln(website, "might be down:", err)
		os.Exit(1)
	}

	c <- fmt.Sprintln(website, "status:", resp.Status)
}
