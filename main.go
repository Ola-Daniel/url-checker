// a website is up if it returns an http response code of 200
// check every 5 seconds
// package main
// takes a user input url and checks if it is up or down
// check if the url is valid
package main

import (
	"fmt"
	"net/http"
	"time"
)


func main() {
	var url string
	fmt.Println("Enter a url: ")
	fmt.Scan(&url)
	for {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println("Status: ", resp.Status)
		}
		time.Sleep(5 * time.Second)
	}
}

