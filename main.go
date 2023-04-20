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
		resp, err := http.Get(website)

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println(website, "status:", resp.Status)
	}
}
