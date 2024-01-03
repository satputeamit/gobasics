// concurrency is not parallelism

// concurrency  - we can have multiple threads executing code. if one thread blocks, another one
// is picked up and worked on

// parallelism -  we can do multiple things at same time. Multiple threads executed at same time.
// requires multiple cpu

package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{"http://google.com", "http://yahoo.com", "http://facebook.com"}
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) // go - creates new go routine/ child go routine
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

	// 	for l := range c {
	// 		fmt.Println("--->", l)
	// 		fmt.Println(<-c)
	// 	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		// fmt.Println(link + " link might be down")
		c <- link + " link might be down"
		return
	}
	// fmt.Println(link + " link is up")
	c <- link + " link is up"
}
