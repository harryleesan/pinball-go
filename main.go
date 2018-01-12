package main

import (
	"fmt"
	"net/http"
)

func main() {
	sites := []string{
		"https://google.com",
		"https://bing.com",
	}

	for _, site := range sites {
		ping(site)
	}

}

func ping(s string) {
	resp, err := http.Get(s)
	if err != nil {
		fmt.Println("Cannot access site:", s)
		fmt.Println(err)
	}

	fmt.Println("Site:", s, "works! Status code:", resp.StatusCode)
}
