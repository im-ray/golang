package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Blocking programm / Non Concurrent
	// for _, link := range links {
	// 	checkLink(link)
	// }

	// using goroutine with help of go keyword
	// for _, link := range links {
	// 	go checkLink(link)
	// }
	// the issues will be there if we use the above programm it may not get time to execute
	//  before finishing of main goroutine 
	// so channel is used to communicate between goroutine

	// learn more about channels
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	//fmt.Println(<- c)
	//fmt.Println(<- c)
	//fmt.Println(<- c)
	//fmt.Println(<- c)
	//fmt.Println(<- c)

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c) // waiting for message to recceive in channel
	// }

	
	// infinite fetch
	// for {
	// 	go checkLink(<-c, c) // interseting the first argument is a string
	// }

	// for l := range c {
	// 	go checkLink(l, c)
	// }

	// for l := range c {
	// 	time.Sleep(5 * time.Second)
	// 	go checkLink(l, c)
	// }

	// lets use function  literal for 
	for l := range c {
		go func (link string)  {
			time.Sleep(5*time.Second)
			checkLink(link, c)
		}(l) // extra set of paranthesis to invoke the function literal
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Might be Down!")
		//c <- "Might be Down I think"
		c <- link
		return
	}
	fmt.Println(link , "is up!")
	//c <- "Yep its up"
	c <- link
}



