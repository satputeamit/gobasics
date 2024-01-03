package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {

	links := []string{
		"http://google.com",
		"http://fb.com",
		"http://youtube.com",
	}

	for _, link := range links {
		go getStatusCode(link)
		wg.Add(1)
	}

	wg.Wait()

}

func getStatusCode(link string) {
	defer wg.Done()
	res, err := http.Get(link)
	if err != nil {
		fmt.Printf("Down %s\n", link)
	} else {
		fmt.Printf("%d of link %s\n", res.StatusCode, link)
	}
}
