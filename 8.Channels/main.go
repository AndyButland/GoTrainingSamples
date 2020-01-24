package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
	}

	// Synchronous version - works, but we have to wait for each check to complete.
	//v1(links)

	// Asynchronous version with go routine - starts to work, but we can't tell, as program ends before
	// checks are made.
	//v2(links)

	// We can use a channel to recieve messages back from the go routines, pausing program flow.
	// But we have to make sure we read from the channel the number of times we expect to have
	// responses.
	//v3(links)

	// Use a loop to get channel responses, pause with local function and re-check.
	v4(links)
}

func v1(links []string) {
	for _, link := range links {
		checkLink(link)
	}
}

func v2(links []string) {
	for _, link := range links {
		go checkLink(link)
	}
}

func checkLink(link string) {
	fmt.Printf("Checking %s\n", link)
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s seems to be down.\n", link)
		return
	}

	fmt.Printf("%s is up!\n", link)
}

func v3(links []string) {

	c := make(chan string)

	for _, link := range links {
		go checkLinkAndUpdateChannel(c, link)
	}

	fmt.Printf("Checked %s\n", <-c)
	fmt.Printf("Checked %s\n", <-c)
	fmt.Printf("Checked %s\n", <-c)
	fmt.Printf("Checked %s\n", <-c)
}

func v4(links []string) {

	c := make(chan string)

	for _, link := range links {
		go checkLinkAndUpdateChannel(c, link)
	}

	for l := range c {
		fmt.Printf("Checked %s\n", l)
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLinkAndUpdateChannel(c, link)
		}(l)
	}
}

func checkLinkAndUpdateChannel(c chan string, link string) {
	fmt.Printf("Checking %s\n", link)
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s seems to be down.\n", link)
		c <- link
		return
	}

	fmt.Printf("%s is up!\n", link)
	c <- link
}
